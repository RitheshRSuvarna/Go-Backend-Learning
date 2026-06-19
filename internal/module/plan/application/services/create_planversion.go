package services

import (
	"context"
	"plan/domain/entity"
	"plan/domain/repository"
	"plan/application/command"
	"plan/application/dto"
)

type CreatePlanVersionService struct {
	versionrepo repository.PlanVersionRepository
}

func NewCreatePlanVersionService(versionrepo repository.PlanVersionRepository) *CreatePlanVersionService {
	return &CreatePlanVersionService{versionrepo: versionrepo}
}

func (s *CreatePlanVerisonService) CreatePlanVersion(ctx context.Context, cmd command.CreatePlanVersionCommand) (dto.PlanVersionDTO, error) {
	version, err := entity.NewPlanVersion()
}
