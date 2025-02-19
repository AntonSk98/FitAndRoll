package mappers

import (
	"fmt"
	"time"
)

func ToTimeOnly(input string) (time.Time, error) {
	mappedTime, err := time.Parse("15:04", input)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return time.Time{}, fmt.Errorf("invalid time format: %s", input)
	}
	return mappedTime, nil
}

func ToWeekDay(day string) (time.Weekday, error) {
	return time.Sunday, nil
}
