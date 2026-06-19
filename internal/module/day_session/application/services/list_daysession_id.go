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

func (d *ListDaySessionServiceID) GetByID(ctx context.Context, id common.DaySessionID) (dto.DaySessionDTO, error) {
	daysession, err := d.dayrepo.GetByID(
		ctx, id,
	)
	if err != nil {
		return dto.DaySessionDTO{}, err
	}
	return dto.ToDaySessionDTO(daysession), nil
}
