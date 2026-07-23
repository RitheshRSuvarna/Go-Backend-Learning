package repository

import (
	"common"
	"context"
	"fmt"

	// psqueries "plan/infrastructure/driven/postgres/queries/plans"
	"plans/domain/entity"
)

func (r *PostgresPlanVersionRepository) ListPlanVersion(ctx context.Context, id common.DaySessionID) ([]*entity.PlanVersion, error) {
	pgid, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, err
	}

	row, err := r.getQueries(ctx).ListPlanVersionsByDaySessionID(ctx, pgid)
	if err != nil {
		return nil, fmt.Errorf("Failed to list planversion:%w", err)
	}
	out := make([]*entity.PlanVersion, len(row))
	for _, row := range row {
		planversion, err := rowToDomainPlanVersion(
			row.ID,
			row.DaySessionID,
			int(row.Version),
			row.Notes.String,
			row.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		out = append(out, planversion)
	}
	return out, nil
}
