package repository

import (
	"common"
	"context"
	"day_session/domain/entity"
	"fmt"
)

func (r *PostgresDaySessionRepository) GetByID(
	ctx context.Context,
	id common.DaySessionID,
) (*entity.DaySession, error) {
	pgID, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, err
	}

	row, err := r.getQueries(ctx).GetByID(
		ctx,
		pgID,
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
