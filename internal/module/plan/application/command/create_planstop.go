package command

import "time"

type CreatePlanStopCommand struct {
	Position         int
	Title            string
	CategoryLabel    string
	ImageURL         string
	PlannedArrival   time.Time
	PlannedDeparture time.Time
	TravelMinutes    int
	StayMinutes      int
	BusyRiskLabel    string
}
