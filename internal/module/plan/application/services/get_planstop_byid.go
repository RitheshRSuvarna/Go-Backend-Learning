package services

import (
	"common"
	"context"
	"plan/application/dto"
	"plan/domain/repository"
)

type GetStopByIDService struct {
	stoprepo repository.PlanStopRepository
}

func NewGetStopByIDService(stoprepo repository.PlanStopRepository) *GetStopByIDService {
	return &GetStopByIDService{stoprepo: stoprepo}
}

func (s *GetStopByIDService) GetByID(ctx context.Context, id common.PlanStopID) (dto.PlanStopDTO, error) {
	planstop, err := s.stoprepo.GetPlanByID(ctx, id)
	if err != nil {
		return dto.PlanStopDTO{}, err
	}
	return dto.ToPlanStopDTO(planstop), nil
}
