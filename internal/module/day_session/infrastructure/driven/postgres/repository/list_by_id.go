package repository

import (
	"common"
	"context"
	"day_session/domain/entity"
	"fmt"
)

func (r *PostgresDaySessionRepository) GetByID(
	ctx context.Context,
	id common.TripID,
) ([]*entity.DaySession, error) {
	pgID, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, err
	}

	rows, err := r.getQueries(ctx).GetByID(
		ctx,
		pgID,
	)
	if err != nil {
		fmt.Println("Repository List Error:", err)
		return nil, err
	}

	daySessions := make([]*entity.DaySession, 0, len(rows))

	for _, row := range rows {
		daySession, err := rowToDomainDaySession(
			row.ID,
			row.TripID,
			row.Date,
			row.StartTime,
			row.StartLabel,
			row.CreatedAt,
		)
		if err != nil {
			fmt.Println("rowToDomainDaySession Error:", err)
			return nil, err
		}

		daySessions = append(daySessions, daySession)
	}
	return daySessions, nil
}
