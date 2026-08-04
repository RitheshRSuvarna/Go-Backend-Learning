package dto

import (
	"day_session/domain/entity"
	"time"
)


// DaySessionDTO represents a day session returned by the API.
type DaySessionDTO struct {
	ID         string
	TripID     string
	Date       string
	StartTime  string
	StartLabel string
	CreatedAT  string
}

func ToDaySessionDTO(t *entity.DaySession) DaySessionDTO {
	return DaySessionDTO {
		ID:         t.ID().String(),
		TripID:     t.TripID().String(),
		Date:       t.Date(),
		StartTime:  t.STime(),
		StartLabel: t.Label(),
		CreatedAT:  t.CreatedAt().Format(time.RFC3339),
	}
}
