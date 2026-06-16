package repository

import (
	"context"
	"common"
	"day_session/domain/entity"
	dsqueries "day_session/infrastructure/driven/postgres/queries/day_session"
	"fmt"
)

func (r *PostgresDaySessionRepository) GetByID(
	ctx context.Context,
	id common.ID,
) (*entity.DaySession, error) {

	row, err := r.getQueries(ctx).GetByID(
		ctx,
		id.String(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get day session: %w", err)
	}

	daySession, err := rowToDomainDaySession(
		row.ID,
		row.TripID,
		row.Date,
		row.StartTime,
		row.StartLabel,
		row.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return daySession, nil
}