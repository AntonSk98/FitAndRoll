package courseattendance

import "time"

// CourseAttendanceParameters represents the filters used when querying for course attendance history.
type CourseAttendanceParameters struct {
	CourseID                           *uint          `json:"courseId,omitempty"`
	FullNameLike                       *string        `json:"fullname,omitempty"`
	CourseLike                         *string        `json:"course,omitempty"`
	ExcludeArchivedCourses             bool           `json:"excludeArchivedCourse"`
	ExcludeTrialAttendanced            bool           `json:"excludeTrialAttendance"`
	ExcludeAttendanceWitouthMemberCard bool           `json:"excludeNoMemberCard"`
	ExcludeAttendanceWithMemberCard    bool           `json:"excludeWithMemberCard"`
	AttendedRange                      *AttendedRange `json:"attendedRange,omitempty"`
}

type AttendedRange struct {
	From *time.Time `json:"from,omitempty"`
	To   *time.Time `json:"to,omitempty"`
}

// CourseAttendanceDto represents a single record in the course attendance history.
type CourseAttendanceDto struct {
	ID             uint   `json:"id"`
	FullName       string `json:"fullname"`
	Course         string `json:"course"`
	AttendedAt     string `json:"attendedAt"`
	AttendanceType string `json:"attendanceType"`
}
