package repository

import (
	"context"
	"common"
	"plan/domain/entity"
	psqueries "plan/infrastructure/driven/postgres/queries/plans"
	"fmt"
)

func (r *PostgresPlanStopRepository) Create(ctx context.Context, planstop *entity.PlanStop) error {
	planversionid, err := uuidStringToPgUUID(planstop.PlanVersionID().String())
	if err != nil {
		return common.NewValidationError("Invalid planversion id:%w", err)
	}

	row, err := r.getQueries(ctx).CreatePlanStopParams(ctx, psqueries.CreatePlanStopParams{
		PlanVersionID: planversionid,
		Position: planstop.Position(),
		Title: planstop.Title(),
		CategoryLabel: planstop.CategoryLabel(),
		ImageURL: planstop.URL(),
		PlannedArrival: planstop.PlannedArrival(),
		PlannedDeparture: planstop.PlannedDeparture(),
		TravelMinutes: planstop.TravelMinutes(),
		StayMinutes: planstop.StayMinutes(),
		BusyRiskLabel: planstop.BusyRiskLabel(),
	})
	if err != nil {
		return fmt.Errorf("Failed to create plan stop:%w", err)
	}

	created, err := rowToDomainPlanStop(
		row.ID,
		row.PlanVersionID,
		row.Position,
		row.Title,
		row.CategoryLabel,
		row.ImageURL,
		row.PlannedArrival,
		row.PlannedDeparture,
		row.TravelMinutes,
		row.StayMinutes,
		row.BuysRiskLabel,
		row.CreatedAt,
	)
	if err != nil {
		return err
	}
	planstop.SetID(created.ID())
	return nil
}