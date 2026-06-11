package repository

import (
	"context"
	"trip/domain/entity"
)

type TripRepository interface {
	Create(ctx context.Context, trip *entity.Trip) error
	List(ctx context.Context) ([]*entity.Trip, error)
}
