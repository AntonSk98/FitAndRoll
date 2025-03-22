package common

import "time"

// Encapsulates parameter for time range filter
type TimeRangeFilter struct {
	From *time.Time `json:"from,omitempty"`
	To   *time.Time `json:"to,omitempty"`
}
