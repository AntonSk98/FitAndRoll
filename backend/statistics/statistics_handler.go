package statistics

import (
	"fit_and_roll/backend/common"
	"fit_and_roll/backend/config"
	"fit_and_roll/backend/membercardattendance"
	"fmt"

	"gorm.io/gorm"
)

// Handler to fetch statistics information
type StatisticsHandler struct {
	dbManager *config.DatabaseManager
}

// Creates a new handler of StatisticsHandler
func NewStatisticsHandler(dbManager *config.DatabaseManager) *StatisticsHandler {
	return &StatisticsHandler{dbManager: dbManager}
}

// Returns paginated participant statistics
func (handler *StatisticsHandler) ParticipantStatistics(params StatisticsParams, pageParams common.PageParams) (*common.Page[StatisticsDto], error) {
	query := handler.dbManager.DB.Model(&membercardattendance.MemberCardAttendance{}).
		Select(`CONCAT(participants.name, ' ', participants.surname) AS name,
		COUNT(CASE WHEN member_card_attendances.attendance_type = ? THEN 1 ELSE NULL END) AS with_member_card,
		COUNT(CASE WHEN member_card_attendances.attendance_type = ? THEN 1 ELSE NULL END) AS trial_attendance,
		COUNT(CASE WHEN member_card_attendances.attendance_type = ? THEN 1 ELSE NULL END) AS without_member_card`,
			membercardattendance.WithMemberCard,
			membercardattendance.TrialAttendance,
			membercardattendance.WithoutMemberCard).
		Joins("JOIN participants ON participants.id = member_card_attendances.participant_id").
		Group("participants.name")

	if params.Name != nil {
		query.Where("LOWER(CONCAT(participants.name, ' ', participants.surname)) LIKE ?", "%"+*params.Name+"%")
	}

	return handler.paginate(query, params, pageParams)
}

// Returns paginated course statistics
func (handler *StatisticsHandler) CourseStatistics(params StatisticsParams, pageParams common.PageParams) (*common.Page[StatisticsDto], error) {
	query := handler.dbManager.DB.Model(&membercardattendance.MemberCardAttendance{}).
		Select(`courses.name AS name,
		COUNT(CASE WHEN member_card_attendances.attendance_type = ? THEN 1 ELSE NULL END) AS with_member_card,
		COUNT(CASE WHEN member_card_attendances.attendance_type = ? THEN 1 ELSE NULL END) AS trial_attendance,
		COUNT(CASE WHEN member_card_attendances.attendance_type = ? THEN 1 ELSE NULL END) AS without_member_card`,
			membercardattendance.WithMemberCard,
			membercardattendance.TrialAttendance,
			membercardattendance.WithoutMemberCard).
		Joins("JOIN courses ON courses.id = member_card_attendances.course_id").
		Group("courses.name")

	if params.Name != nil {
		query.Where("LOWER(courses.name) LIKE ?", "%"+*params.Name+"%")
	}

	return handler.paginate(query, params, pageParams)
}

func (handler *StatisticsHandler) paginate(query *gorm.DB, params StatisticsParams, pageParams common.PageParams) (*common.Page[StatisticsDto], error) {
	var statistics []StatisticsDto
	var total int64

	if params.AttendedRange != nil {
		if err := handler.withRange(query, params); err != nil {
			return nil, err
		}
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count statistics: %w", err)
	}

	if err := query.
		Scopes(handler.dbManager.Paginate(pageParams.Page, pageParams.Size)).
		Order("with_member_card DESC, trial_attendance DESC, without_member_card DESC").
		Scan(&statistics).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch statistics: %w", err)
	}

	return &common.Page[StatisticsDto]{
		Data:  statistics,
		Total: int(total),
		Page:  pageParams.Page,
		Size:  pageParams.Size,
	}, nil
}

func (handler *StatisticsHandler) withRange(query *gorm.DB, params StatisticsParams) error {
	if params.AttendedRange.From != nil {
		from, err := common.ConvertUTCToBerlin(*params.AttendedRange.From)
		if err != nil {
			return err
		}
		query.Where("member_card_attendances.created_at >= ?", from)
	}
	if params.AttendedRange.To != nil {
		to, err := common.ConvertUTCToBerlin(*params.AttendedRange.To)
		if err != nil {
			return err
		}
		query.Where("member_card_attendances.created_at <= ?", to)
	}

	return nil
}
