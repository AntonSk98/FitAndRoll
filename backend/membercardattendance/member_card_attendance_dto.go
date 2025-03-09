package membercardattendance

type CourseAttendanceCommand struct {
	MemberCardID   *uint  `json:"memberCard,omitempty"`
	CourseID       uint   `json:"course"`
	ParticipantID  uint   `json:"participant"`
	AttendanceType string `json:"attendanceType"`
}

type MemberCardHistoryEntry struct {
	Course     string `json:"course"`
	AttendedAt string `json:"attendedAt"`
}
