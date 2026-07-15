package repository

import (
	"common"
	"context"
	"fmt"
	"plan/domain/entity"
	// pvqueries "plan/infrastructure/driven/postgres/queries/plans"
)

func (r *PostgresPlanVersionRepository) GetActivePlan(ctx context.Context, id common.DaySessionID) (*entity.PlanVersion, error) {
	pgid, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, err
	}

	row, err := r.getQueries(ctx).GetActivePlan(ctx, pgid)
	if err != nil {
		return nil, fmt.Errorf("Failed to get active plan:%w", err)
	}

	activeplan, err := rowToDomainPlanVersion(
		row.ID,
		row.DaySessionID,
		int(row.Version),
		row.Notes.String,
		row.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return activeplan, nil
}
