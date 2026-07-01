package repository

import (
	"common"
	"context"
	"plan/domain/entity"
)

type PlanVersionRepository interface {
	Create(ctx context.Context, planversion *entity.PlanVersion) error
	GetVersionByID(ctx context.Context, versionID common.PlanVersionID) (*entity.PlanVersion, error)
	ListPlanVersion(ctx context.Context, id common.DaySessionID) ([]*entity.PlanVersion, error)
}
