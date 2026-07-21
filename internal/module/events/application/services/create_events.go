package services

import (
	"context"
	"events/domain/entity"
	"events/domain/repository"
	"events/application/dto"
	"events/application/command"
)

type CreateEventService struct{
	repo repository.EventsRepository
}

func NewCreateEventsService(repo repository.EventsRepository) *CreateEventService {
	return &CreateEventService{repo: repo}
}

func (e *CreateEventService) CreateEvents(ctx context.Context, cmd command.CreateEventsCommand) (dto.EventsDTO, error) {
	event, err := entity.NewEvents(cmd.DaySessionID, cmd.EventType, cmd.Payload,)
	if err != nil {
		return dto.EventsDTO{}, err
	}
	
	if err := e.repo.CreateEvents(ctx, event); err != nil {
		return dto.EventsDTO{}, err
	}
	return dto.ToEventDTO(event), nil
}