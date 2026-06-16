package services

import (
	"context"
	"common"
	"day_session/domain/repository"
	"day_session/application/command"
	"day_session/application/dto"
)

type ListDaySessionServiceID struct {
	dayrepo repository.DaySessionRepository
}

func NewListDaySessionService(dayrepo repository.DaySessionRepository) *ListDaySessionServiceID {
	return &ListDaySessionServiceID{dayrepo: dayrepo}
}

func (d* ListDaySessionServiceID) GetByIDAndDate(ctx context.Context, cmd command.CreateDaySessionCommand, tripid common.TripID) (dto.DaySessionDTO, error) {
	daysession, err := d.dayrepo.GetByTripIDAndDate(
		ctx, cmd.Date, tripid,
	)
	if err != nil {
		return dto.DaySessionDTO{}, err
	}
	return dto.ToDaySessionDTO(daysession), nil
}