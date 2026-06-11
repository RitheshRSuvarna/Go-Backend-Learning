package repository

import (
	"common"
	"fmt"
	"trip/domain/entity"

	"github.com/jackc/pgx/v5/pgtype"
)


func rowToDomainTrip(
	id pgtype.UUID,
	destination string,
	startDate, endDate pgtype.Date,
	travelersCount int32,
	createdAt pgtype.Timestamptz,
)(*entity.Trip, error) {
	tripID, err := common.NewTripID(pgTypeToString(id))
	if err != nil {
		return nil, fmt.Errorf("Invalid trip id: %v", err)
	}

	if !createdAt.Valid {
		return nil, fmt.Errorf("invalid create_At")
	}

	return entity.RestoreTrip(
		tripID,
		destination,
		pgDateToString(startDate),
		pgDateToString(endDate),
		int (travelersCount),
		common.NewTime(createdAt.Time),
	), nil
}
