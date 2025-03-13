package membercardattendance

// Command to attend a course.
// If a member card identifier is provided, it means that a user will attend a course using this card.
// Otherwise, depending on the AttendaceType it could be a trial visit or a visit without a member card.
type CourseAttendanceCommand struct {
	MemberCardID   *uint  `json:"memberCard,omitempty"`
	CourseID       uint   `json:"course"`
	ParticipantID  uint   `json:"participant"`
	AttendanceType string `json:"attendanceType"`
}

// Entry that represent the course attendance with a member card.
// This includes the name of a course and time when the course was attended.
type MemberCardHistoryEntry struct {
	Course     string `json:"course"`
	AttendedAt string `json:"attendedAt"`
}
