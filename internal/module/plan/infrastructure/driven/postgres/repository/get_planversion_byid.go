package repository

import (
	"common"
	"context"
	"fmt"
	"plan/domain/entity"
	// pvqueries "plan/infrastructure/driven/postgres/queries/plans"
)

func (r *PostgresPlanVersionRepository) GetVersionByID(ctx context.Context, id common.PlanVersionID) (*entity.PlanVersion, error) {
	pgid, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, err
	}

	row, err := r.getQueries(ctx).GetPlanVersionByID(ctx, pgid)
	if err != nil {
		return nil, fmt.Errorf("Failed to get plan version:%w", err)
	}

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

	return planversion, nil
}
