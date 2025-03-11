package courseattendance

type CourseAttendanceParameters struct {
	CourseID                           *uint   `json:"courseId,omitempty"`
	FullNameLike                       *string `json:"fullname,omitempty"`
	CourseLike                         *string `json:"course,omitempty"`
	ExcludeArchivedCourses             bool    `json:"excludeArchivedCourse"`
	ExcludeTrialAttendanced            bool    `json:"excludeTrialAttendance"`
	ExcludeAttendanceWitouthMemberCard bool    `json:"excludeNoMemberCard"`
	ExcludeAttendanceWithMemberCard    bool    `json:"excludeWithMemberCard"`
}

type CourseAttendanceDto struct {
	FullName       string `json:"fullname"`
	Course         string `json:"course"`
	AttendedAt     string `json:"attendedAt"`
	AttendanceType string `json:"attendanceType"`
}
