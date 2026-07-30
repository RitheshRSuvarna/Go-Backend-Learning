package services

import (
	"common"
	"context"
	"encoding/json"
	"events/application/command"
	"events/application/dto"
	"events/domain/entity"
	"events/domain/repository"
	"fmt"
	plansentity "plans/domain/entity"
	plansrepository "plans/domain/repository"
)

type CreateEventService struct {
	repo         repository.EventsRepository
	planRepo     plansrepository.PlanVersionRepository
	planStopRepo plansrepository.PlanStopRepository
}

func NewCreateEventsService(repo repository.EventsRepository, planRepo plansrepository.PlanVersionRepository, planStopRepo plansrepository.PlanStopRepository) *CreateEventService {
	return &CreateEventService{
		repo:         repo,
		planRepo:     planRepo,
		planStopRepo: planStopRepo,
	}
}

func (e *CreateEventService) CreateEvents(ctx context.Context, cmd command.CreateEventsCommand) (dto.EventsDTO, error) {

	daySessionID, err := common.NewDaySessionID(cmd.DaySessionID)
	if err != nil {
		return dto.EventsDTO{}, err
	}

	activePlan, err := e.planRepo.GetActivePlan(ctx, daySessionID)
	if err != nil {
		return dto.EventsDTO{}, err
	}

	planStops, err := e.planStopRepo.ListStop(ctx, activePlan.ID())
	if err != nil {
		return dto.EventsDTO{}, err
	}

	events, err := e.repo.GetEvents(ctx, daySessionID)
	if err != nil {
		return dto.EventsDTO{}, err
	}

	var latestReached *entity.Events

	for _, event := range events {
		if event.EventType() == "Reached" {
			latestReached = event
			break
		}
	}

	if latestReached == nil {
		return dto.EventsDTO{}, fmt.Errorf("no reached event found")
	}

	planStopID, err := latestReached.PlanStopID()
	if err != nil {
		return dto.EventsDTO{}, err
	}

	var currentStop *plansentity.PlanStop

	for _, stop := range planStops {
		if stop.ID().String() == planStopID {
			currentStop = stop
			break
		}
	}

	if currentStop == nil {
		return dto.EventsDTO{}, fmt.Errorf("matching plan stop not found")
	}

	payload := entity.Payload{
		PlanStopID: currentStop.ID().String(),
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return dto.EventsDTO{}, err
	}

	event, err := entity.NewEvents(
		cmd.DaySessionID,
		cmd.EventType,
		payloadJSON,
	)
	if err != nil {
		return dto.EventsDTO{}, err
	}

	if err := e.repo.CreateEvents(ctx, event); err != nil {
		return dto.EventsDTO{}, err
	}

	return dto.ToEventDTO(event), nil
}
