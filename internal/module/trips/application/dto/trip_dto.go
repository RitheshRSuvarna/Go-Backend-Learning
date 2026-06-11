package dto

import (
	"time"
	"trip/domain/entity"
)

type TripDTO struct {
	ID             string
	Destination    string
	StartDate      string
	EndDate        string
	TravelersCount int
	CreatedAt      string
}

func ToTripDTO(t *entity.Trip) TripDTO {
	return TripDTO{
		ID:             t.ID().String(),
		Destination:    t.Destination(),
		StartDate:      t.StartDate(),
		EndDate:        t.EndDate(),
		TravelersCount: t.TravelersCount(),
		CreatedAt:      t.CreatedAt().Format(time.RFC3339),
	}
}
