package services

import (
	// "time"
	"common"
	"context"
	"plan/application/command"
	"plan/application/dto"
	"plan/domain/repository"
	"plan/domain/entity"
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

	planstop, err := entity.NewPlanStop(
		id, cmd.Position, cmd.Title, cmd.CategoryLabel, 
		cmd.ImageURL, cmd.PlannedArrival, cmd.PlannedDeparture, 
		cmd.TravelMinutes, cmd.StayMinutes, cmd.BusyRiskLabel)
		if err != nil {
			return dto.PlanStopDTO{}, err
		}
		if err := s.planrepo.Create(ctx, planstop); err != nil {
			return dto.PlanStopDTO{}, err
		}
		return dto.ToPlanStopDTO(planstop),nil
}