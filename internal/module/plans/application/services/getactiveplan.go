package services

import (
	"common"
	"context"
	"plan/application/dto"
	"plan/domain/repository"
)

type GetByIDPlanVersionService struct {
	versionrepo repository.PlanVersionRepository
}

func NewGetByIDPlanVersionService(versionrepo repository.PlanVersionRepository) *GetByIDPlanVersionService {
	return &GetByIDPlanVersionService{versionrepo: versionrepo}
}

func (s *GetByIDPlanVersionService) GetActivePlan(ctx context.Context, id common.DaySessionID) (dto.PlanVersionDTO, error) {
	planversion, err := s.versionrepo.GetActivePlan(ctx, id)
	if err != nil {
		return dto.PlanVersionDTO{}, err
	}
	return dto.ToPlanVersionDTO(planversion), nil
}
