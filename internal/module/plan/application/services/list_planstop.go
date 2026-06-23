package services 

import (
	"common"
	"context"
	"plan/application/dto"
	"plan/domain/repository"
)

type ListPlanStopService struct {
	planstoprepo repository.PlanStopRepository
}

func NewListPlanStopService(planstoprepo repository.PlanStopRepository) *ListPlanStopService {
	return &ListPlanStopService{planstoprepo: planstoprepo}
}

func (s *ListPlanStopService) ListPlanStop(ctx context.Context, id common.PlanVersionID) ([]dto.PlanStopDTO, error) {
	planstop, err := s.planstoprepo.ListStop(ctx, id)
	if err != nil {
		return nil, err
	}

	out := make([]dto.PlanStopDTO, 0, len(planstop))
	for _, t := range planstop {
		out = append(out, dto.ToPlanStopDTO(t))
	}
	return nil, err
}