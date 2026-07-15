package services

import (
	"common"
	"context"
	"day_session/application/dto"
	"day_session/domain/repository"
)

type ListDaySessionServiceID struct {
	dayrepo repository.DaySessionRepository
}

func NewDaySessionListService(dayrepo repository.DaySessionRepository) *ListDaySessionServiceID {
	return &ListDaySessionServiceID{dayrepo: dayrepo}
}

func (d *ListDaySessionServiceID) GetByID(ctx context.Context, id common.TripID) ([]dto.DaySessionDTO, error) {
	daysession, err := d.dayrepo.GetByID(
		ctx, id,
	)
	if err != nil {
		return nil, err
	}

	out := make([]dto.DaySessionDTO, 0, len(daysession))

	for _, ds := range daysession {
		out = append(out, dto.ToDaySessionDTO(ds))
	}
	return out, nil
}
