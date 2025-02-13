package courses

import "time"

type CourseRequest struct {
	Name        string                 `json:"name" validate:"required"`
	Description string                 `json:"description"`
	Schedules   []ScheduleEntryRequest `json:"schedules"`
}

type ScheduleEntryRequest struct {
	Day  time.Weekday `json:"day" validate:"required"`
	Time time.Time    `json:"time" validate:"required"`
}
