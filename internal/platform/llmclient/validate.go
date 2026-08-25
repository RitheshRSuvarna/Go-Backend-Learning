package llmclient

import (
	"fmt"
	// "time"
)



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
	if stop.PlannedArrival.IsZero() {
    	return fmt.Errorf("planned_arrival is required")
	}	
	if stop.PlannedDeparture.IsZero() {
    	return fmt.Errorf("planned_departure is required")
	}
	if stop.TravelMinutes < 0 {
		return fmt.Errorf("travel_minutes must not be negative")
	}
	if stop.StayMinutes <= 0 {
		return fmt.Errorf("stay_minutes must be greater than zero")
	}
	return nil
}
