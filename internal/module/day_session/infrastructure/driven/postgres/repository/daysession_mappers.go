package repository

import (
	"common"
	"fmt"
	"day_session/domain/entity"

	"github.com/jackc/pgx/v5/pgtype"
)

func rowToDomainDaySession (
	id pgtype.UUID,
	tripID common.TripID,
	date pgtype.Date,
	startTime,startLabel string,
	createdAt pgtype.Timestamptz,
)(*entity.DaySession, error) {
	daysessionID, err := common.NewDaySessionID(pgTypeToString(id))
	if err != nil {
		return nil, fmt.Errorf("Invalid day session id: %v", err)
	}

	if !createdAt.Valid {
		return nil, fmt.Errorf("invalid created_At")
	}

	return entity.RestoreDaySession(
		daysessionID,
		tripID,
		pgDateToString(date),
		startTime,
		startLabel,
		common.NewTime(createdAt.Time),
	), nil
}