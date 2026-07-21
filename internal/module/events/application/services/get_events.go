package services

import (
	"context"
	"common"
	"events/application/dto"
	"events/domain/repository"
)

type GetEventService struct {
	repo repository.EventsRepository
}

func NewGetEventService(repo repository.EventsRepository) *GetEventService {
	return &GetEventService{repo: repo}
}

func (e *GetEventService) GetEvents(ctx context.Context, id common.DaySessionID) ([]dto.EventsDTO, error) {
	// did, err := common.NewDaySessionID(id)
	// if err != nil {
	// 	return nil, common.NewValidationError("Invalid daysessionid", err)
	// }
	events, err := e.repo.GetEvents(ctx, id)
	if err != nil {
		return nil, err
	}

	out := make([]dto.EventsDTO, 0, len(events))
	for _, e := range events {
		out = append(out, dto.ToEventDTO(e))
	}
	return out, nil
}
