package services

import (
	"context"
	"fmt"

	"day_session/domain/port"
)

type GenerateLLMPlanService struct {
	planGenerator port.PlanGenerator
}

func NewGenerateLLMPlanService(planGenerator port.PlanGenerator) *GenerateLLMPlanService {
	return &GenerateLLMPlanService{planGenerator: planGenerator}
}

func (s *GenerateLLMPlanService) GeneratePlan(ctx context.Context, daySessionID string) (port.PlanResponse, error) {
	if daySessionID == "" {
		return port.PlanResponse{}, fmt.Errorf("day session id is required")
	}

	return s.planGenerator.GeneratePlan(ctx, port.PlanRequest{
		DaySessionID: daySessionID,
	})
}
