package services

import (
	"common"
	"context"
	"plan/application/dto"
	"plan/domain/repository"
)

type ListPlanVerionservice struct {
	versionrepo repository.PlanVersionRepository
}

func NewListPlanVersionService(versionrepo repository.PlanVersionRepository) *ListPlanVerionservice {
	return &ListPlanVerionservice{versionrepo: versionrepo}
}

func (s *ListPlanVerionservice) ListVersion(ctx context.Context, id common.DaySessionID) ([]dto.PlanVersionDTO, error) {
	planversion, err := s.versionrepo.ListVerion(ctx, id)
	if err != nil {
		return nil, err
	}

	out := make([]dto.PlanVersionDTO, 0, len(planversion))
	for _, t := range planversion {
		out = append(out, dto.ToPlanVersionDTO(t))
	}
	return out, nil
}
