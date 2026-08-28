package entity

import "testing"

func TestNewTrip_Valid(t *testing.T) {
	trip, err := NewTrip(
		"Goa",
		"2026-08-20",
		"2026-08-25",
		2,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if trip == nil {
		t.Fatal("expected trip, got nil")
	}

	if trip.Destination() != "Goa" {
		t.Errorf("expected destination Goa, got %s", trip.Destination())
	}

	if trip.StartDate() != "2026-08-20" {
		t.Errorf("expected start date 2026-08-20, got %s", trip.StartDate())
	}

	if trip.EndDate() != "2026-08-25" {
		t.Errorf("expected end date 2026-08-25, got %s", trip.EndDate())
	}

	if trip.TravelersCount() != 2 {
		t.Errorf("expected travelers count 2, got %d", trip.TravelersCount())
	}
}

func TestNewTrip_EmptyDestination(t *testing.T) {
	trip, err := NewTrip(
		"",
		"2026-08-20",
		"2026-08-25",
		2,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if trip != nil {
		t.Fatal("expected nil trip, got trip")
	}
}

func TestNewTrip_EmptyStartDate(t *testing.T) {
	trip, err := NewTrip(
		"Goa",
		"",
		"2026-08-25",
		2,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if trip != nil {
		t.Fatal("expected nil trip, got trip")
	}
}

func TestNewTrip_EmptyEndDate(t *testing.T) {
	trip, err := NewTrip(
		"Goa",
		"2026-08-20",
		"",
		2,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if trip != nil {
		t.Fatal("expected nil trip, got trip")
	}
}

func TestNewTrip_InvalidTravelerCount(t *testing.T) {
	trip, err := NewTrip(
		"Goa",
		"2026-08-20",
		"2026-08-25",
		0,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if trip != nil {
		t.Fatal("expected nil trip, got trip")
	}
}

func TestNewTrip_InvalidStartDate(t *testing.T) {
	trip, err := NewTrip(
		"Goa",
		"20-08-2026",
		"2026-08-25",
		2,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if trip != nil {
		t.Fatal("expected nil trip, got trip")
	}
}

func TestNewTrip_InvalidEndDate(t *testing.T) {
	trip, err := NewTrip(
		"Goa",
		"2026-08-20",
		"25-08-2026",
		2,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if trip != nil {
		t.Fatal("expected nil trip, got trip")
	}
}