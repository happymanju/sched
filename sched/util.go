package sched

import (
	"strings"
	"time"
)

func ParseISO(date string) (time.Time, error) {
	yearIdx := strings.Index(date, "-")
	monthIdx := strings.Index(date[yearIdx+1:], "-")
	dayIdx := strings.Index(date[monthIdx+1:], "-")
	hourIdx := strings.Index(date[dayIdx+1:], "T")
	minIdx := strings.Index(date, ":")

}
