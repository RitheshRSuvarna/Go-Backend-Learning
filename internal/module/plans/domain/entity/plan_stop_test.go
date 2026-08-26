package entity

import (
	"testing"
	"time"

	"common"
)

func TestNewPlanStop_Valid(t *testing.T) {
	planVersionID, err := common.NewPlanVersionID("plan-version-1")
	if err != nil {
		t.Fatalf("failed to create plan version ID: %v", err)
	}

	arrival := time.Date(
		2026, 8, 20,
		10, 0, 0, 0,
		time.UTC,
	)

	departure := time.Date(
		2026, 8, 20,
		11, 0, 0, 0,
		time.UTC,
	)

	stop, err := NewPlanStop(
		planVersionID,
		1,
		"Baga Beach",
		"Beach",
		"https://example.com/baga.jpg",
		arrival,
		departure,
		30,
		60,
		"Medium",
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if stop == nil {
		t.Fatal("expected PlanStop, got nil")
	}

	if stop.PlanVersionID() != planVersionID {
		t.Errorf(
			"expected plan version ID %v, got %v",
			planVersionID,
			stop.PlanVersionID(),
		)
	}

	if stop.Position() != 1 {
		t.Errorf(
			"expected position 1, got %d",
			stop.Position(),
		)
	}

	if stop.Title() != "Baga Beach" {
		t.Errorf(
			"expected title Baga Beach, got %s",
			stop.Title(),
		)
	}

	if stop.CategoryLabel() != "Beach" {
		t.Errorf(
			"expected category Beach, got %s",
			stop.CategoryLabel(),
		)
	}

	if stop.URL() != "https://example.com/baga.jpg" {
		t.Errorf(
			"expected image URL https://example.com/baga.jpg, got %s",
			stop.URL(),
		)
	}

	if !stop.PlannedArrival().Equal(arrival) {
		t.Errorf(
			"expected arrival %v, got %v",
			arrival,
			stop.PlannedArrival(),
		)
	}

	if !stop.PlannedDeparture().Equal(departure) {
		t.Errorf(
			"expected departure %v, got %v",
			departure,
			stop.PlannedDeparture(),
		)
	}

	if stop.TravelMinutes() != 30 {
		t.Errorf(
			"expected travel minutes 30, got %d",
			stop.TravelMinutes(),
		)
	}

	if stop.StayMinutes() != 60 {
		t.Errorf(
			"expected stay minutes 60, got %d",
			stop.StayMinutes(),
		)
	}

	if stop.BusyRiskLabel() != "Medium" {
		t.Errorf(
			"expected busy risk label Medium, got %s",
			stop.BusyRiskLabel(),
		)
	}
}

func TestNewPlanStop_EmptyPlanVersionID(t *testing.T) {
	planVersionID, err := common.NewPlanVersionID("plan-version-1")
	if err != nil {
		t.Fatalf("failed to create plan version ID: %v", err)
	}

	arrival := time.Date(
		2026, 8, 20,
		10, 0, 0, 0,
		time.UTC,
	)

	departure := time.Date(
		2026, 8, 20,
		11, 0, 0, 0,
		time.UTC,
	)

	stop, err := NewPlanStop(
		planVersionID,
		1,
		"Baga Beach",
		"Beach",
		"https://example.com/baga.jpg",
		arrival,
		departure,
		30,
		60,
		"Medium",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if stop != nil {
		t.Fatal("expected nil PlanStop, got PlanStop")
	}
}

func TestNewPlanStop_InvalidPosition(t *testing.T) {
	planVersionID, err := common.NewPlanVersionID("plan-version-1")
	if err != nil {
		t.Fatalf("failed to create plan version ID: %v", err)
	}

	arrival := time.Date(
		2026, 8, 20,
		10, 0, 0, 0,
		time.UTC,
	)

	departure := time.Date(
		2026, 8, 20,
		11, 0, 0, 0,
		time.UTC,
	)

	stop, err := NewPlanStop(
		planVersionID,
		0,
		"Baga Beach",
		"Beach",
		"https://example.com/baga.jpg",
		arrival,
		departure,
		30,
		60,
		"Medium",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if stop != nil {
		t.Fatal("expected nil PlanStop, got PlanStop")
	}
}

func TestNewPlanStop_EmptyTitle(t *testing.T) {
	planVersionID, err := common.NewPlanVersionID("plan-version-1")
	if err != nil {
		t.Fatalf("failed to create plan version ID: %v", err)
	}

	arrival := time.Date(
		2026, 8, 20,
		10, 0, 0, 0,
		time.UTC,
	)

	departure := time.Date(
		2026, 8, 20,
		11, 0, 0, 0,
		time.UTC,
	)

	stop, err := NewPlanStop(
		planVersionID,
		1,
		"",
		"Beach",
		"https://example.com/baga.jpg",
		arrival,
		departure,
		30,
		60,
		"Medium",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if stop != nil {
		t.Fatal("expected nil PlanStop, got PlanStop")
	}
}

func TestNewPlanStop_EmptyCategoryLabel(t *testing.T) {
	planVersionID, err := common.NewPlanVersionID("plan-version-1")
	if err != nil {
		t.Fatalf("failed to create plan version ID: %v", err)
	}

	arrival := time.Date(
		2026, 8, 20,
		10, 0, 0, 0,
		time.UTC,
	)

	departure := time.Date(
		2026, 8, 20,
		11, 0, 0, 0,
		time.UTC,
	)

	stop, err := NewPlanStop(
		planVersionID,
		1,
		"Baga Beach",
		"",
		"https://example.com/baga.jpg",
		arrival,
		departure,
		30,
		60,
		"Medium",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if stop != nil {
		t.Fatal("expected nil PlanStop, got PlanStop")
	}
}

func TestNewPlanStop_EmptyImageURL(t *testing.T) {
	planVersionID, err := common.NewPlanVersionID("plan-version-1")
	if err != nil {
		t.Fatalf("failed to create plan version ID: %v", err)
	}

	arrival := time.Date(
		2026, 8, 20,
		10, 0, 0, 0,
		time.UTC,
	)

	departure := time.Date(
		2026, 8, 20,
		11, 0, 0, 0,
		time.UTC,
	)

	stop, err := NewPlanStop(
		planVersionID,
		1,
		"Baga Beach",
		"Beach",
		"",
		arrival,
		departure,
		30,
		60,
		"Medium",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if stop != nil {
		t.Fatal("expected nil PlanStop, got PlanStop")
	}
}

func TestNewPlanStop_InvalidTravelMinutes(t *testing.T) {
	planVersionID, err := common.NewPlanVersionID("plan-version-1")
	if err != nil {
		t.Fatalf("failed to create plan version ID: %v", err)
	}

	arrival := time.Date(
		2026, 8, 20,
		10, 0, 0, 0,
		time.UTC,
	)

	departure := time.Date(
		2026, 8, 20,
		11, 0, 0, 0,
		time.UTC,
	)

	stop, err := NewPlanStop(
		planVersionID,
		1,
		"Baga Beach",
		"Beach",
		"https://example.com/baga.jpg",
		arrival,
		departure,
		0,
		60,
		"Medium",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if stop != nil {
		t.Fatal("expected nil PlanStop, got PlanStop")
	}
}

func TestNewPlanStop_InvalidStayMinutes(t *testing.T) {
	planVersionID, err := common.NewPlanVersionID("plan-version-1")
	if err != nil {
		t.Fatalf("failed to create plan version ID: %v", err)
	}

	arrival := time.Date(
		2026, 8, 20,
		10, 0, 0, 0,
		time.UTC,
	)

	departure := time.Date(
		2026, 8, 20,
		11, 0, 0, 0,
		time.UTC,
	)

	stop, err := NewPlanStop(
		planVersionID,
		1,
		"Baga Beach",
		"Beach",
		"https://example.com/baga.jpg",
		arrival,
		departure,
		30,
		0,
		"Medium",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if stop != nil {
		t.Fatal("expected nil PlanStop, got PlanStop")
	}
}

func TestNewPlanStop_EmptyBusyRiskLabel(t *testing.T) {
	planVersionID, err := common.NewPlanVersionID("plan-version-1")
	if err != nil {
		t.Fatalf("failed to create plan version ID: %v", err)
	}

	arrival := time.Date(
		2026, 8, 20,
		10, 0, 0, 0,
		time.UTC,
	)

	departure := time.Date(
		2026, 8, 20,
		11, 0, 0, 0,
		time.UTC,
	)

	stop, err := NewPlanStop(
		planVersionID,
		1,
		"Baga Beach",
		"Beach",
		"https://example.com/baga.jpg",
		arrival,
		departure,
		30,
		60,
		"",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if stop != nil {
		t.Fatal("expected nil PlanStop, got PlanStop")
	}
}
