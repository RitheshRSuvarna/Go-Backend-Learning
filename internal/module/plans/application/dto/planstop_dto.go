package dto

import (
	"plans/domain/entity"
	"time"
)

type PlanStopDTO struct {
	PlanversionID    string
	Position         int
	Title            string
	CategoryLabel    string
	ImageURL         string
	PlannedArrival   string
	PlannedDeparture string
	TravelMinutes    int
	StayMinutes      int
	BusyRiskLabel    string
	CreatedAt        string
}

func ToPlanStopDTO(p *entity.PlanStop) PlanStopDTO {
	return PlanStopDTO{
		PlanversionID:    p.PlanVersionID().String(),
		Position:         p.Position(),
		Title:            p.Title(),
		CategoryLabel:    p.CategoryLabel(),
		ImageURL:         p.URL(),
		PlannedArrival:   p.PlannedArrival().String(),
		PlannedDeparture: p.PlannedDeparture().String(),
		TravelMinutes:    p.TravelMinutes(),
		StayMinutes:      p.StayMinutes(),
		BusyRiskLabel:    p.BusyRiskLabel(),
		CreatedAt:        p.CreatedAt().Format(time.RFC3339),
	}
}
