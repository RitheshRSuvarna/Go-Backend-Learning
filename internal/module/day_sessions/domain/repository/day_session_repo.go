package repository

import (
	"common"
	"context"
	"day_session/domain/entity"
)

type DaySessionRepository interface {
	CreateDaySession(ctx context.Context, daysession *entity.DaySession) error
	GetDaySession(ctx context.Context, id common.DaySessionID ) (*entity.DaySession, error)
	ListDaySession(ctx context.Context, id common.TripID) ([]*entity.DaySession, error)
	UpdateActivePlan(ctx context.Context, daySessionID common.DaySessionID, planVersionID common.PlanVersionID,) error
}
