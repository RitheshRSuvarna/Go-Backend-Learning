package services

import (
	"common"
	"context"
	"day_session/application/dto"
	"day_session/domain/repository"
)

type ListDaySessionService struct {
	dayrepo repository.DaySessionRepository
}

func NewListDaySessionService(dayrepo repository.DaySessionRepository) *ListDaySessionService {
	return &ListDaySessionService{dayrepo: dayrepo}
}

func (d *ListDaySessionService) GetByTripIDAndDate(ctx context.Context, tripid common.TripID, date string) (dto.DaySessionDTO, error) {
	daysession, err := d.dayrepo.GetByTripIDAndDate(
		ctx, tripid, date,
	)
	if err != nil {
		return dto.DaySessionDTO{}, err
	}
	return dto.ToDaySessionDTO(daysession), nil
}
