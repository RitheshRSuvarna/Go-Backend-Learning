package port

import "context"

type DaySessionGenerator interface {
    Generate(ctx context.Context, input GenerateDaySessionsInput) ([]GeneratedDaySession, error)
}

type GenerateDaySessionsInput struct {
    TripID       string
    Destination  string
    StartDate    string
    EndDate      string
    Travelers    int
}

type GeneratedDaySession struct {
    Date       string `json:"date"`
    StartTime  string `json:"start_time"`
    StartLabel string `json:"start_label"`
}