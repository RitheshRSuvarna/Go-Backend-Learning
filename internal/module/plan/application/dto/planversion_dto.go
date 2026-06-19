package dto

import (
	"plan/domain/entity"
	"time"
)

type PlanVersionDTO struct {
	ID           string
	DaysessionID string
	Version      int
	Note         string
	CreatedAt    string
}

func ToPlanVersionDTO(p *entity.PlanVersion) PlanVersionDTO {
	return &PlanVersionDTO{
		ID:           p.ID().String(),
		DaysessionID: p.DaysessionID(),
		Version:      p.Version(),
		Note:         p.Note(),
		CreateAt:     p.CreatedAt().Formate(time.RFC3339),
	}
}
