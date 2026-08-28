package dto

type ItineraryDTO struct {
	Trip        TripDTO         `json:"trip"`
	DaySessions []DaySessionDTO `json:"day_sessions"`
}

type DaySessionDTO struct {
	ID           string           `json:"id"`
	Date         string           `json:"date"`
	PlanVersions []PlanVersionDTO `json:"plan_versions"`
}

type PlanVersionDTO struct {
	ID        string        `json:"id"`
	Version   int           `json:"version"`
	Notes     string        `json:"notes"`
	PlanStops []PlanStopDTO `json:"plan_stops"`
}

type PlanStopDTO struct {
	ID            string `json:"id"`
	Position      int    `json:"position"`
	Title         string `json:"title"`
	CategoryLabel string `json:"category_label"`
	ImageURL      string `json:"image_url"`
	TravelMinutes int    `json:"travel_minutes"`
	StayMinutes   int    `json:"stay_minutes"`
	BusyRiskLabel string `json:"busy_risk_label"`
}
