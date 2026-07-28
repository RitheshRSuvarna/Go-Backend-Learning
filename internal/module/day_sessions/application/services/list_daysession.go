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

func (d *ListDaySessionService) ListDaysession(ctx context.Context, id common.TripID) ([]dto.DaySessionDTO, error) {
	daysession, err := d.dayrepo.ListDaySession(
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
