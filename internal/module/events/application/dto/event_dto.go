package dto

import (
	"encoding/json"
	"events/domain/entity"
	"time"
)

type EventsDTO struct {
	ID           string
	DaySessionID string
	EventType    string
	Payload      json.RawMessage
	CreatedAt    string
}

func ToEventDTO(t *entity.Events) EventsDTO {
	return EventsDTO{
		ID:           t.ID().String(),
		DaySessionID: t.DaysessionID().String(),
		EventType:    t.EventType(),
		Payload:      t.Payload(),
		CreatedAt:    t.CreatedAt().Format(time.RFC3339),
	}
}
