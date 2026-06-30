package repository

import (
	"common"
	"context"
	"plan/domain/entity"
	pvqueries "plan/infrastructure/driven/postgres/queries/plans"
	"fmt"
)

func (r *PostgresPlanVersionRepository) Create(ctx context.Context, planversion *entity.PlanVersion) error {
	daysessionid, err := uuidStringToPgUUID(planversion.DaySessionID().String())
	if err != nil {
		return common.NewValidationError("Invalid Daysession id", err)
	}

	row, err := r.getQueries(ctx).CreatePlanVersion(ctx, pvqueries.CreatePlanVersionParams{
		DaySessionID: daysessionid,
		Version: planversion.Version(),
		Note: planversion.Note(),
	})
	if err != nil {
		return fmt.Errorf("Failed to create plan version: %w", err)
	}

	created, err := rowToDomainPlanVersion(
		row.ID,
		row.DaySessionID,
		row.Version,
		row.Note,
		row.CreatedAt,
	)
	if err != nil {
		return err
	}

	planversion.SetID(created.ID())
	return nil	
}
