package entity

import (
	"encoding/json"
	"testing"
)

func TestNewEvents_Valid(t *testing.T) {
	payload := json.RawMessage(`{"plan_stop_id":"stop-1"}`)

	event, err := NewEvents(
		"day-session-1",
		"Reached",
		payload,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if event == nil {
		t.Fatal("expected Events, got nil")
	}

	if event.DaysessionID().String() != "day-session-1" {
		t.Errorf(
			"expected day session ID day-session-1, got %s",
			event.DaysessionID().String(),
		)
	}

	if event.EventType() != "Reached" {
		t.Errorf(
			"expected event type Reached, got %s",
			event.EventType(),
		)
	}

	if string(event.Payload()) != string(payload) {
		t.Errorf(
			"expected payload %s, got %s",
			payload,
			event.Payload(),
		)
	}
}

func TestNewEvents_InvalidDaySessionID(t *testing.T) {
	payload := json.RawMessage(`{"plan_stop_id":"stop-1"}`)

	event, err := NewEvents(
		"",
		"Reached",
		payload,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if event != nil {
		t.Fatal("expected nil Events, got Events")
	}
}

func TestNewEvents_EmptyEventType(t *testing.T) {
	payload := json.RawMessage(`{"plan_stop_id":"stop-1"}`)

	event, err := NewEvents(
		"day-session-1",
		"",
		payload,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if event != nil {
		t.Fatal("expected nil Events, got Events")
	}
}

func TestNewEvents_InvalidEventType(t *testing.T) {
	payload := json.RawMessage(`{"plan_stop_id":"stop-1"}`)

	event, err := NewEvents(
		"day-session-1",
		"InvalidEvent",
		payload,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if event != nil {
		t.Fatal("expected nil Events, got Events")
	}
}

func TestNewEvents_EmptyPayload(t *testing.T) {
	event, err := NewEvents(
		"day-session-1",
		"Reached",
		json.RawMessage{},
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if event != nil {
		t.Fatal("expected nil Events, got Events")
	}
}

func TestNewEvents_InvalidJSONPayload(t *testing.T) {
	payload := json.RawMessage(`{"plan_stop_id":}`)

	event, err := NewEvents(
		"day-session-1",
		"Reached",
		payload,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if event != nil {
		t.Fatal("expected nil Events, got Events")
	}
}

func TestNewEvents_EventTypes(t *testing.T) {
	eventTypes := []string{
		"Reached",
		"Delay",
		"Skip",
	}

	for _, eventType := range eventTypes {
		t.Run(eventType, func(t *testing.T) {
			payload := json.RawMessage(`{"plan_stop_id":"stop-1"}`)

			event, err := NewEvents(
				"day-session-1",
				eventType,
				payload,
			)

			if err != nil {
				t.Fatalf(
					"expected no error for event type %s, got %v",
					eventType,
					err,
				)
			}

			if event == nil {
				t.Fatal("expected Events, got nil")
			}

			if event.EventType() != eventType {
				t.Errorf(
					"expected event type %s, got %s",
					eventType,
					event.EventType(),
				)
			}
		})
	}
}
