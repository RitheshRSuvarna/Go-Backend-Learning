package repository

import (
	"context"
	"common"
	"day_session/domain/entity"
	dsqueries "day_session/infrastructure/driven/postgres/queries/day_session"
	"fmt"
)

func (r *PostgresDaySessionRepository) GetByIDAndDate(ctx context.Context,
	tripID common.TripID, date entity.DaySession) (*entity.DaySession, error) {

	row, err := r.getQueries(ctx).GetDaySessionByIDAndDate(
		ctx, dsqueries.GetDaySessionByIDAndDateparams {
		TripID: tripID,
		Date: date,
	},
)
	if err != nil {
		return nil, fmt.Errorf("Failed to get day session:%w", err)
	}
		daySession, err := rowToDomainDaySession (
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
	return daysession, nil
}