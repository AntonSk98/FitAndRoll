package courseattendance

import (
	"fit_and_roll/backend/config"
	"fit_and_roll/backend/membercardattendance"
)

type CourseAttendanceController struct {
	dbManager *config.DatabaseManager
}

func NewCourseAttendanceController(dbManager *config.DatabaseManager) *CourseAttendanceController {
	return &CourseAttendanceController{dbManager: dbManager}
}

func (controller *CourseAttendanceController) FindOverallAttendanceHistory() ([]CourseAttendanceDto, error) {
	var results []CourseAttendanceDto

	err := controller.dbManager.DB.Model(&membercardattendance.MemberCardAttendance{}).
		Select("CONCAT(participants.name, ' ', participants.surname) as full_name, " +
			"courses.name as course, " +
			"CASE WHEN courses.deleted IS NOT NULL THEN true ELSE false END AS archived_course, " +
			"DATE(member_card_attendances.created_at) as attended_at, " +
			"member_card_attendances.attendance_type as attendance_type").
		Joins("JOIN participants ON participants.id = member_card_attendances.participant_id").
		Joins("JOIN courses ON courses.id = member_card_attendances.course_id").
		Unscoped().
		Order("member_card_attendances.created_at DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}
