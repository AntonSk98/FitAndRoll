package courses

import (
	"time"

	"gorm.io/gorm"
)

// Represents a course in the application.
//
// Fields:
//   - ID: uint, the unique identifier for the course (primary key).
//   - Name: string, the name of the course.
//   - Description: string, a description of the course.
//   - Schedules: []ScheduleEntry, a list of schedule entries for the course.
//     This list is managed by GORM, which uses the foreign key `CourseID`.
//   - CreatedAt: time.Time, the timestamp when the course was created.
//   - Deleted: gorm.DeletedAt, the soft delete marker (if the course is deleted).
type Course struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Schedules   []ScheduleEntry `gorm:"foreignKey:CourseID;onDelete:CASCADE" json:"schedules"`
	CreatedAt   time.Time
	Deleted     gorm.DeletedAt
}

// Represents a schedule entry for a course.
//
// Fields:
//   - ID: uint, the unique identifier for the schedule entry (primary key).
//   - Day: time.Weekday, the day of the week the course will take place (e.g., Monday, Tuesday).
//   - Time: time.Time, the time when the course will occur (e.g., "10:00" or "15:30").
//   - CourseID: uint, the ID of the course this schedule entry is associated with.
type ScheduleEntry struct {
	ID       uint         `gorm:"primaryKey" json:"id"`
	Day      time.Weekday `json:"day"`
	Time     time.Time    `json:"time"`
	CourseID uint         `json:"course_id"`
}
