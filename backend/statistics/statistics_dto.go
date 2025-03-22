package statistics

import "fit_and_roll/backend/common"

// Encapsulates filter parameters for statistics
type StatisticsParams struct {
	Name          *string                 `json:"name,omitempty"`
	AttendedRange *common.TimeRangeFilter `json:"attendedRange,omitempty"`
}

// Represents statistics DTO
type StatisticsDto struct {
	Name              string `json:"name"`
	WithMemberCard    uint   `json:"withMemberCard"`
	TrialAttendance   uint   `json:"trialAttendance"`
	WithoutMemberCard uint   `json:"withoutMemberCard"`
}
