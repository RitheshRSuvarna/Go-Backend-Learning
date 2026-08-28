package entity

import (
	"testing"

	"common"
)

func TestNewPlanVersion_Valid(t *testing.T) {
	daySessionID, err := common.NewDaySessionID("day-session-1")
	if err != nil {
		t.Fatalf("failed to create day session ID: %v", err)
	}

	planVersion, err := NewPlanVersion(
		daySessionID,
		1,
		"Initial plan",
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if planVersion == nil {
		t.Fatal("expected PlanVersion, got nil")
	}

	if planVersion.DaySessionID() != daySessionID {
		t.Errorf(
			"expected day session ID %v, got %v",
			daySessionID,
			planVersion.DaySessionID(),
		)
	}

	if planVersion.Version() != 1 {
		t.Errorf(
			"expected version 1, got %d",
			planVersion.Version(),
		)
	}

	if planVersion.Note() != "Initial plan" {
		t.Errorf(
			"expected note Initial plan, got %s",
			planVersion.Note(),
		)
	}
}

func TestNewPlanVersion_EmptyDaySessionID(t *testing.T) {
	var daySessionID common.DaySessionID

	planVersion, err := NewPlanVersion(
		daySessionID,
		1,
		"Initial plan",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if planVersion != nil {
		t.Fatal("expected nil PlanVersion, got PlanVersion")
	}
}

func TestNewPlanVersion_InvalidVersion(t *testing.T) {
	daySessionID, err := common.NewDaySessionID("day-session-1")
	if err != nil {
		t.Fatalf("failed to create day session ID: %v", err)
	}

	planVersion, err := NewPlanVersion(
		daySessionID,
		0,
		"Initial plan",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if planVersion != nil {
		t.Fatal("expected nil PlanVersion, got PlanVersion")
	}
}

func TestNewPlanVersion_EmptyNote(t *testing.T) {
	daySessionID, err := common.NewDaySessionID("day-session-1")
	if err != nil {
		t.Fatalf("failed to create day session ID: %v", err)
	}

	planVersion, err := NewPlanVersion(
		daySessionID,
		1,
		"",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if planVersion != nil {
		t.Fatal("expected nil PlanVersion, got PlanVersion")
	}
}
