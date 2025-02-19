package courses

type CourseRequest struct {
	Name        string                 `json:"name" validate:"required"`
	Description string                 `json:"description"`
	Schedules   []ScheduleEntryRequest `json:"schedules"`
}

type ScheduleEntryRequest struct {
	Day  string `json:"day" validate:"required"`
	Time string `json:"time" validate:"required"`
}
