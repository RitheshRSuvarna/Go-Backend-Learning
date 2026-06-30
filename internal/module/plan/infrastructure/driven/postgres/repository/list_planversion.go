package repository

import (
	"fmt"
	"context"
	"common"
	psqueries "plan/infrastructure/driven/postgres/queries/plans"
	"plan/domain/entity"
)

func(r *PostgresPlanVersionRepository) ListPlanversion(ctx context.Context, id common.DaySessionID) (*entity.PlanVersion, error) {
	pgid, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, err
	}

	row, err := r.getQueries(ctx).ListPlanversion(ctx, pgid)
	if err != nil {
		return nil, fmt.Errorf("Failed to list planversion:%w", err)
	}

	planversion, err := rowToDomainPlanVersion(
		row.ID,
		row.DaysessionID,
		row.Version,
		row.Note,
		row.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return planversion, nil
}