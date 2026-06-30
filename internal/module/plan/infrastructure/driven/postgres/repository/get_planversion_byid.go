package repository

import (
	"fmt"
	"common"
	"context"
	"plan/domain/entity"
	pvqueries "plan/infrastructure/driven/postgres/queries/plans"
)

func (r *PostgresPlanVersionRepository) GetByID(ctx context.Context, id common.DaySessionID) (*entity.PlanVersion, error) {
	pgid, err :=uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, err
	}

	row, err := r.getQueries(ctx).GetByID(ctx, pgid)
	if err != nil {
		return nil, fmt.Errorf("Failed to get plan version:%w", err)
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