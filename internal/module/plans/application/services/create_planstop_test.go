package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"common"
	"plans/application/command"
	"plans/domain/entity"
)

// Fake repository
type fakePlanStopRepository struct {
	createCalled bool
	createdStop  *entity.PlanStop
	createErr    error
}

func (f *fakePlanStopRepository) Create(
	ctx context.Context,
	planstop *entity.PlanStop,
) error {
	f.createCalled = true
	f.createdStop = planstop

	return f.createErr
}

// Add these if your actual PlanStopRepository interface
// contains additional methods.
func (f *fakePlanStopRepository) ListStop(
	ctx context.Context,
	id common.PlanVersionID,
) ([]*entity.PlanStop, error) {
	return nil, nil
}

// Helper to create a valid PlanVersionID
func createTestPlanVersionID(t *testing.T) common.PlanVersionID {
	t.Helper()

	id, err := common.NewPlanVersionID("plan-version-1")
	if err != nil {
		t.Fatalf("failed to create plan version ID: %v", err)
	}

	return id
}

// ------------------------------------------------------------
// SUCCESS TEST
// ------------------------------------------------------------

func TestCreatePlanStopService_CreateStop_Success(t *testing.T) {

	repo := &fakePlanStopRepository{}

	service := NewCreatePlanStopService(repo)

	planVersionID := createTestPlanVersionID(t)

	cmd := command.CreatePlanStopCommand{
		Position:         1,
		Title:            "Baga Beach",
		CategoryLabel:    "Beach",
		ImageURL:         "https://example.com/baga.jpg",
		PlannedArrival:   time.Date(2026, 8, 20, 10, 0, 0, 0, time.UTC),
		PlannedDeparture: time.Date(2026, 8, 20, 12, 0, 0, 0, time.UTC),
		TravelMinutes:    20,
		StayMinutes:      120,
		BusyRiskLabel:    "Low",
	}

	result, err := service.CreateStop(
		context.Background(),
		planVersionID,
		cmd,
	)

	// Service should not return an error.
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// DTO should contain the correct data.
	if result.Title != "Baga Beach" {
		t.Errorf(
			"expected title Baga Beach, got %s",
			result.Title,
		)
	}

	if result.Position != 1 {
		t.Errorf(
			"expected position 1, got %d",
			result.Position,
		)
	}

	// Repository should have been called.
	if !repo.createCalled {
		t.Fatal("expected repository Create to be called")
	}

	// Repository should receive the created PlanStop.
	if repo.createdStop == nil {
		t.Fatal("expected created PlanStop, got nil")
	}

	if repo.createdStop.Title() != "Baga Beach" {
		t.Errorf(
			"expected repository to receive Baga Beach, got %s",
			repo.createdStop.Title(),
		)
	}
}

// ------------------------------------------------------------
// INVALID POSITION
// ------------------------------------------------------------

func TestCreatePlanStopService_CreateStop_InvalidPosition(t *testing.T) {

	repo := &fakePlanStopRepository{}

	service := NewCreatePlanStopService(repo)

	planVersionID := createTestPlanVersionID(t)

	cmd := command.CreatePlanStopCommand{
		Position:         0, // invalid
		Title:            "Baga Beach",
		CategoryLabel:    "Beach",
		ImageURL:         "https://example.com/baga.jpg",
		PlannedArrival:   time.Date(2026, 8, 20, 10, 0, 0, 0, time.UTC),
		PlannedDeparture: time.Date(2026, 8, 20, 12, 0, 0, 0, time.UTC),
		TravelMinutes:    20,
		StayMinutes:      120,
		BusyRiskLabel:    "Low",
	}

	_, err := service.CreateStop(
		context.Background(),
		planVersionID,
		cmd,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Because entity.NewPlanStop failed,
	// repository.Create should NOT be called.
	if repo.createCalled {
		t.Fatal("repository Create should not be called")
	}
}

// ------------------------------------------------------------
// EMPTY TITLE
// ------------------------------------------------------------

func TestCreatePlanStopService_CreateStop_EmptyTitle(t *testing.T) {

	repo := &fakePlanStopRepository{}

	service := NewCreatePlanStopService(repo)

	planVersionID := createTestPlanVersionID(t)

	cmd := command.CreatePlanStopCommand{
		Position:         1,
		Title:            "", // invalid
		CategoryLabel:    "Beach",
		ImageURL:         "https://example.com/baga.jpg",
		PlannedArrival:   time.Date(2026, 8, 20, 10, 0, 0, 0, time.UTC),
		PlannedDeparture: time.Date(2026, 8, 20, 12, 0, 0, 0, time.UTC),
		TravelMinutes:    20,
		StayMinutes:      120,
		BusyRiskLabel:    "Low",
	}

	_, err := service.CreateStop(
		context.Background(),
		planVersionID,
		cmd,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if repo.createCalled {
		t.Fatal("repository Create should not be called")
	}
}

// ------------------------------------------------------------
// REPOSITORY ERROR
// ------------------------------------------------------------

func TestCreatePlanStopService_CreateStop_RepositoryError(t *testing.T) {

	expectedError := errors.New("database error")

	repo := &fakePlanStopRepository{
		createErr: expectedError,
	}

	service := NewCreatePlanStopService(repo)

	planVersionID := createTestPlanVersionID(t)

	cmd := command.CreatePlanStopCommand{
		Position:         1,
		Title:            "Baga Beach",
		CategoryLabel:    "Beach",
		ImageURL:         "https://example.com/baga.jpg",
		PlannedArrival:   time.Date(2026, 8, 20, 10, 0, 0, 0, time.UTC),
		PlannedDeparture: time.Date(2026, 8, 20, 12, 0, 0, 0, time.UTC),
		TravelMinutes:    20,
		StayMinutes:      120,
		BusyRiskLabel:    "Low",
	}

	_, err := service.CreateStop(
		context.Background(),
		planVersionID,
		cmd,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, expectedError) {
		t.Errorf(
			"expected %v, got %v",
			expectedError,
			err,
		)
	}

	// Repository should have been called.
	if !repo.createCalled {
		t.Fatal("expected repository Create to be called")
	}
}
