package repository

import (
	"context"
	"common"
	"trip/domain/entity"
)

type TripRepository interface {
	Create(ctx context.Context, trip *entity.Trip) error
	List(ctx context.Context) ([]*entity.Trip, error)
	GetByID(ctx context.Context, id common.TripID) (*entity.Trip, error)
}
