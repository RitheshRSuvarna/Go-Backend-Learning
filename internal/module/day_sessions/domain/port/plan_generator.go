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

type RePlanRequest struct {
	DaySessionID string
	Destination  string
    Date         string
    StartTime    string
    StartLabel   string

    ExistingPlanStops []RePlanStop
}

type RePlanStop struct {
	Position         int
	Title            string
	CategoryLabel    string
	ImageURL         string
	PlannedArrival   time.Time
	PlannedDeparture time.Time
	TravelMinutes    int
	StayMinutes      int
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
	Replan(ctx context.Context, req RePlanRequest) (PlanResponse, error)
}
