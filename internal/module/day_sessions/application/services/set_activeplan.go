package services

import (
	"context"
	"common"
	"day_session/domain/repository"
)

type SetActivePlanService struct {
	setplanrepo repository.DaySessionRepository
}

func NewSetActivePlanService(setplanrepo repository.DaySessionRepository) *SetActivePlanService {
	return &SetActivePlanService{setplanrepo: setplanrepo}
}

func (s * SetActivePlanService) UpdateActivePlan(ctx context.Context, daysessionid common.DaySessionID, planversionid common.PlanVersionID) error {
	return s.setplanrepo.UpdateActivePlan(ctx, daysessionid, planversionid)
}