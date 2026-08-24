package services

import (
	"common"
	"context"
	"fmt"
	busy "plans/busy"

	"day_session/domain/port"
	"day_session/domain/repository"
	"plans/domain/entity"
	planrepo "plans/domain/repository"
	triprepository "trip/domain/repository"
)

type GenerateLLMPlanService struct {
	planGenerator   port.PlanGenerator
	daySessionRepo  repository.DaySessionRepository
	tripRepo        triprepository.TripRepository
	planVersionRepo planrepo.PlanVersionRepository
	planStopRepo    planrepo.PlanStopRepository
}

func NewGenerateLLMPlanService(
	planGenerator port.PlanGenerator,
	daySessionRepo repository.DaySessionRepository,
	tripRepo triprepository.TripRepository,
	planVersionRepo planrepo.PlanVersionRepository,
	planStopRepo planrepo.PlanStopRepository,
) *GenerateLLMPlanService {
	return &GenerateLLMPlanService{
		planGenerator:   planGenerator,
		daySessionRepo:  daySessionRepo,
		tripRepo:        tripRepo,
		planVersionRepo: planVersionRepo,
		planStopRepo:    planStopRepo,
	}
}

func (s *GenerateLLMPlanService) GeneratePlan(
	ctx context.Context,
	daySessionID string,
) (port.PlanResponse, error) {

	if daySessionID == "" {
		return port.PlanResponse{}, fmt.Errorf(
			"day session id is required",
		)
	}

	id, err := common.NewDaySessionID(daySessionID)
	if err != nil {
		return port.PlanResponse{}, err
	}

	daySession, err := s.daySessionRepo.GetDaySession(ctx, id)
	if err != nil {
		return port.PlanResponse{}, err
	}

	trip, err := s.tripRepo.GetByID(
		ctx,
		daySession.TripID(),
	)
	if err != nil {
		return port.PlanResponse{}, err
	}

	plan, err := s.planGenerator.GeneratePlan(
		ctx,
		port.PlanRequest{
			DaySessionID: daySessionID,
			Destination:  trip.Destination(),
			Date:         daySession.Date(),
			StartTime:    daySession.STime(),
			StartLabel:   daySession.Label(),
		},
	)
	if err != nil {
		return port.PlanResponse{}, err
	}

	latestVersion, err := s.planVersionRepo.GetLatestVersion(ctx, id)
	if err != nil {
    	return port.PlanResponse{}, fmt.Errorf( "failed to get latest plan version: %w", err)
	}

	nextVersion := latestVersion + 1

	version, err := entity.NewPlanVersion(id, nextVersion, "Initial version")
	if err != nil {
		return port.PlanResponse{}, err
	}

	if err := s.planVersionRepo.Create(ctx, version); err != nil {
		return port.PlanResponse{}, err
	}

	// 3. Save every AI-generated stop
	for _, aiStop := range plan.Stops {

		busyrisk := busy.Label(aiStop.CategoryLabel, aiStop.PlannedArrival)
		planStop, err := entity.NewPlanStop(
			version.ID(),
			aiStop.Position,
			aiStop.Title,
			aiStop.CategoryLabel,
			aiStop.ImageURL,
			aiStop.PlannedArrival,
			aiStop.PlannedDeparture,
			aiStop.TravelMinutes,
			aiStop.StayMinutes,
			busyrisk,
		)
		if err != nil {
			return port.PlanResponse{}, err
		}

		if err := s.planStopRepo.Create(ctx, planStop); err != nil {
			return port.PlanResponse{}, err
		}
	}
	return plan, nil
}
