package port

import (
	"context"
	"time"
)

type PlanRequest struct {
    DaySessionID string
    Destination  string
    Date         string
    StartTime    string
    StartLabel   string
}

type PlanStop struct {
	Position         int
	Title            string
	CategoryLabel    string
	ImageURL         string
	PlannedArrival   time.Time
	PlannedDeparture time.Time
	TravelMinutes    int
	StayMinutes      int
}

type PlanResponse struct {
	Stops []PlanStop
}

type PlanGenerator interface {
	GeneratePlan(ctx context.Context, request PlanRequest) (PlanResponse, error)
}
