package llmclient

import (
	"fmt"
	"regexp"
)

var hhmmPattern = regexp.MustCompile(`^(?:[01]\d|2[0-3]):[0-5]\d$`)

func ValidateStop(stop Stop) error {
	if stop.Position <= 0 {
		return fmt.Errorf("position must be greater than zero")
	}
	if stop.Title == "" {
		return fmt.Errorf("title is required")
	}
	if stop.CategoryLabel == "" {
		return fmt.Errorf("category_label is required")
	}
	if stop.PlannedArrival == "" || !hhmmPattern.MatchString(stop.PlannedArrival) {
		return fmt.Errorf("planned_arrival must use HH:MM format")
	}
	if stop.PlannedDeparture == "" || !hhmmPattern.MatchString(stop.PlannedDeparture) {
		return fmt.Errorf("planned_departure must use HH:MM format")
	}
	if stop.TravelMinutes < 0 {
		return fmt.Errorf("travel_minutes must not be negative")
	}
	if stop.StayMinutes <= 0 {
		return fmt.Errorf("stay_minutes must be greater than zero")
	}
	return nil
}
