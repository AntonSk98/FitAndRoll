package membercardattendance

import (
	"fmt"
	"time"
)

// Represent the type of course attendance
// There are three type how a course can be attended: with a member card, as a trial session, or without a member card
type AttendanceType string

// Constants that define possible course attendance types.
const (
	WithMemberCard    AttendanceType = "WITH_MEMBER_CARD"
	TrialAttendance   AttendanceType = "TRIAL_ATTENDANCE"
	WithoutMemberCard AttendanceType = "WITHOUT_MEMBER_CARD"
)

// Entity that represents the fact of attending a course by participant
// It is possible to attend a course without member card. In this case the id of the member card shall be null
type MemberCardAttendance struct {
	ID             uint `gorm:"primaryKey"`
	MemberCardID   *uint
	CourseID       uint `gorm:"not null"`
	ParticipantID  uint `gorm:"not null"`
	CreatedAt      time.Time
	AttendanceType AttendanceType `gorm:"type:varchar(20);not null"`
}

// Creates a new MemberCardAttendance
func NewCourseAttendance(courseAttendance CourseAttendanceCommand) (*MemberCardAttendance, error) {
	mappedAttendanceType, error := mapToAttendanceType(courseAttendance.AttendanceType)
	if error != nil {
		return nil, error
	}

	return &MemberCardAttendance{
		MemberCardID:   courseAttendance.MemberCardID,
		CourseID:       courseAttendance.CourseID,
		ParticipantID:  courseAttendance.ParticipantID,
		AttendanceType: mappedAttendanceType,
	}, nil
}

// Mapper that takes the attendance type as a string and attempts to map it one of the attendance type constants.
// If it cannot be mapped to any known attendance types, the error is returned.
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
