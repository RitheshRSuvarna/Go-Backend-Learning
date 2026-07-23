package services 

import (
	"common"
	"context"
	"plans/application/dto"
	"plans/domain/repository"
	"plans/busy"
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
	for _, stop := range planstop {
        dto := dto.ToPlanStopDTO(stop)

        dto.BusyRiskLabel = busy.Label(
            stop.CategoryLabel(),
            stop.PlannedArrival().Format("15:04"),
        )

        out = append(out, dto)
    }

	return out, nil
}