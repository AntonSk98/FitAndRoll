package courseattendance

import (
	"fit_and_roll/backend/common"
)

// CourseAttendanceParameters represents the filters used when querying for course attendance history.
type CourseAttendanceParameters struct {
	CourseID                           *uint                   `json:"courseId,omitempty"`
	FullNameLike                       *string                 `json:"fullname,omitempty"`
	CourseLike                         *string                 `json:"course,omitempty"`
	ExcludeArchivedCourses             bool                    `json:"excludeArchivedCourse"`
	ExcludeTrialAttendanced            bool                    `json:"excludeTrialAttendance"`
	ExcludeAttendanceWitouthMemberCard bool                    `json:"excludeNoMemberCard"`
	ExcludeAttendanceWithMemberCard    bool                    `json:"excludeWithMemberCard"`
	AttendedRange                      *common.TimeRangeFilter `json:"attendedRange,omitempty"`
}

// CourseAttendanceDto represents a single record in the course attendance history.
type CourseAttendanceDto struct {
	ID             uint   `json:"id"`
	FullName       string `json:"fullname"`
	Course         string `json:"course"`
	AttendedAt     string `json:"attendedAt"`
	AttendanceType string `json:"attendanceType"`
}
