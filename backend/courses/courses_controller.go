package courses

import (
	"fit_and_roll/backend/config"
	"fit_and_roll/backend/mappers"
	"fmt"
	"log"

	"fit_and_roll/backend/common"

	"gorm.io/gorm"
)

type CourseController struct {
	dbManager *config.DatabaseManager
}

func NewCoursesController(dbManager *config.DatabaseManager) *CourseController {
	return &CourseController{dbManager: dbManager}
}

func (controller *CourseController) FindCourses(page int, size int) (*common.Page[Course], error) {
	var courses []Course
	var total int64

	result := controller.dbManager.Paginate(page, size).Preload("Schedules").Where("archived = ?", false).Select("id, name").Find(&courses)

	if result.Error != nil {
		return nil, fmt.Errorf("error fetching courses: %v", result.Error)
	}

	if err := controller.dbManager.DB.Model(&Course{}).Where("archived = ?", false).Count(&total).Error; err != nil {
		return nil, fmt.Errorf("error to count the total number of courses: %v", err)
	}

	return &common.Page[Course]{
		Data:  courses,
		Total: int(total),
		Page:  page,
		Size:  size,
	}, nil
}

func (controller *CourseController) FindCourseDetails(id uint) (*Course, error) {
	var course Course
	result := controller.dbManager.DB.Preload("Schedules").First(&course, id)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch course details. Error: %v", result.Error)
	}

	return &course, nil
}

func (controller *CourseController) NewCourse(request CourseRequest) error {
	fmt.Println("In new course")
	fmt.Println(request)
	return controller.dbManager.DB.Transaction(func(tx *gorm.DB) error {
		course := Course{
			Name:        request.Name,
			Description: request.Description,
			Archived:    false,
		}

		result := tx.Create(&course)

		if result.Error != nil {
			log.Println("Failed to create a course", result.Error)
			return result.Error
		}

		var schedulePolicy []ScheduleEntry

		for _, schedule := range request.Schedules {
			scheduleEntry, err := toScheduleEntry(schedule)
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

func (controller *CourseController) UpdateCourse(id uint, request CourseRequest) error {
	return controller.dbManager.Transactional(func(tx *gorm.DB) error {
		course, err := controller.FindCourseDetails(id)
		if err != nil {
			return err
		}
		course.Name = request.Name
		course.Description = request.Description

		var schedulePolicy []ScheduleEntry
		for _, schedule := range request.Schedules {
			scheduleEntry, err := toScheduleEntry(schedule)
			if err != nil {
				return err
			}
			schedulePolicy = append(schedulePolicy, *scheduleEntry)
		}

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
	course, err := controller.FindCourseDetails(id)
	if err != nil {
		return err
	}

	course.Archived = true
	if err := controller.dbManager.DB.Save(&course).Error; err != nil {
		return err
	}
	return nil
}

func toScheduleEntry(schedule ScheduleEntryRequest) (*ScheduleEntry, error) {
	mappedDay, dayErr := mappers.ToWeekDay(schedule.Day)

	if dayErr != nil {
		return nil, dayErr
	}

	mappedTime, timeErr := mappers.ToTimeOnly(schedule.Time)
	if timeErr != nil {
		return nil, timeErr
	}

	return &ScheduleEntry{
		Day:  mappedDay,
		Time: mappedTime,
	}, nil
}
