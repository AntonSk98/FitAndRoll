package courseattendance

type CourseAttendanceDto struct {
	FullName       string `json:"fullname"`
	Course         string `json:"course"`
	ArchivedCourse bool   `json:"archivedCourse"`
	AttendedAt     string `json:"attendedAt"`
	AttendanceType string `json:"attendanceType"`
}
