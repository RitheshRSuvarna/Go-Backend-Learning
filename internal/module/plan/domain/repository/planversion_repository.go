package repository

import (
	"common"
	"context"
	"plan/domain/entity"
)

type PlanVersionRepository interface {
	Create(ctx context.Context, planversion *entity.PlanVersion) error
	GetActivePlan(ctx context.Context, id common.DaySessionID) (*entity.PlanVersion, error)
	ListPlanVersion(ctx context.Context, id common.DaySessionID) ([]*entity.PlanVersion, error)
}
