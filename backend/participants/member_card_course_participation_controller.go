package participants

import "fit_and_roll/backend/config"

type MemberCardCourseParticipationController struct {
	dbManager *config.DatabaseManager
}

func NewMemberCardCourseParticipationController(dbManager *config.DatabaseManager) *MemberCardCourseParticipationController {
	return &MemberCardCourseParticipationController{dbManager: dbManager}
}

// findParticipant(name or surname...) -> {name|surname|hasValidMemberCard}
// takePartInCourse(participantId, courseId, actionType) -> CourseParticipantEntry(manyToOne to Course... manyToOne to Participant)
// undoTakePartInCourse(participantId, courseId, actionType)
