package repository

import (
	"common"
	"context"
	"fmt"

	// psqueries "plan/infrastructure/driven/postgres/queries/plans"
	"plans/domain/entity"
)

func (r *PostgresPlanStopRepository) ListStop(ctx context.Context, id common.PlanVersionID) ([]*entity.PlanStop, error) {
	pgid, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, err
	}

	row, err := r.getQueries(ctx).ListPlanStopsByPlanVersionID(ctx, pgid)
	if err != nil {
		return nil, fmt.Errorf("Failed to list plan stop:%w", err)
	}

	out := make([]*entity.PlanStop, len(row))
	for _, row := range row {
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
		out = append(out, planstop)
	}
	return out, nil
}
