package repository

import (
	"common"
	"context"
	"day_session/domain/entity"
	dsqueries "day_session/infrastructure/driven/postgres/queries/day_session"
	"fmt"
)

func (r *PostgresDaySessionRepository) Create(ctx context.Context, daysession *entity.DaySession) error {
	date, err := dateStringToPGDate(daysession.Date())
	if err != nil {
		return common.NewValidationError("Invalid date", err)
	}
	tripID, err := uuidStringToPgUUID(daysession.TripID().String())
	if err != nil {
		return common.NewValidationError("Invalid tripid", err)
	}

	row, err := r.getQueries(ctx).CreateDaySession(ctx, dsqueries.CreateDaySessionParams{
		TripID:     tripID,
		Date:       date,
		StartTime:  daysession.STime(),
		StartLabel: daysession.Label(),
	})
	if err != nil {
		return fmt.Errorf("filed to create session: %w", err)
	}

	created, err := rowToDomainDaySession(
		row.ID,
		row.TripID,
		row.Date,
		row.StartTime,
		row.StartLabel,
		row.CreatedAt,
	)
	if err != nil {
		return err
	}
	daysession.SetID(created.ID())
	return nil
}
