package llmclient

import (
	"fmt"

	"github.com/google/uuid"
)

type PlanRequest struct {
	DaySessionID string `json:"day_session_id"`
}

func (r PlanRequest) Validate() error {
	if r.DaySessionID == "" {
		return fmt.Errorf("day_session_id is required")
	}
	if _, err := uuid.Parse(r.DaySessionID); err != nil {
		return fmt.Errorf("day_session_id must be a valid UUID: %w", err)
	}
	return nil
}

type Stop struct {
	Position         int    `json:"position"`
	Title            string `json:"title"`
	CategoryLabel    string `json:"category_label"`
	ImageURL         string `json:"image_url"`
	PlannedArrival   string `json:"planned_arrival"`
	PlannedDeparture string `json:"planned_departure"`
	TravelMinutes    int    `json:"travel_minutes"`
	StayMinutes      int    `json:"stay_minutes"`
}

type PlanResponse struct {
	Stops []Stop `json:"stops"`
}

func (r PlanResponse) Validate() error {
	if len(r.Stops) == 0 {
		return fmt.Errorf("stops must contain at least one item")
	}
	for i, stop := range r.Stops {
		if err := ValidateStop(stop); err != nil {
			return fmt.Errorf("stops[%d]: %w", i, err)
		}
	}
	return nil
}
