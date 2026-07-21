package repository

import (
	"fmt"
	"context"
	"events/domain/entity"
	"common"

)

func (r *PostgresEventsRepository) GetEvents(ctx context.Context, id common.DaySessionID) ([]*entity.Events, error) {
	pgID, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, err
	}
	rows, err := r.getQueries(ctx).GetEvents(ctx, pgID)
	if err != nil {
		return nil, fmt.Errorf("Unable to get events:%w", err)
	}

	out := make([]*entity.Events, 0, len(rows))

	for _, row := range rows {
		event, err := rowToDomainEvents(
			row.ID,
			row.DaySessionID,
			row.Type,
			common.NewTime(row.Ts.Time),
			row.PayloadJson,
			row.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		out = append(out, event)
	}
	return out, nil
}