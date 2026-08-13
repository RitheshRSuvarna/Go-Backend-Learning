package services

import (
	"context"
	"fmt"
	"common"

	"day_session/domain/port"
	"day_session/domain/repository"
	triprepository "trip/domain/repository"
)

type GenerateLLMPlanService struct {
    planGenerator  port.PlanGenerator
    daySessionRepo repository.DaySessionRepository
    tripRepo       triprepository.TripRepository
}

func NewGenerateLLMPlanService(
    planGenerator port.PlanGenerator,
    daySessionRepo repository.DaySessionRepository,
    tripRepo triprepository.TripRepository,
) *GenerateLLMPlanService {
    return &GenerateLLMPlanService{
        planGenerator:  planGenerator,
        daySessionRepo: daySessionRepo,
        tripRepo:       tripRepo,
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

    return s.planGenerator.GeneratePlan(
        ctx,
        port.PlanRequest{
            DaySessionID: daySessionID,
            Destination:  trip.Destination(),
            Date:         daySession.Date(),
            StartTime:    daySession.STime(),
            StartLabel:   daySession.Label(),
        },
    )
}