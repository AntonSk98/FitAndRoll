package courses

import "fit_and_roll/backend/mappers"

type CourseDto struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Schedules   []ScheduleDto `json:"schedules"`
}

type CourseDetailsDto struct {
	CourseDto
}

type ScheduleDto struct {
	Day  string `json:"day"`
	Time string `json:"time"`
}

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

func NewCourseDetailsDto(course Course) *CourseDetailsDto {
	var schedulesDto []ScheduleDto
	for _, schedule := range course.Schedules {
		schedulesDto = append(schedulesDto, toScheduleDto(schedule))
	}

	return &CourseDetailsDto{
		CourseDto: CourseDto{
			Name:        course.Name,
			Description: course.Description,
			Schedules:   schedulesDto,
		},
	}
}

func toScheduleDto(schedule ScheduleEntry) ScheduleDto {
	return ScheduleDto{
		Day:  mappers.ToWeekDayString(schedule.Day),
		Time: mappers.ToTimeString(schedule.Time),
	}
}
