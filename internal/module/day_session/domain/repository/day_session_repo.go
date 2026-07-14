package repository

import (
	"common"
	"context"
	"day_session/domain/entity"
)

type DaySessionRepository interface {
	Create(ctx context.Context, daysession *entity.DaySession) error
	GetByTripIDAndDate(ctx context.Context, tripID common.TripID, Date string) (*entity.DaySession, error)
	GetByID(ctx context.Context, id common.TripID) ([]*entity.DaySession, error)
}
