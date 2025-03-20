package courseattendance

import (
	"fit_and_roll/backend/common"
	"fit_and_roll/backend/config"
	"fit_and_roll/backend/membercardattendance"
	"fit_and_roll/backend/participants"
	"fmt"

	"gorm.io/gorm"
)

// CourseAttendanceHandler retrieves the attendance history of participants in courses.
// It provides functionality to fetch either the overall participation history or history filtered by specific course identifier.
// The handler relies on the provided database manager to query the database for relevant attendance data.
type CourseAttendanceHandler struct {
	dbManager *config.DatabaseManager
}

// NewCourseAttendanceHandler creates a new CourseAttendanceHandler.
//
// It initializes the handler with a database manager instance.
//
// Parameters:
//   - dbManager: A pointer to the database manager instance.
//
// Returns:
//   - A pointer to the CourseAttendanceHandler instance.
func NewCourseAttendanceHandler(dbManager *config.DatabaseManager) *CourseAttendanceHandler {
	return &CourseAttendanceHandler{dbManager: dbManager}
}

// FindCourseAttendanceHistory retrieves the course attendance history based on the provided filters and pagination parameters.
//
// Parameters:
//   - courseAttendanceParameters: Filters to apply when querying attendance records.
//   - pageParams: Pagination parameters (page number and page size).
//
// Returns:
//   - A paginated list of CourseAttendanceDto containing attendance records.
//   - An error if the operation fails.
func (controller *CourseAttendanceHandler) FindCourseAttendanceHistory(courseAttendanceParameters CourseAttendanceParameters, pageParams common.PageParams) (*common.Page[CourseAttendanceDto], error) {
	var total int64
	var results []CourseAttendanceDto

	// Initialize base query
	query := controller.baseQuery()

	// Apply filters
	err := controller.applyFilters(query, courseAttendanceParameters)

	if err != nil {
		return nil, err
	}

	// Count total results
	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count attendance records: %w", err)
	}

	// Fetch paginated results
	if err := query.
		Scopes(controller.dbManager.Paginate(pageParams.Page, pageParams.Size)).
		Order("member_card_attendances.created_at DESC").
		Scan(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch attendance history: %w", err)
	}

	if err := controller.withReadableAttendedAtDate(results); err != nil {
		return nil, err
	}

	return &common.Page[CourseAttendanceDto]{
		Data:  results,
		Total: int(total),
		Page:  pageParams.Page,
		Size:  pageParams.Size,
	}, nil
}

// Removes attendance from the course attendance history.
// If a participant attended a course using a member card, the attendance is restored.
func (handler *CourseAttendanceHandler) UndoCourseAttendance(id *uint) error {
	if id == nil {
		return fmt.Errorf("course attendance history entry cannot be null")
	}

	return handler.dbManager.Transactional(func(tx *gorm.DB) error {
		var memberCardAttendance membercardattendance.MemberCardAttendance
		if err := tx.First(&memberCardAttendance, *id).Error; err != nil {
			return err
		}

		if memberCardAttendance.AttendanceType == membercardattendance.WithMemberCard {
			var memberCard participants.MemberCard
			if err := tx.Unscoped().First(&memberCard, *memberCardAttendance.MemberCardID).Error; err != nil {
				return err
			}
			memberCard.UndoAttendance()
			tx.Save(&memberCard)
		}

		tx.Unscoped().Delete(&memberCardAttendance)

		return nil
	})
}

// withReadableAttendedAtDate formats the 'AttendedAt' field in each record of the provided
// CourseAttendanceDto slice to a more readable date format. If an error occurs while
// formatting any date, it returns an error.
//
// This method modifies the 'AttendedAt' field in the original results slice to use a
// human-readable datetime format (e.g., "2006-01-02 15:04").
func (controller *CourseAttendanceHandler) withReadableAttendedAtDate(results []CourseAttendanceDto) error {
	for i := range results {
		formattedDate, err := common.ToDateTimeString(results[i].AttendedAt)
		if err != nil {
			return fmt.Errorf("failed to format date: %w", err)
		}
		results[i].AttendedAt = formattedDate
	}
	return nil
}

// baseQuery initializes the base query for attendance history
func (controller *CourseAttendanceHandler) baseQuery() *gorm.DB {
	return controller.dbManager.DB.Model(&membercardattendance.MemberCardAttendance{}).
		Select(`
			member_card_attendances.id,
			CONCAT(participants.name, ' ', participants.surname) AS full_name,
			courses.name AS course,
			member_card_attendances.created_at AS attended_at,
			member_card_attendances.attendance_type AS attendance_type`).
		Joins("JOIN participants ON participants.id = member_card_attendances.participant_id").
		Joins("JOIN courses ON courses.id = member_card_attendances.course_id").
		Unscoped()
}

// applyFilters applies filtering conditions based on provided parameters
func (controller *CourseAttendanceHandler) applyFilters(query *gorm.DB, params CourseAttendanceParameters) error {
	if params.CourseID != nil {
		query.Where("courses.id = ?", *params.CourseID)
	}

	if params.FullNameLike != nil {
		query.Where("LOWER(CONCAT(participants.name, ' ', participants.surname)) LIKE ?", "%"+*params.FullNameLike+"%")
	}

	if params.CourseLike != nil {
		query.Where("LOWER(courses.name) LIKE ?", "%"+*params.CourseLike+"%")
	}

	if params.ExcludeArchivedCourses {
		query.Where("courses.deleted IS NULL")
	}

	if params.ExcludeTrialAttendanced {
		query.Where("member_card_attendances.attendance_type != ?", membercardattendance.TrialAttendance)
	}

	if params.ExcludeAttendanceWithMemberCard {
		query.Where("member_card_attendances.attendance_type != ?", membercardattendance.WithMemberCard)
	}

	if params.ExcludeAttendanceWitouthMemberCard {
		query.Where("member_card_attendances.attendance_type != ?", membercardattendance.WithoutMemberCard)
	}

	if params.AttendedRange != nil {
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
	}

	return nil
}
