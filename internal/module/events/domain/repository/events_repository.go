package repository

import (
	"common"
	"context"
	"events/domain/entity"
)

type EventsRepository interface{
	CreateEvents(ctx context.Context, events *entity.Events) error
	GetEvents(ctx context.Context, id common.DaySessionID) ([]*entity.Events, error)
}