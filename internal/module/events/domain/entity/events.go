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

	if EventType == "" {
		return nil, common.NewValidationError("Status cannot be empty", nil)
	}

	switch EventType {
	case "Reached", "Delay", "Skip":
	default:
		return nil, common.NewValidationError("Invalid status", nil)
	}

	if len(payload) == 0 {
		return nil, common.NewValidationError("payload cannot be empty", nil)
	}

	if !json.Valid(payload) {
		return nil, common.NewValidationError("payload must be valid JSON", nil)
	}

	now := common.Now()

	return &Events{
		daySessionID: daysessionid,
		eventType:    EventType,
		ts:           now,
		payload:      payload,
		createdAt:    now,
	}, nil
}

func (e *Events) ID() common.EventsID               { return e.id }
func (e *Events) DaysessionID() common.DaySessionID { return e.daySessionID }
func (e *Events) EventType() string                 { return e.eventType }
func (e *Events) TS() common.Time                   { return e.ts }
func (e *Events) Payload() json.RawMessage          { return e.payload }
func (e *Events) CreatedAt() common.Time            { return e.createdAt }

func (e *Events) SetID(id common.EventsID) {
	e.id = id
}

func (e *Events) AssignPersistance(id common.EventsID, createdAt common.Time) {
	e.id = id
	e.createdAt = createdAt
}

func RestoreEvents(
	id common.EventsID,
	daysessionID common.DaySessionID,
	eventType string,
	ts common.Time,
	payload json.RawMessage,
	createdAt common.Time,
) *Events {
	return &Events{
		id:           id,
		daySessionID: daysessionID,
		eventType:    eventType,
		ts:           ts,
		payload:      payload,
		createdAt:    createdAt,
	}
}
