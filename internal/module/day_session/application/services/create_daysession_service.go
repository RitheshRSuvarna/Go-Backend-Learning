package services

import (
	"context"
	"day_session/application/command"
	"day_session/application/dto"
	"day_session/domain/entity"
	"day_session/domain/repository"
)

type CreateDaySessionService struct {
	dayrepo repository.DaySessionRepository
}

func NewDaySessionService(dayrepo repository.DaySessionRepository) *CreateDaySessionService {
	return &CreateDaySessionService{dayrepo: dayrepo}
}

func (d *CreateDaySessionService) CreateDaySession(ctx context.Context, cmd command.CreateDaySessionCommand) (dto.DaySessionDTO, error) {
	daysession, err := entity.NewDaySession(cmd.TripID, cmd.Date, cmd.StartTime, cmd.StartLabel)
	if err != nil {
		return dto.DaySessionDTO{}, err
	}
	if err := d.dayrepo.Create(ctx, daysession); err != nil {
		return dto.DaySessionDTO{}, err
	}
	return dto.ToDaySessionDTO(daysession), nil
}
