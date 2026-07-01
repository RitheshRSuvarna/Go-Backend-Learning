package services

import (
	"common"
	"context"
	"plan/application/dto"
	"plan/domain/repository"
)

type ListStopByIDService struct {
	stoprepo repository.PlanStopRepository
}

func NewListStopByID(stoprepo repository.PlanStopRepository) *ListStopByIDService {
	return &ListStopByIDService{stoprepo: stoprepo}
}

func (s *ListStopByIDService) ListByID(ctx context.Context, id common.PlanStopID) (dto.PlanStopDTO, error) {
	planstop, err := s.stoprepo.GetPlanByID(ctx, id)
	if err != nil {
		return dto.PlanStopDTO{}, err
	}
	return dto.ToPlanStopDTO(planstop), nil
}
