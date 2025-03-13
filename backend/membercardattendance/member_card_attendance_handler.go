package membercardattendance

import (
	"fit_and_roll/backend/common"
	"fit_and_roll/backend/config"
	"fit_and_roll/backend/participants"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Provides methods to to fetch active member cards, ability to attend a course with a member card, and the history of attended course with a member card.
type MemberCardAttendanceHandler struct {
	dbManager *config.DatabaseManager
}

// Creates a new instance of the handler
func NewMemberCardAttendanceHandler(dbManager *config.DatabaseManager) *MemberCardAttendanceHandler {
	return &MemberCardAttendanceHandler{dbManager: dbManager}
}

// Finds all active member cards by participant identifier.
// A member card is considered valid if it still has a capacity that could be used to attend a course.
func (controller *MemberCardAttendanceHandler) FindActiveMemberCards(participant uint) ([]participants.MemberCardInfo, error) {
	var activeMemberCards []participants.MemberCard
	result := controller.dbManager.DB.Where("participant_id = ?", participant).Find(&activeMemberCards)
	if result.Error != nil {
		return nil, result.Error
	}

	var activeMemberCardInfo []participants.MemberCardInfo

	for _, activeMemberCard := range activeMemberCards {
		activeMemberCardInfo = append(activeMemberCardInfo, *participants.NewMemberCardInfo(activeMemberCard))
	}
	return activeMemberCardInfo, nil
}

// Records a participant's attendance for a course.
// This method creates an entry that answers the question:
//   - "Which user attended which course, and if attended using a member card, which member card was used?"
func (controller *MemberCardAttendanceHandler) AttendCourse(courseAttendanceCommand CourseAttendanceCommand) error {
	return controller.dbManager.Transactional(func(tx *gorm.DB) error {
		// Handle Member Card processing if it's provided
		if courseAttendanceCommand.MemberCardID != nil {
			if err := controller.attendWithMemberCard(tx, *courseAttendanceCommand.MemberCardID); err != nil {
				return err
			}
		}

		// Create Course Attendance from the command
		courseAttendance, err := NewCourseAttendance(courseAttendanceCommand)
		if err != nil {
			return err
		}

		// Save Course Attendance
		if err := tx.Create(courseAttendance).Error; err != nil {
			return err
		}

		return nil
	})
}

// Loads the history of attended courses by passed member card and participant identifiers.
func (controller *MemberCardAttendanceHandler) LoadMemberCardCourseHistory(participantId uint, memberCardId uint) ([]MemberCardHistoryEntry, error) {

	var courseAttendanceProjection []struct {
		CourseName string    `gorm:"column:name"`
		AttendedAt time.Time `gorm:"column:created_at"`
	}

	err := controller.dbManager.DB.
		Model(&MemberCardAttendance{}).
		Select("courses.name, member_card_attendances.created_at").
		Joins("JOIN courses ON courses.id = member_card_attendances.course_id").
		Unscoped().
		Where("member_card_attendances.participant_id = ? AND member_card_attendances.member_card_id = ?", participantId, memberCardId).
		Order("member_card_attendances.created_at DESC").
		Scan(&courseAttendanceProjection).Error

	if err != nil {
		return nil, err
	}

	var historyEntries []MemberCardHistoryEntry

	for _, courseAttendanceEntry := range courseAttendanceProjection {
		memberCardHistoryEntry := &MemberCardHistoryEntry{
			Course:     courseAttendanceEntry.CourseName,
			AttendedAt: common.ToDateString(courseAttendanceEntry.AttendedAt),
		}

		historyEntries = append(historyEntries, *memberCardHistoryEntry)
	}

	return historyEntries, nil
}

func (controller *MemberCardAttendanceHandler) attendWithMemberCard(tx *gorm.DB, memberCardID uint) error {
	var memberCard participants.MemberCard
	if err := tx.First(&memberCard, memberCardID).Error; err != nil {
		return fmt.Errorf("could not find member card with ID %v: %w", memberCardID, err)
	}

	// Handle course visit logic based on the capacity
	memberCard.VisitCourse()

	// Save updated member card state
	if err := tx.Save(&memberCard).Error; err != nil {
		return fmt.Errorf("failed to save member card %v: %w", memberCard.ID, err)
	}

	return nil
}
