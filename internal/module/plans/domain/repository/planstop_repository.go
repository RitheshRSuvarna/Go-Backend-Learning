package repository

import (
	"common"
	"context"
	"plans/domain/entity"
)

type PlanStopRepository interface {
	Create(ctx context.Context, planstop *entity.PlanStop) error
	ListStop(ctx context.Context, id common.PlanVersionID) ([]*entity.PlanStop, error)
}
