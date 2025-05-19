package courses

import (
	"fit_and_roll/backend/common"
	"time"
)

// CourseDto represents the DTO for a course. It contains only the necessary fields for oveview.
type CourseDto struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Schedules   []ScheduleDto `json:"schedules"`
}

// CourseDetailsDto extened version of CourseDto when retrieven the details of a course.
type CourseDetailsDto struct {
	CourseDto
}

// ScheduleDto represents a DTO for a scheduled entry denoting when a course takes place.
type ScheduleDto struct {
	Day  string `json:"day"`
	Time string `json:"time"`
}

// Creates a new Course DTO from the course entity.
func NewCourseDto(course Course) *CourseDto {
	var schedulesDto []ScheduleDto
	for _, schedule := range course.Schedules {
		schedulesDto = append(schedulesDto, toScheduleDto(schedule))
	}

	return &CourseDto{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		Schedules:   schedulesDto,
	}
}

// Cretes a new Course Details DTO from the course entity.
func NewCourseDetailsDto(course Course) *CourseDetailsDto {
	var schedulesDto []ScheduleDto
	for _, schedule := range course.Schedules {
		schedulesDto = append(schedulesDto, toScheduleDto(schedule))
	}

	return &CourseDetailsDto{
		CourseDto: CourseDto{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			Schedules:   schedulesDto,
		},
	}
}

func toScheduleDto(schedule ScheduleEntry) ScheduleDto {
	return ScheduleDto{
		Day:  common.ToWeekDayString(schedule.Day, time.Monday),
		Time: common.ToTimeString(schedule.Time),
	}
}
