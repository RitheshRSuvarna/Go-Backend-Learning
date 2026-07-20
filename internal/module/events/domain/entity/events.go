package entity

import (
	"common"
	"encoding/json"
)

type Events struct {
	id           common.EventsID
	daySessionID common.DaySessionID
	eventType    string
	ts           common.Time
	payload      json.RawMessage
	createdAt    common.Time
}

func NewEvents(DaysessionID, EventType string, payload json.RawMessage) (*Events, error) {
	daysessionid, err := common.NewDaySessionID(DaysessionID)
	if err != nil {
		return nil, common.NewValidationError("Invalid daysession id", err)
	}

	if eventtype == "" {
		return nil, common.NewValidationError("Event type cannot be empty", err)
	}

	if len(payload) == 0 {
    return nil, common.NewValidationError("payload cannot be empty", nil)
	}

	if !json.Valid(payload) {
    return nil, common.NewValidationError("payload must be valid JSON", nil)
}





}
