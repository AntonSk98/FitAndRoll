package courses

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Schedules   []ScheduleEntry `gorm:"foreignKey:CourseID;onDelete:CASCADE" json:"schedules"`
	CreatedAt   time.Time
	Deleted     gorm.DeletedAt
}

type ScheduleEntry struct {
	ID       uint         `gorm:"primaryKey" json:"id"`
	Day      time.Weekday `json:"day"`
	Time     time.Time    `json:"time"`
	CourseID uint         `json:"course_id"`
}
