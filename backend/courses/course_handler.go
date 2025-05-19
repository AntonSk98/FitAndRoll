package courses

import (
	"fit_and_roll/backend/common"
	"fit_and_roll/backend/config"
	"fmt"

	"gorm.io/gorm"
)

// CourseHandler manages courses.
type CourseHandler struct {
	dbManager *config.DatabaseManager
}

// NewCourseHandler creates a new instance of CourseHandler.
//
// Parameters:
//   - dbManager: A pointer to the DatabaseManager used for database interactions.
//
// Returns:
//   - A pointer to a new CourseHandler instance.
func NewCourseHandler(dbManager *config.DatabaseManager) *CourseHandler {
	return &CourseHandler{dbManager: dbManager}
}

// FindCourses retrieves a list of courses based on the provided filter parameters and pagination.
// It performs a query on the database to fetch courses matching the filter criteria and then returns
// the results as a paginated list of CourseDto objects.
//
// Parameters:
//   - filter: FindCourseParams struct containing filter criteria (e.g., course name).
//   - pageParams: Pagination parameters, including the page number and page size.
//
// Returns:
//   - A pointer to a common.Page containing a list of CourseDto objects, along with pagination info.
//   - An error if the operation fails.
func (controller *CourseHandler) FindCourses(filter FindCourseParams, pageParams common.PageParams) (*common.Page[CourseDto], error) {
	var coursesDto []CourseDto
	var courses []Course
	var total int64

	baseQuery := controller.dbManager.DB.Preload("Schedules").Order("created_at desc")

	if filter.Name != "" {
		baseQuery = baseQuery.Where("LOWER(name) LIKE LOWER(?)", "%"+filter.Name+"%")
	}

	if err := baseQuery.Model(&Course{}).Count(&total).Error; err != nil {
		return nil, fmt.Errorf("error to count the total number of courses: %v", err)
	}

	if err := baseQuery.Scopes(controller.dbManager.Paginate(pageParams.Page, pageParams.Size)).Find(&courses).Error; err != nil {
		return nil, fmt.Errorf("error fetching courses: %v", err)
	}

	for _, course := range courses {
		coursesDto = append(coursesDto, *NewCourseDto(course))
	}

	return &common.Page[CourseDto]{
		Data:  coursesDto,
		Total: int(total),
		Page:  pageParams.Page,
		Size:  pageParams.Size,
	}, nil
}

// FindCourseDetails retrieves the details of a specific course by its ID.
// It fetches the course, along with its associated schedule entries, and returns them as a CourseDetailsDto.
//
// Parameters:
//   - id: The unique identifier of the course to fetch.
//
// Returns:
//   - A pointer to a CourseDetailsDto containing the course details and schedule information.
//   - An error if the operation fails.
func (controller *CourseHandler) FindCourseDetails(id uint) (*CourseDetailsDto, error) {
	course, err := controller.findCourse(id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch course details. Error: %v", err)
	}

	return NewCourseDetailsDto(*course), nil
}

// NewCourse creates a new course and persists it in the database.
// The function handles creating the course and associating it with the provided schedule entries.
//
// Parameters:
//   - request: A CourseRequest struct containing the details of the course to create.
//
// Returns:
//   - An error if the operation fails.
func (controller *CourseHandler) NewCourse(request CourseRequest) error {
	return controller.dbManager.Transactional(func(tx *gorm.DB) error {
		course := Course{
			Name:        request.Name,
			Description: request.Description,
		}

		result := tx.Create(&course)

		if result.Error != nil {
			return result.Error
		}

		var schedulePolicy []ScheduleEntry

		for _, schedule := range request.Schedules {
			scheduleEntry, err := schedule.Map()
			if err != nil {
				return err
			}
			schedulePolicy = append(schedulePolicy, *scheduleEntry)
		}

		if err := tx.Model(&course).Association("Schedules").Append(schedulePolicy); err != nil {
			return err
		}
		return nil
	})
}

// UpdateCourse updates an existing course, including its schedule entries.
// It ensures that the provided request is valid before performing the update.
// It replaces the course's schedule entries with the new ones.
//
// Parameters:
//   - request: An UpdateCourseRequest struct containing the updated course details.
//
// Returns:
//   - An error if the operation fails.
func (controller *CourseHandler) UpdateCourse(request UpdateCourseRequest) error {
	return controller.dbManager.Transactional(func(tx *gorm.DB) error {
		if err := request.Validate(); err != nil {
			return err
		}

		course, err := controller.findCourse(*request.ID)
		if err != nil {
			return err
		}
		course.Name = request.Name
		course.Description = request.Description

		var schedulePolicy []ScheduleEntry
		for _, schedule := range request.Schedules {
			scheduleEntry, err := schedule.Map()
			if err != nil {
				return err
			}
			schedulePolicy = append(schedulePolicy, *scheduleEntry)
		}

		// Permanent Delete: Physically deletes the association records from the database.
		// Therefore, there are 2x 'Unscoped' used
		if err := tx.Unscoped().Model(&course).Association("Schedules").Unscoped().Replace(schedulePolicy); err != nil {
			return err
		}

		if err := tx.Save(&course).Error; err != nil {
			return err
		}
		return nil
	})
}

// ArchiveCourse archives (soft deletes) a course by its ID.
//
// Parameters:
//   - id: The unique identifier of the course to archive.
//
// Returns:
//   - An error if the operation fails.
func (controller *CourseHandler) ArchiveCourse(id uint) error {
	if err := controller.dbManager.DB.Delete(&Course{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (controller *CourseHandler) findCourse(id uint) (*Course, error) {
	var course Course
	result := controller.dbManager.DB.Preload("Schedules").First(&course, id)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch course details. Error: %v", result.Error)
	}

	return &course, nil
}
