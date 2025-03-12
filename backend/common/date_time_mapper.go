package common

import (
	"fmt"
	"time"
)

const (
	sunday    = "sunday"
	monday    = "monday"
	tuesday   = "tuesday"
	wednesday = "wednesday"
	thursday  = "thursday"
	friday    = "friday"
	saturday  = "saturday"
)

var weekdayToString = map[time.Weekday]string{
	time.Sunday:    sunday,
	time.Monday:    monday,
	time.Tuesday:   tuesday,
	time.Wednesday: wednesday,
	time.Thursday:  thursday,
	time.Friday:    friday,
	time.Saturday:  saturday,
}

var stringToWeekday = map[string]time.Weekday{
	sunday:    time.Sunday,
	monday:    time.Monday,
	tuesday:   time.Tuesday,
	wednesday: time.Wednesday,
	thursday:  time.Thursday,
	friday:    time.Friday,
	saturday:  time.Saturday,
}

// ToTimeString converts a time.Time value to a string in "HH:MM" format.
func ToTimeString(input time.Time) string {
	return input.Format("15:04")
}

// ToWeekDayString converts a time.Weekday value to its string representation.
// If the provided weekday is invalid, it falls back to "Monday".
func ToWeekDayString(input time.Weekday) string {
	dayStr, ok := weekdayToString[input]
	if !ok {
		fmt.Printf("Warning: invalid weekday '%v', falling back to 'Monday'\n", input)
		return weekdayToString[time.Monday]
	}
	return dayStr
}

// ToTimeOnly parses a string in "HH:MM" format and returns a time.Time object.
// Returns an error if the format is invalid.
func ToTimeOnly(input string) (time.Time, error) {
	mappedTime, err := time.Parse("15:04", input)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %s", input)
	}
	return mappedTime, nil
}

// ToWeekDay converts a string representation of a weekday to a time.Weekday value.
// Returns an error if the input string is not a valid weekday.
func ToWeekDay(dayStr string) (time.Weekday, error) {
	day, ok := stringToWeekday[dayStr]
	if !ok {
		return -1, fmt.Errorf("invalid weekday string: %s", dayStr)
	}
	return day, nil
}

// ToDateString converts a time.Time value to a string in "YYYY-MM-DD" format.
func ToDateString(dateTime time.Time) string {
	return dateTime.Format(time.DateOnly)
}

// ToDateTimeString converts an RFC3339Nano formatted string to "YYYY-MM-DD HH:MM" format.
// Returns an error if parsing fails.
func ToDateTimeString(dateWithTimezoneString string) (string, error) {
	dateWithTimezone, err := time.Parse(time.RFC3339Nano, dateWithTimezoneString)
	if err != nil {
		return "", err
	}

	return dateWithTimezone.Format("2006-01-02 15:04"), nil
}
