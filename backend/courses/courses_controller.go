package courses

import (
	"fit_and_roll/backend/common"
	"fit_and_roll/backend/config"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type CourseController struct {
	dbManager *config.DatabaseManager
}

func NewCoursesController(dbManager *config.DatabaseManager) *CourseController {
	return &CourseController{dbManager: dbManager}
}

func (controller *CourseController) FindCourses(filter FindCourseParams, pageParams common.PageParams) (*common.Page[CourseDto], error) {
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

func (controller *CourseController) FindCourseDetails(id uint) (*CourseDetailsDto, error) {
	course, err := controller.findCourse(id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch course details. Error: %v", err)
	}

	return NewCourseDetailsDto(*course), nil
}

func (controller *CourseController) NewCourse(request CourseRequest) error {
	return controller.dbManager.Transactional(func(tx *gorm.DB) error {
		course := Course{
			Name:        request.Name,
			Description: request.Description,
		}

		result := tx.Create(&course)

		if result.Error != nil {
			log.Println("Failed to create a course", result.Error)
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

func (controller *CourseController) UpdateCourse(request UpdateCourseRequest) error {
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

func (controller *CourseController) ArchiveCourse(id uint) error {
	if err := controller.dbManager.DB.Delete(&Course{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (controller *CourseController) findCourse(id uint) (*Course, error) {
	var course Course
	result := controller.dbManager.DB.Preload("Schedules").First(&course, id)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch course details. Error: %v", result.Error)
	}

	return &course, nil
}
