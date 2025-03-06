package participants

import "fit_and_roll/backend/config"

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

// findParticipant(name or surname...) -> {name|surname|hasValidMemberCard}
// takePartInCourse(participantId, courseId, actionType) -> CourseParticipantEntry(manyToOne to Course... manyToOne to Participant)
// undoTakePartInCourse(participantId, courseId, actionType)
