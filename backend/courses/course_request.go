package courses

import "fmt"

type UpdateCourseRequest struct {
	CourseRequest
	ID *uint `json:"id"`
}

type CourseRequest struct {
	Name        string                 `json:"name" validate:"required"`
	Description string                 `json:"description"`
	Schedules   []ScheduleEntryRequest `json:"schedules"`
}

type ScheduleEntryRequest struct {
	Day  string `json:"day" validate:"required"`
	Time string `json:"time" validate:"required"`
}

type FindCourseParams struct {
	Name string `json:"name"`
}

func (request *CourseRequest) Validate() error {
	if request.Name == "" {
		return fmt.Errorf("name is a required parameter")
	}
	return nil
}

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
