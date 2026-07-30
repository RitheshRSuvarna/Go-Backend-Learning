package repository

import (
	"common"
	"context"
	"day_session/domain/entity"
	"fmt"
)

func (r *PostgresDaySessionRepository) GetDaySession(
	ctx context.Context, daysessionid common.DaySessionID,
) (*entity.DaySession, error) {

	pgUUID, err := uuidStringToPgUUID(daysessionid.String())
	if err != nil {
		return nil, err
	}

	row, err := r.getQueries(ctx).GetDaySession(ctx,pgUUID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get day session:%w", err)
	}
	daySession, err := rowToDomainDaySession(
		row.ID,
		row.TripID,
		row.Date,
		row.StartTime,
		row.StartLabel,
		row.ActivePlanVersionID,
		row.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return daySession, nil
}
