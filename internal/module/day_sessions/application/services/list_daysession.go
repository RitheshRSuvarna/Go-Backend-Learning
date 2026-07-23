package services

import (
	"common"
	"context"
	"day_session/application/dto"
	"encoding/json"
	evententity "events/domain/entity"
	"fmt"
	planentity "plans/domain/entity"

	daysessionrepo "day_session/domain/repository"
	"day_session/drift"
	eventrepo "events/domain/repository"
	planrepo "plans/domain/repository"
)

type ListDaySessionService struct {
	dayRepo      daysessionrepo.DaySessionRepository
	planRepo     planrepo.PlanVersionRepository
	planStopRepo planrepo.PlanStopRepository
	eventRepo    eventrepo.EventsRepository
}

func NewListDaySessionService(dayRepo daysessionrepo.DaySessionRepository,
	planRepo planrepo.PlanVersionRepository, planStopRepo planrepo.PlanStopRepository,
	eventRepo eventrepo.EventsRepository) *ListDaySessionService {
	return &ListDaySessionService{
		dayRepo:      dayRepo,
		planRepo:     planRepo,
		planStopRepo: planStopRepo,
		eventRepo:    eventRepo,
	}
}

func (d *ListDaySessionService) GetByTripIDAndDate(ctx context.Context, tripid common.TripID, date string) (dto.DaySessionResponseDTO, error) {
	daysession, err := d.dayRepo.GetByTripIDAndDate(ctx, tripid, date)
	if err != nil {
		return dto.DaySessionResponseDTO{}, err
	}

	activePlan, err := d.planRepo.GetActivePlan(ctx, daysession.ID())
	if err != nil {
		return dto.DaySessionResponseDTO{}, err
	}

	planStops, err := d.planStopRepo.ListStop(ctx, activePlan.ID())
	if err != nil {
		return dto.DaySessionResponseDTO{}, err
	}

	events, err := d.eventRepo.GetEvents(ctx, daysession.ID())
	if err != nil {
		return dto.DaySessionResponseDTO{}, err
	}

	var latestEvent *evententity.Events

	for _, event := range events {
		if event.EventType() != "Reached" {
			continue
		}

		if latestEvent == nil || event.TS().Time().After(latestEvent.TS().Time()) {
			latestEvent = event
		}
	}

	if latestEvent == nil {
		return dto.DaySessionResponseDTO{}, fmt.Errorf("no reached event found")
	}
	var matchedStop *planentity.PlanStop
	var payload map[string]any

	err = json.Unmarshal(latestEvent.Payload(), &payload)
	if err != nil {
		return dto.DaySessionResponseDTO{}, err
	}

	planStopID, ok := payload["plan_stop_id"].(string)
	if !ok {
		return dto.DaySessionResponseDTO{}, fmt.Errorf("plan_stop_id not found in payload")
	}

	for _, stop := range planStops {
		if stop.ID().Value() == planStopID {
			matchedStop = stop
			break
		}
	}
	// for _, stop := range planStops {
	// 	if stop.ID == latestEvent.stopID() {
	//     	matchedStop = stop
	//     	break
	// 	}
	// }

	plannedArrival := matchedStop.PlannedArrival()
	actualArrival := latestEvent.TS()

	status := drift.Compute(plannedArrival, actualArrival)

	return dto.DaySessionResponseDTO{
		DaySession: dto.ToDaySessionDTO(daysession),
		Status:     status,
	}, nil
}
