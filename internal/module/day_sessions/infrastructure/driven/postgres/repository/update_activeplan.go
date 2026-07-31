package repository

import (
	"context"
	"common"
	dsqueries "day_session/infrastructure/driven/postgres/queries/day_session"
)

func (r *PostgresDaySessionRepository) UpdateActivePlan(ctx context.Context, daysessionid common.DaySessionID, planversionid common.PlanVersionID) error{
	id, err := uuidStringToPgUUID(daysessionid.String())
	if err != nil {
		return common.NewValidationError("Invalid daysessionid", err)
	}
	
	activeplan, err := uuidStringToPgUUID(planversionid.String())
	if err != nil {
		return common.NewValidationError("Invalid planversionid", err)
	}

	return r.queries.UpdateActivePlan(ctx, dsqueries.UpdateActivePlanParams{
		ID:                  id,
		ActivePlanVersionID: activeplan,
	})

}