package courses

import (
	"fit_and_roll/backend/common"
	"fmt"
)

// UpdateCourseRequest represents a request to update an existing course.
type UpdateCourseRequest struct {
	CourseRequest
	ID *uint `json:"id"`
}

// CourseRequest represents the structure of a course request, which includes
// details like the name, description, and schedules of the course.
type CourseRequest struct {
	Name        string                 `json:"name" validate:"required"`
	Description string                 `json:"description"`
	Schedules   []ScheduleEntryRequest `json:"schedules"`
}

// ScheduleEntryRequest represents an individual schedule entry for a course.
// It includes the day and time when a course takes place.
type ScheduleEntryRequest struct {
	Day  string `json:"day" validate:"required"`
	Time string `json:"time" validate:"required"`
}

// FindCourseParams contains the parameters used for searching courses.
type FindCourseParams struct {
	Name string `json:"name"`
}

// Checks if the fields of the CourseRequest are valid.
func (request *CourseRequest) Validate() error {
	if request.Name == "" {
		return fmt.Errorf("name is a required parameter")
	}
	return nil
}

// Checks if the fields of the UpdateCourseRequest are valid.
func (request *UpdateCourseRequest) Validate() error {
	// Validate fields inherited from CourseRequest
	if err := request.CourseRequest.Validate(); err != nil {
		return err
	}

	// Validate ID for UpdateCourseRequest
	if request.ID == nil {
		return fmt.Errorf("id is a required parameter for updating a course")
	}
	return nil
}

// Map converts the ScheduleEntryRequest to a ScheduleEntry object. It handles converting
func (scheduleEntryRequest *ScheduleEntryRequest) Map() (*ScheduleEntry, error) {
	mappedDay, dayErr := common.ToWeekDay(scheduleEntryRequest.Day)

	if dayErr != nil {
		return nil, dayErr
	}

	mappedTime, timeErr := common.ToTimeOnly(scheduleEntryRequest.Time)
	if timeErr != nil {
		return nil, timeErr
	}

	return &ScheduleEntry{
		Day:  mappedDay,
		Time: mappedTime,
	}, nil
}
