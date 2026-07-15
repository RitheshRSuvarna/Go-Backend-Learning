package repository

import (
	"common"
	"context"
	"fmt"
	"plan/domain/entity"
	psqueries "plan/infrastructure/driven/postgres/queries/plans"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *PostgresPlanStopRepository) Create(ctx context.Context, planstop *entity.PlanStop) error {
	planversionid, err := uuidStringToPgUUID(planstop.PlanVersionID().String())
	if err != nil {
		return common.NewValidationError("Invalid planversion id:%w", err)
	}

	row, err := r.getQueries(ctx).CreatePlanStop(ctx, psqueries.CreatePlanStopParams{
		PlanVersionID: planversionid,
		Position:      int32(planstop.Position()),
		Title:         planstop.Title(),
		CategoryLabel: planstop.CategoryLabel(),
		ImageUrl: pgtype.Text{
			String: planstop.URL(),
			Valid:  true,
		},
		PlannedArrival: timeToPgTimestamp(planstop.PlannedArrival()),
		// PlannedArrival: pgtype.Timestamptz{
		// Time: planstop.PlannedArrival(),
		// Valid: true,
		// },
		PlannedDeparture: timeToPgTimestamp(planstop.PlannedDeparture()),
		// PlannedDeparture: pgtype.Timestamptz{
		// Time: planstop.PlannedDeparture(),
		// Valid: true,
		// },
		TravelMinutes: int32(planstop.TravelMinutes()),
		StayMinutes:   int32(planstop.StayMinutes()),
		BusyRiskLabel: planstop.BusyRiskLabel(),
	})
	if err != nil {
		return fmt.Errorf("Failed to create plan stop:%w", err)
	}

	created, err := rowToDomainPlanStop(
		row.ID,
		row.PlanVersionID,
		int(row.Position),
		row.Title,
		row.CategoryLabel,
		row.ImageUrl.String,
		row.PlannedArrival.Time,
		row.PlannedDeparture.Time,
		int(row.TravelMinutes),
		int(row.StayMinutes),
		row.BusyRiskLabel,
		row.CreatedAt,
	)
	if err != nil {
		return err
	}
	planstop.SetID(created.ID())
	return nil
}
