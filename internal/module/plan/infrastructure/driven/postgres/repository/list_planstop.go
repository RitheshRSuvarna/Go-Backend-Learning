package repository

import (
	"fmt"
	"context"
	"common"
	psqueries "plan/infrastructure/driven/postgres/queries/plans"
	"plan/domain/entity"
)

func(r *PostgresPlanStopRepository) ListPlanStop(ctx context.Context, id common.PlanVersionID) (*entity.PlanStop, error) {
	pgid, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, err
	}

	row, err := r.getQueries(ctx).Listplanstop(ctx, pgid)
	if err != nil {
		return nil, fmt.Errorf("Failed to list plan stop:%w", err)
	}

	planstop, err := rowToDomainPlanStop(
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
		return nil, err
	}
	return planstop, nil
}