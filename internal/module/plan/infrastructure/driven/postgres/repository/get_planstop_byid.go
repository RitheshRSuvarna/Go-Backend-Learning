package repository

import (
	"common"
	"context"
	"fmt"
	"plan/domain/entity"
	// psqueries "plan/infrastructure/driven/postgres/queries/plans"
)

func (r *PostgresPlanStopRepository) GetPlanByID(ctx context.Context, id common.PlanStopID) (*entity.PlanStop, error) {
	pgid, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, err
	}

	row, err := r.getQueries(ctx).GetPlanStopByID(ctx, pgid)
	if err != nil {
		return nil, fmt.Errorf("Failed to get plan stop:%w", err)
	}

	planstop, err := rowToDomainPlanStop(
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
		return nil, err
	}
	return planstop, nil

}
