package command

type CreatePlanStopCommand struct {
	Position         int
	Title            string
	CategoryLabel    string
	ImageURL         string
	PlannedArrival   string
	PlannedDeparture string
	TravelMinutes    int
	StayMinutes      int
	BusyRiskLabel    string
}
