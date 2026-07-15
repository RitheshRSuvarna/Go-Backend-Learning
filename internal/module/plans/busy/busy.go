package busy

import (
	"strings"
	"time"
)

const (
	Low    = "low"
	Medium = "med"
	High   = "high"
)

func Label(categoryLabel, plannedArrival string) string {
	t, err := time.Parse("15:04", plannedArrival)
	if err != nil {
		return Low
	}

	minutes := t.Hour()*60 + t.Minute()

	switch strings.ToLower(strings.TrimSpace(categoryLabel)) {

	case "shopping":
		if between(minutes, 12, 0, 17, 0) {
			return High
		}

	case "sightseeing":
		if between(minutes, 11, 0, 15, 0) {
			return Medium
		}

	case "dining":
		if between(minutes, 12, 0, 14, 0) ||
			between(minutes, 19, 0, 21, 0) {
			return High
		}
	}

	return Low
}

func between(value, startHour, startMinute, endHour, endMinute int) bool {
	start := startHour*60 + startMinute
	end := endHour*60 + endMinute

	return value >= start && value <= end
}