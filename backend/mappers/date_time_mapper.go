package mappers

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

func ToTimeString(input time.Time) string {
	return input.Format("15:04")
}

func ToWeekDayString(input time.Weekday) string {
	dayStr, ok := weekdayToString[input]
	if !ok {
		fmt.Printf("Warning: invalid weekday '%v', falling back to 'Monday'\n", input)
		return weekdayToString[time.Monday]
	}
	return dayStr
}

func ToTimeOnly(input string) (time.Time, error) {
	mappedTime, err := time.Parse("15:04", input)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %s", input)
	}
	return mappedTime, nil
}

func ToWeekDay(dayStr string) (time.Weekday, error) {
	day, ok := stringToWeekday[dayStr]
	if !ok {
		return -1, fmt.Errorf("invalid weekday string: %s", dayStr)
	}
	return day, nil
}
