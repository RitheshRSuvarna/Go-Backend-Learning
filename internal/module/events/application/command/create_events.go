package command

import (
	"encoding/json"
)

type CreateEventsCommand struct {
	DaySessionID string
	EventType string
	Payload json.RawMessage
}