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

func (s *GetByIDPlanVersionService) GetPlanVersionByID(ctx context.Context, id common.PlanVersionID) (dto.PlanVersionDTO, error) {
	planversion, err := s.versionrepo.GetVersionByID(ctx, id)
	if err != nil {
		return dto.PlanVersionDTO{}, err
	}
	return dto.ToPlanVersionDTO(planversion), nil
}
