package repository

import (
	"common"
	"context"
	"fmt"
	"plan/domain/entity"
	pvqueries "plan/infrastructure/driven/postgres/queries/plans"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *PostgresPlanVersionRepository) Create(ctx context.Context, planversion *entity.PlanVersion) error {
	daysessionid, err := uuidStringToPgUUID(planversion.DaySessionID().String())
	if err != nil {
		return common.NewValidationError("Invalid Daysession id", err)
	}

	row, err := r.getQueries(ctx).CreatePlanVersion(ctx, pvqueries.CreatePlanVersionParams{
		DaySessionID: daysessionid,
		Version:      int32(planversion.Version()),
		Notes: pgtype.Text{
			String: planversion.Note(),
			Valid:  true,
		},
	})
	if err != nil {
		return fmt.Errorf("Failed to create plan version: %w", err)
	}

	created, err := rowToDomainPlanVersion(
		row.ID,
		row.DaySessionID,
		int(row.Version),
		row.Notes.String,
		row.CreatedAt,
	)
	if err != nil {
		return err
	}

	planversion.SetID(created.ID())
	return nil
}
