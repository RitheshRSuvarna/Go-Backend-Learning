package repository

import (
	"context"
	"common"
	"plan/domain/entity"
)

type PlanStopRepository interface{
	Create(ctx context.Context, planstop *entity.PlanStop) error
	ListStop(ctx context.Context, id common.PlanVersionID) ([]*entity.PlanStop, error)
	GetByID(ctx context.Context, planID common.PlanStopID) (*entity.PlanStop, error)
}