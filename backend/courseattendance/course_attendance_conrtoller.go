package courseattendance

import (
	"fit_and_roll/backend/common"
	"fit_and_roll/backend/config"
	"fit_and_roll/backend/membercardattendance"
	"fmt"
)

type CourseAttendanceController struct {
	dbManager *config.DatabaseManager
}

func NewCourseAttendanceController(dbManager *config.DatabaseManager) *CourseAttendanceController {
	return &CourseAttendanceController{dbManager: dbManager}
}

func (controller *CourseAttendanceController) FindCourseAttendanceHistory(courseAttendanceParameters CourseAttendanceParameters, pageParams common.PageParams) (*common.Page[CourseAttendanceDto], error) {
	var total int64
	var results []CourseAttendanceDto

	query := controller.dbManager.DB.Model(&membercardattendance.MemberCardAttendance{}).
		Select("CONCAT(participants.name, ' ', participants.surname) as full_name, " +
			"courses.name as course, " +
			"member_card_attendances.created_at as attended_at, " +
			"member_card_attendances.attendance_type as attendance_type").
		Joins("JOIN participants ON participants.id = member_card_attendances.participant_id").
		Joins("JOIN courses ON courses.id = member_card_attendances.course_id").
		Unscoped()

	if courseAttendanceParameters.CourseID != nil {
		query = query.Where("courses.id = ?", *courseAttendanceParameters.CourseID)
	}

	if courseAttendanceParameters.FullNameLike != nil {
		query = query.Where("LOWER(CONCAT(participants.name, ' ', participants.surname)) LIKE ?", "%"+*courseAttendanceParameters.FullNameLike+"%")
	}

	if courseAttendanceParameters.CourseLike != nil {
		query = query.Where("LOWER(courses.name) LIKE ?", "%"+*courseAttendanceParameters.CourseLike+"%")
	}

	if courseAttendanceParameters.ExcludeArchivedCourses {
		query = query.Where("courses.deleted IS NULL")
	}

	if courseAttendanceParameters.ExcludeTrialAttendanced {
		query = query.Where("member_card_attendances.attendance_type != ?", membercardattendance.TrialAttendance)
	}

	if courseAttendanceParameters.ExcludeAttendanceWithMemberCard {
		query = query.Where("member_card_attendances.attendance_type != ?", membercardattendance.WithMemberCard)
	}

	if courseAttendanceParameters.ExcludeAttendanceWitouthMemberCard {
		query = query.Where("member_card_attendances.attendance_type != ?", membercardattendance.WithoutMemberCard)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := query.Scopes(controller.dbManager.Paginate(pageParams.Page, pageParams.Size)).
		Order("member_card_attendances.created_at DESC").
		Scan(&results).Error; err != nil {
		return nil, err
	}

	for index := range results {
		mappedDateTime, err := common.ToDateTimeString(results[index].AttendedAt)
		if err != nil {
			return nil, err
		}
		results[index].AttendedAt = mappedDateTime
	}

	fmt.Println(pageParams)

	return &common.Page[CourseAttendanceDto]{
		Data:  results,
		Total: int(total),
		Page:  pageParams.Page,
		Size:  pageParams.Size,
	}, nil
}
