package services

import (
	"common"
	"context"
	"errors"
	"fmt"
	busy "plans/busy"

	"github.com/jackc/pgx/v5"

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
		if errors.Is(err, pgx.ErrNoRows) {
			// No plan exists yet.
			// This is the first plan.
			latestVersion = 0
		} else {
			return port.PlanResponse{}, fmt.Errorf(
				"failed to get latest plan version: %w",
				err,
			)
		}
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

func (s *GenerateLLMPlanService) Replan(
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

	// 1. Get DaySession
	daySession, err := s.daySessionRepo.GetDaySession(ctx, id)
	if err != nil {
		return port.PlanResponse{}, err
	}

	// 2. Get Trip
	trip, err := s.tripRepo.GetByID(
		ctx,
		daySession.TripID(),
	)
	if err != nil {
		return port.PlanResponse{}, err
	}

	// 3. Get current/latest plan version
	currentVersion, err := s.planVersionRepo.GetActivePlan(ctx, id)
	if err != nil {
		return port.PlanResponse{}, err
	}

	// 4. Get existing plan stops
	existingStops, err := s.planStopRepo.ListStop(
		ctx,
		currentVersion.ID(),
	)
	if err != nil {
		return port.PlanResponse{}, err
	}

	replanstops := make([]port.RePlanStop, 0, len(existingStops))

	for _, stop := range existingStops {
		replanstops = append(replanstops, port.RePlanStop{
			Position:         stop.Position(),
			Title:             stop.Title(),
			CategoryLabel:    stop.CategoryLabel(),
			ImageURL:         stop.URL(),
			PlannedArrival:   stop.PlannedArrival(),
			PlannedDeparture: stop.PlannedDeparture(),
			TravelMinutes:    stop.TravelMinutes(),
			StayMinutes:      stop.StayMinutes(),
		})
	}

	// 5. Send existing plan to AI for replanning
	replan, err := s.planGenerator.Replan(
		ctx,
		port.RePlanRequest{
			DaySessionID:      daySessionID,
			Destination:       trip.Destination(),
			Date:              daySession.Date(),
			StartTime:         daySession.STime(),
			StartLabel:        daySession.Label(),
			ExistingPlanStops: replanstops,
		},
	)
	if err != nil {
		return port.PlanResponse{}, err
	}

	// 6. Get latest version number
	latestVersion, err := s.planVersionRepo.GetLatestVersion(ctx, id)
	if err != nil {
		return port.PlanResponse{}, fmt.Errorf(
			"failed to get latest plan version: %w",
			err,
		)
	}

	nextVersion := latestVersion + 1

	// 7. Create new PlanVersion
	version, err := entity.NewPlanVersion(
		id,
		nextVersion,
		"Replanned version",
	)
	if err != nil {
		return port.PlanResponse{}, err
	}

	if err := s.planVersionRepo.Create(ctx, version); err != nil {
		return port.PlanResponse{}, fmt.Errorf(
			"failed to create plan version: %w",
			err,
		)
	}

	// 8. Save new AI-generated stops
	for _, aiStop := range replan.Stops {

		busyRisk := busy.Label(
			aiStop.CategoryLabel,
			aiStop.PlannedArrival,
		)

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
			busyRisk,
		)

		if err != nil {
			return port.PlanResponse{}, err
		}

		if err := s.planStopRepo.Create(ctx, planStop); err != nil {
			return port.PlanResponse{}, fmt.Errorf(
				"failed to create plan stop: %w",
				err,
			)
		}
	}

	return replan, nil
}
