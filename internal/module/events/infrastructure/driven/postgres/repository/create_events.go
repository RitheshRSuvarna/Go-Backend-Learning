package repository

import (
	"common"
	"context"
	"events/domain/entity"
	eventqueries "events/infrastructure/driven/postgres/queries/events"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *PostgresEventsRepository) CreateEvents(ctx context.Context, event *entity.Events) error {
	daysessionid, err := uuidStringToPgUUID(event.DaysessionID().String())
	if err != nil {
		return common.NewValidationError("Invalid daysessionId", err)
	}

	row, err := r.getQueries(ctx).CreateEvents(ctx, eventqueries.CreateEventsParams{
		DaySessionID: daysessionid,
		Type:         event.EventType(),
		Ts: pgtype.Timestamptz{
			Time:  event.TS().Value(),
			Valid: true,
		},
		PayloadJson: event.Payload(),
	})
	if err != nil {
		return fmt.Errorf("Unable to create event:%w", err)
	}

	created, err := rowToDomainEvents(
		row.ID,
		row.DaySessionID,
		row.Type,
		common.NewTime(row.Ts.Time),
		row.PayloadJson,
		row.CreatedAt,
	)
	if err != nil {
		return err
	}
	event.SetID(created.ID())
	return nil

}
