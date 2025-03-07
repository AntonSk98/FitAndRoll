package participants

import (
	"fit_and_roll/backend/config"
	"fmt"

	"gorm.io/gorm"
)

type MemberCardCourseParticipationController struct {
	dbManager *config.DatabaseManager
}

func NewMemberCardCourseParticipationController(dbManager *config.DatabaseManager) *MemberCardCourseParticipationController {
	return &MemberCardCourseParticipationController{dbManager: dbManager}
}

func (controller *MemberCardCourseParticipationController) FindActiveMemberCards(participant uint) ([]MemberCardInfo, error) {
	var activeMemberCards []MemberCard
	result := controller.dbManager.DB.Where("participant_id = ?", participant).Find(&activeMemberCards)
	if result.Error != nil {
		return nil, result.Error
	}

	var activeMemberCardInfo []MemberCardInfo

	for _, activeMemberCard := range activeMemberCards {
		activeMemberCardInfo = append(activeMemberCardInfo, *NewMemberCardInfo(activeMemberCard))
	}
	return activeMemberCardInfo, nil
}

func (controller *MemberCardCourseParticipationController) AttendCourse(courseAttendanceCommand CourseAttendanceCommand) error {
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

func (controller *MemberCardCourseParticipationController) attendWithMemberCard(tx *gorm.DB, memberCardID uint) error {
	var memberCard MemberCard
	if err := tx.First(&memberCard, memberCardID).Error; err != nil {
		return fmt.Errorf("could not find member card with ID %v: %w", memberCardID, err)
	}

	// Handle course visit logic based on the capacity
	if !memberCard.visitCourse() {
		// Mark card as used if it cannot be visited due to capacity
		memberCard.markAsUsed()
	}

	// Save updated member card state
	if err := tx.Save(&memberCard).Error; err != nil {
		return fmt.Errorf("failed to save member card %v: %w", memberCard.ID, err)
	}

	return nil
}

// findParticipant(name or surname...) -> {name|surname|hasValidMemberCard}
// takePartInCourse(participantId, courseId, actionType) -> CourseParticipantEntry(manyToOne to Course... manyToOne to Participant)
// undoTakePartInCourse(participantId, courseId, actionType)
