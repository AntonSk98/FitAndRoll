package common

import (
	"fmt"
	"time"

	_ "time/tzdata"

	"gorm.io/gorm"
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
// If the provided weekday is invalid, it falls back to the passed fallback parameter.
func ToWeekDayString(input time.Weekday, fallback time.Weekday) string {
	dayStr, ok := weekdayToString[input]
	if !ok {
		return weekdayToString[fallback]
	}
	return dayStr
}

// ToDateOnly parses a string in "YYYY-MM-DD" format and returns a time.Time object
// Returns an error if the format is invalid
func ToDateOnly(input string) (time.Time, error) {
	mappedTime, err := time.Parse(time.DateOnly, input)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %s", input)
	}

	return mappedTime, nil
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

// ToDateStringOrEmpty converts a time.Time value to astring in "YYYY-MM-DD" format.
// If the passed time is empty the empty string is returned.
func ToDateStringOrEmpty(dateTime *time.Time) string {
	if dateTime == nil {
		return ""
	}

	return ToDateString(*dateTime)
}

// Returns a string if it is not nil.
// Otherwise an empty string is returned.
func EmptyIfNil(str *string) string {
	if str == nil {
		return ""
	}

	return *str
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

// ToDateTime converts a time.Time to "YYYY-MM-DD HH:MM" format.
func ToDateTime(timestamp time.Time) string {
	return timestamp.Format(time.DateTime)
}

// FormatDeletedAt converts gorm.DeletedAt to a string (formatted datetime or empty if invalid).
func FormatDeletedAt(optionalTimestamp gorm.DeletedAt) string {
	if optionalTimestamp.Valid {
		return ToDateTime(optionalTimestamp.Time) // Assuming ToDateTime is already in common
	}
	return ""
}

// BoolToString returns "1" is the passed boolean is true, "0" otherwise.
func BoolToString(boolean bool) string {
	if boolean {
		return "1"
	}

	return "0"
}

func ConvertUTCToBerlin(utcTime time.Time) (time.Time, error) {
	berlinLocation, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to load Europe/Berlin location: %v", err)
	}

	berlinTime := utcTime.In(berlinLocation)

	return berlinTime, nil
}
