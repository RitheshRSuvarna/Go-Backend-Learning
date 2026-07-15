package repository

import (
	"common"
	"context"
	"day_session/domain/entity"
	dsqueries "day_session/infrastructure/driven/postgres/queries/day_session"
	"fmt"
)

func (r *PostgresDaySessionRepository) GetByTripIDAndDate(
	ctx context.Context,
	tripID common.TripID,
	date string,
) (*entity.DaySession, error) {

	pgDate, err := dateStringToPGDate(date)
	if err != nil {
		return nil, err
	}
	pgUUID, err := uuidStringToPgUUID(tripID.String())
	if err != nil {
		return nil, err
	}

	row, err := r.getQueries(ctx).GetDaySessionByIDAndDate(
		ctx,
		dsqueries.GetDaySessionByIDAndDateParams{
			TripID: pgUUID,
			Date:   pgDate,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to get day session:%w", err)
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
