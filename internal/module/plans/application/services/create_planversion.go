package services

import (
	"common"
	"context"
	"plans/application/command"
	"plans/application/dto"
	"plans/domain/entity"
	"plans/domain/repository"
)

type CreatePlanVersionService struct {
	versionrepo repository.PlanVersionRepository
}

func NewCreatePlanVersionService(versionrepo repository.PlanVersionRepository) *CreatePlanVersionService {
	return &CreatePlanVersionService{versionrepo: versionrepo}
}

func (s *CreatePlanVersionService) CreatePlanVersion(ctx context.Context, id common.DaySessionID, cmd command.CreatePlanVersionCommand) (dto.PlanVersionDTO, error) {
	version, err := entity.NewPlanVersion(id, cmd.Version, cmd.Note)
	if err != nil {
		return dto.PlanVersionDTO{}, err
	}
	if err = s.versionrepo.Create(ctx, version); err != nil {
		return dto.PlanVersionDTO{}, err
	}
	return dto.ToPlanVersionDTO(version), nil
}
