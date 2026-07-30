package repository

import (
	"common"
	"day_session/domain/entity"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func rowToDomainDaySession(
	id pgtype.UUID,
	tripID pgtype.UUID,
	date pgtype.Date,
	startTime, startLabel string,
	activePlanVersionID pgtype.UUID,
	createdAt pgtype.Timestamptz,
) (*entity.DaySession, error) {
	daysessionID, err := common.NewDaySessionID(pgTypeToString(id))
	if err != nil {
		return nil, fmt.Errorf("Invalid day session id: %v", err)
	}

	var activePlanID *common.PlanVersionID

	if activePlanVersionID.Valid {
    	pid, err := common.NewPlanVersionID(pgTypeToString(activePlanVersionID))
    	if err != nil {
        	return nil, fmt.Errorf("invalid active plan version id: %v", err)
    	}
    	activePlanID = &pid
	}
	
	tripiD, err := common.NewTripID(pgTypeToString(tripID))
	if err != nil {
		return nil, fmt.Errorf("invalid trip id: %v", err)
	}

	if !createdAt.Valid {
		return nil, fmt.Errorf("invalid created_At")
	}

	return entity.RestoreDaySession(
		daysessionID,
		tripiD,
		pgDateToString(date),
		startTime,
		startLabel,
		activePlanID,
		common.NewTime(createdAt.Time),
	), nil
}
