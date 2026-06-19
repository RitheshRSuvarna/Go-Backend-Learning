package repository

import (
	"context"
	"common"
	"plan/domain/entity"
)

type PlanVersionRepository interface{
	Create(ctx context.Context, planversion *entity.PlanVersion) error
	ListVerion(ctx context.Context) ([]*entity.PlanVersion, error)
	GetVersion(ctx context.Context, versionID common.PlanVersionID) (*entity.PlanVersion, error)
}