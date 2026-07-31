package services

import (
	"common"
	"context"
	daysessionRepo "day_session/domain/repository"
	"plans/application/dto"
	"plans/domain/repository"
)

type GetByIDPlanVersionService struct {
	versionrepo    repository.PlanVersionRepository
	daySessionRepo daysessionRepo.DaySessionRepository
}

func NewGetByIDPlanVersionService(versionrepo repository.PlanVersionRepository, daySessionRepo daysessionRepo.DaySessionRepository) *GetByIDPlanVersionService {
	return &GetByIDPlanVersionService{
		versionrepo:    versionrepo,
		daySessionRepo: daySessionRepo,
	}
}

func (s *GetByIDPlanVersionService) GetActivePlan(ctx context.Context, id common.DaySessionID) (dto.PlanVersionDTO, error) {

	daysession, err := s.daySessionRepo.GetDaySession(ctx, id)
	if err != nil {
		return dto.PlanVersionDTO{}, err
	}

	if daysession.ActivePlanVersionID() != nil {
		planversion, err := s.versionrepo.GetByID(ctx, *daysession.ActivePlanVersionID())
		if err != nil {
			return dto.PlanVersionDTO{}, err
		}
		return dto.ToPlanVersionDTO(planversion), nil
	}

	planversion, err := s.versionrepo.GetActivePlan(ctx, id)
	if err != nil {
		return dto.PlanVersionDTO{}, err
	}

	return dto.ToPlanVersionDTO(planversion), nil

}
