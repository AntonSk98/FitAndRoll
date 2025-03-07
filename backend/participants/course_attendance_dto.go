package participants

type CourseAttendanceCommand struct {
	MemberCardID   *uint  `json:"memberCard,omitempty"`
	CourseID       uint   `json:"course"`
	ParticipantID  uint   `json:"participant"`
	AttendanceType string `json:"attendanceType"`
}
