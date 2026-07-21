package repository

import (
	"common"
	"encoding/json"
	"events/domain/entity"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func rowToDomainEvents(
	id, daysessionid pgtype.UUID,
	eventtype string,
	ts common.Time,
	payload json.RawMessage,
	createdAt pgtype.Timestamptz,
) (*entity.Events, error) {
	eventid, err := common.NewEventsID(pgTypeToString(id))
	if err != nil {
		return nil, fmt.Errorf("Invalid eventid:%v", err)
	}

	daysessionID, err := common.NewDaySessionID(pgTypeToString(daysessionid))
	if err != nil {
		return nil, fmt.Errorf("Invalid day session id: %v", err)
	}

	if !createdAt.Valid {
		return nil, fmt.Errorf("invalid create_At")
	}

	return entity.RestoreEvents(
		eventid,
		daysessionID,
		eventtype,
		ts,
		payload,
		common.NewTime(createdAt.Time),
	), nil
}
