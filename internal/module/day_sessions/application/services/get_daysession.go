package services

import (
	"common"
	"context"
	"day_session/application/dto"

	evententity "events/domain/entity"
	"fmt"
	planentity "plans/domain/entity"

	daysessionrepo "day_session/domain/repository"
	"day_session/drift"
	eventrepo "events/domain/repository"
	planrepo "plans/domain/repository"
)

type GetDaySessionService struct {
	dayRepo      daysessionrepo.DaySessionRepository
	planRepo     planrepo.PlanVersionRepository
	planStopRepo planrepo.PlanStopRepository
	eventRepo    eventrepo.EventsRepository
}

func NewGetDaySessionService(dayRepo daysessionrepo.DaySessionRepository,
	planRepo planrepo.PlanVersionRepository, planStopRepo planrepo.PlanStopRepository,
	eventRepo eventrepo.EventsRepository) *GetDaySessionService {
	return &GetDaySessionService{
		dayRepo:      dayRepo,
		planRepo:     planRepo,
		planStopRepo: planStopRepo,
		eventRepo:    eventRepo,
	}
}

func (d *GetDaySessionService) GetDaySession(ctx context.Context, daysessionid common.DaySessionID) (dto.DaySessionResponseDTO, error) {
	daysession, err := d.dayRepo.GetDaySession(ctx, daysessionid)
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

	// var payload map[string]any

	planStopID, err := latestEvent.PlanStopID()
	if err != nil {
		return dto.DaySessionResponseDTO{}, err
	}

	var matchedStop *planentity.PlanStop

	for _, stop := range planStops {
		if stop == nil {
			continue
		}

		if stop.ID().Value() == planStopID {
			matchedStop = stop
			break
		}
	}

	if matchedStop == nil {
		return dto.DaySessionResponseDTO{}, fmt.Errorf("matching plan stop not found")
	}

	// err = json.Unmarshal(latestEvent.Payload(), &payload)
	// if err != nil {
	// 	return dto.DaySessionResponseDTO{}, err
	// }

	// planStopID, ok := payload["plan_stop_id"].(string)
	// if !ok {
	// 	return dto.DaySessionResponseDTO{}, fmt.Errorf("plan_stop_id not found in payload")
	// }

	// for _, stop := range planStops {
	// 	if stop.ID().Value() == planStopID {
	// 		matchedStop = stop
	// 		break
	// 	}
	// }

	// var planStopID string
	// planStopID, err := latestEvent.PlanStopID()
	// if err != nil {
	// 	return dto.DaySessionResponseDTO{}, err
	// }

	// for _, stop := range planStops {
	// 	if stop.ID == planStopID {
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
