package services

import (
	"common"
	"context"
	"fmt"
	"plans/application/command"
	"plans/application/dto"
	"plans/domain/entity"
	"plans/domain/repository"
)

type CreatePlanStopService struct {
	planrepo repository.PlanStopRepository
}

func NewCreatePlanStopService(planrepo repository.PlanStopRepository) *CreatePlanStopService {
	return &CreatePlanStopService{planrepo: planrepo}
}

func (s *CreatePlanStopService) CreateStop(ctx context.Context, id common.PlanVersionID, cmd command.CreatePlanStopCommand) (dto.PlanStopDTO, error) {

	// arrival, err := time.Parse(
	// 	time.RFC3339,
	// 	cmd.PlannedArrival,
	// )
	// if err != nil {
	// 	return dto.PlanStopDTO{}, err
	// }

	// departure, err := time.Parse(
	// 	time.RFC3339,
	// 	cmd.PlannedDeparture,
	// )
	// if err != nil {
	// 	return dto.PlanStopDTO{}, err
	// }
	fmt.Println("Entered CreateStopService")

	planstop, err := entity.NewPlanStop(
		id, cmd.Position, cmd.Title, cmd.CategoryLabel,
		cmd.ImageURL, cmd.PlannedArrival, cmd.PlannedDeparture,
		cmd.TravelMinutes, cmd.StayMinutes, cmd.BusyRiskLabel)
	if err != nil {
		return dto.PlanStopDTO{}, err
	}

	fmt.Println("Before repository Create")

	if err := s.planrepo.Create(ctx, planstop); err != nil {
		return dto.PlanStopDTO{}, err
	}

	fmt.Println("After repository Create")

	return dto.ToPlanStopDTO(planstop), nil
}
