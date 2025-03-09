package attendancehistory

import (
	"fmt"
	"time"
)

type AttendanceType string

const (
	WithMemberCard    AttendanceType = "WITH_MEMBER_CARD"
	TrialAttendance   AttendanceType = "TRIAL_ATTENDANCE"
	WithoutMemberCard AttendanceType = "WITHOUT_MEMBER_CARD"
)

type CourseAttendance struct {
	ID             uint `gorm:"primaryKey"`
	MemberCardID   *uint
	CourseID       uint `gorm:"not null"`
	ParticipantID  uint `gorm:"not null"`
	CreatedAt      time.Time
	AttendanceType AttendanceType `gorm:"type:varchar(20);not null"`
}

func NewCourseAttendance(courseAttendance CourseAttendanceCommand) (*CourseAttendance, error) {
	mappedAttendanceType, error := mapToAttendanceType(courseAttendance.AttendanceType)
	if error != nil {
		return nil, error
	}

	return &CourseAttendance{
		MemberCardID:   courseAttendance.MemberCardID,
		CourseID:       courseAttendance.CourseID,
		ParticipantID:  courseAttendance.ParticipantID,
		AttendanceType: mappedAttendanceType,
	}, nil
}

func mapToAttendanceType(value string) (AttendanceType, error) {
	switch value {
	case string(WithMemberCard):
		return WithMemberCard, nil
	case string(TrialAttendance):
		return TrialAttendance, nil
	case string(WithoutMemberCard):
		return WithoutMemberCard, nil
	default:
		return "", fmt.Errorf("invalid attendance type %v", value)
	}
}
