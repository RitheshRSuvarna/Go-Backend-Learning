package llmclient

import (
	"fmt"
	"time"
)

func ValidateTime(value string) error {
	_, err := time.Parse("15:04", value)
	if err != nil {
		return fmt.Errorf("invalid time format: %s", value)
	}
	return nil
}

func ValidateStop(stop Stop) error {
	if stop.Title == "" {
		return fmt.Errorf("title is required")
	}

	if stop.CategoryLabel == "" {
		return fmt.Errorf("category_label is required")
	}

	if stop.PlannedArrival == "" {
		return fmt.Errorf("planned_arrival is required")
	}

	if stop.PlannedDeparture == "" {
		return fmt.Errorf("planned_departure is required")
	}

	if err := ValidateTime(stop.PlannedArrival); err != nil {
		return err
	}

	if err := ValidateTime(stop.PlannedDeparture); err != nil {
		return err
	}

	return nil
}

func ValidatePlan(resp PlanResponse) error {
	if len(resp.Stops) == 0 {
		return fmt.Errorf("plan contains no stops")
	}

	for _, stop := range resp.Stops {
		if err := ValidateStop(stop); err != nil {
			return err
		}
	}

	return nil
}