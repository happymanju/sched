package sched

import (
	"fmt"
	"strconv"
	"time"
)

func ParseISO(year string, month string, day string, hour string, minute string) (time.Time, error) {
	//todo

	parsedYear, _ := strconv.Atoi(year)
	parsedMonth, _ := strconv.Atoi(month)
	parsedDay, _ := strconv.Atoi(day)
	parsedHour, _ := strconv.Atoi(hour)
	parsedMin, _ := strconv.Atoi(minute)

	t := time.Date(parsedYear, time.Month(parsedMonth), parsedDay, parsedHour, parsedMin, 0, 0, time.Local)
	return t, nil
}

func MakeTimestamp() string {
	t := time.Now()
	year := strconv.Itoa(t.Year())
	month := strconv.Itoa(int(t.Month()))
	day := strconv.Itoa(t.Day())
	hour := strconv.Itoa(t.Hour())
	min := strconv.Itoa(t.Minute())

	if len(hour) < 2 {
		hour = "0" + hour
	}
	if len(min) < 2 {
		min = "0" + min
	}
	return fmt.Sprintf("%s%s%sT%s%s", year, month, day, hour, min)
}
