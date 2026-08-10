package llmclient

type DaySessionRequest struct {
    Destination string `json:"destination"`
    StartDate   string `json:"start_date"`
    EndDate     string `json:"end_date"`
    Travelers   int    `json:"travelers"`
}

type GeneratedDaySession struct {
    Date       string `json:"date"`
    StartTime  string `json:"start_time"`
    StartLabel string `json:"start_label"`
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

type ReplanResponse struct {
	Stops []Stop `json:"stops"`
}