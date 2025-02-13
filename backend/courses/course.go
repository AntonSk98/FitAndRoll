package courses

import "time"

type Course struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Schedules   []ScheduleEntry `gorm:"foreignKey:CourseID;"`
	Archived    bool            `gorm:"default:false"`
}

type ScheduleEntry struct {
	ID       uint `gorm:"primaryKey"`
	Day      time.Weekday
	Time     time.Time
	CourseID uint
}
