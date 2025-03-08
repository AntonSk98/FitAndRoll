package participants

import (
	"fit_and_roll/backend/config"
	"fit_and_roll/backend/mappers"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type MemberCardController struct {
	dbManager *config.DatabaseManager
}

func NewMemberCardController(dbManager *config.DatabaseManager) *MemberCardController {
	return &MemberCardController{dbManager: dbManager}
}

func (controller *MemberCardController) FindAllMemberCards(participantId uint) ([]MemberCardInfo, error) {
	var participantMemberCards []MemberCard
	var participantMemberCardInfos []MemberCardInfo

	baseQuery := controller.dbManager.DB.Model(&MemberCard{}).Unscoped().Where("participant_id = ?", participantId)

	if err := baseQuery.
		Order("created_at desc").
		Find(&participantMemberCards).Error; err != nil {
		return nil, err
	}

	for _, memparticipantMemberCard := range participantMemberCards {
		participantMemberCardInfos = append(participantMemberCardInfos, *NewMemberCardInfo(memparticipantMemberCard))
	}

	return participantMemberCardInfos, nil
}

func (controller *MemberCardController) IssueNewMemberCard(participantId uint) error {
	return controller.dbManager.Transactional(func(tx *gorm.DB) error {
		participantPointer, err := requireParticipant(participantId, tx)
		if err != nil {
			return err
		}

		memberCard := NewMemberCard(participantPointer.ID)

		if err := tx.Create(&memberCard).Error; err != nil {
			return err
		}

		return nil
	})
}

func (controller *MemberCardController) UndoIssueNewMemberCard(participantId uint, memberCardId uint) error {
	var memberCard MemberCard

	if err := controller.dbManager.DB.Where("id = ? AND participant_id = ?", memberCardId, participantId).First(&memberCard).Error; err != nil {
		return err
	}

	if !memberCard.isUnused() {
		return fmt.Errorf("unable to delete the member card as it has already been used")
	}

	if err := controller.dbManager.DB.Unscoped().Delete(&memberCard).Error; err != nil {
		return err
	}

	return nil
}

func (controller *MemberCardController) LoadMemberCardCourseHistory(participantId uint, memberCardId uint) ([]MemberCardHistoryEntry, error) {

	var courseAttendanceProjection []struct {
		CourseName string    `gorm:"column:name"`
		AttendedAt time.Time `gorm:"column:created_at"`
	}

	err := controller.dbManager.DB.
		Model(&CourseAttendance{}).
		Select("courses.name, course_attendances.created_at").
		Joins("JOIN courses ON courses.id = course_attendances.course_id").
		Where("course_attendances.participant_id = ? AND course_attendances.member_card_id = ?", participantId, memberCardId).
		Order("course_attendances.created_at DESC").
		Scan(&courseAttendanceProjection).Error

	if err != nil {
		return nil, err
	}

	var historyEntries []MemberCardHistoryEntry

	for _, courseAttendanceEntry := range courseAttendanceProjection {
		memberCardHistoryEntry := &MemberCardHistoryEntry{
			Course:     courseAttendanceEntry.CourseName,
			AttendedAt: mappers.ToDateString(courseAttendanceEntry.AttendedAt),
		}

		historyEntries = append(historyEntries, *memberCardHistoryEntry)
	}

	fmt.Println(historyEntries)

	return historyEntries, nil
}

func requireParticipant(participantId uint, db *gorm.DB) (*Participant, error) {
	var participant Participant

	if err := db.First(&participant, participantId).Error; err != nil {
		return nil, err
	}

	return &participant, nil
}
