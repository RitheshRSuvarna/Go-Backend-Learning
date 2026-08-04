package dto

import (
	"plans/domain/entity"
	"time"
)

// PlanVersionDTO represents a plan version returned by the API.
type PlanVersionDTO struct {
	ID           string
	DaysessionID string
	Version      int
	Note         string
	CreatedAt    string
}

func ToPlanVersionDTO(p *entity.PlanVersion) PlanVersionDTO {
	return PlanVersionDTO{
		ID:           p.ID().String(),
		DaysessionID: p.DaySessionID().String(),
		Version:      p.Version(),
		Note:         p.Note(),
		CreatedAt:    p.CreatedAt().Format(time.RFC3339),
	}
}
