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

func NewDaySessionListService(dayrepo repository.DaySessionRepository) *ListDaySessionService {
	return &ListDaySessionService{dayrepo: dayrepo}
}

func (d *ListDaySessionService) GetByID(ctx context.Context, id common.DaySessionID) (dto.DaySessionDTO, error) {
	daysession, err := d.dayrepo.GetByID(
		ctx, id,
	)
	if err != nil {
		return dto.DaySessionDTO{}, err
	}
	return dto.ToDaySessionDTO(daysession), nil
}
