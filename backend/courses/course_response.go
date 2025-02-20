package courses

import "fit_and_roll/backend/mappers"

type CourseDto struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Schedules   []ScheduleDto `json:"schedules"`
}

type ScheduleDto struct {
	Day  string `json:"day"`
	Time string `json:"time"`
}

func NewCourseDto(course Course) *CourseDto {
	var schedulesDto []ScheduleDto
	for _, schedule := range course.Schedules {
		schedulesDto = append(schedulesDto, ScheduleDto{
			Day:  mappers.ToWeekDayString(schedule.Day),
			Time: mappers.ToTimeString(schedule.Time),
		})
	}

	return &CourseDto{
		Name:        course.Name,
		Description: course.Description,
		Schedules:   schedulesDto,
	}
}
