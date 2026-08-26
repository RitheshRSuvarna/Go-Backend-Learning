package entity

import (
	"testing"
)

// import "common"

func TestNewDaySession_Valid(t *testing.T) {
	daySession, err := NewDaySession(
		"trip-1",
		"2026-08-20",
		"09:00",
		"Hotel",
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if daySession == nil {
		t.Fatal("expected DaySession, got nil")
	}

	if daySession.TripID().String() != "trip-1" {
		t.Errorf(
			"expected trip ID trip-1, got %s",
			daySession.TripID().String(),
		)
	}

	if daySession.Date() != "2026-08-20" {
		t.Errorf(
			"expected date 2026-08-20, got %s",
			daySession.Date(),
		)
	}

	if daySession.STime() != "09:00" {
		t.Errorf(
			"expected start time 09:00, got %s",
			daySession.STime(),
		)
	}

	if daySession.Label() != "Hotel" {
		t.Errorf(
			"expected start label Hotel, got %s",
			daySession.Label(),
		)
	}

	if daySession.ActivePlanVersionID() != nil {
		t.Error("expected active plan version ID to be nil")
	}
}

func TestNewDaySession_InvalidTripID(t *testing.T) {
	daySession, err := NewDaySession(
		"",
		"2026-08-20",
		"09:00",
		"Hotel",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if daySession != nil {
		t.Fatal("expected nil DaySession, got DaySession")
	}
}

func TestNewDaySession_EmptyDate(t *testing.T) {
	daySession, err := NewDaySession(
		"trip-1",
		"",
		"09:00",
		"Hotel",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if daySession != nil {
		t.Fatal("expected nil DaySession, got DaySession")
	}
}

func TestNewDaySession_EmptyStartTime(t *testing.T) {
	daySession, err := NewDaySession(
		"trip-1",
		"2026-08-20",
		"",
		"Hotel",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if daySession != nil {
		t.Fatal("expected nil DaySession, got DaySession")
	}
}

func TestNewDaySession_EmptyStartLabel(t *testing.T) {
	daySession, err := NewDaySession(
		"trip-1",
		"2026-08-20",
		"09:00",
		"",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if daySession != nil {
		t.Fatal("expected nil DaySession, got DaySession")
	}
}

func TestNewDaySession_InvalidDate(t *testing.T) {
	daySession, err := NewDaySession(
		"trip-1",
		"20-08-2026",
		"09:00",
		"Hotel",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if daySession != nil {
		t.Fatal("expected nil DaySession, got DaySession")
	}
}
