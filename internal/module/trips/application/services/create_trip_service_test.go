package services

import (
	"common"
	"context"
	"errors"
	"testing"

	"trip/application/command"
	"trip/application/dto"
	"trip/domain/entity"
)

// Fake repository used only for testing.
type fakeTripRepository struct {
	createCalled bool
	createError  error
	createdTrip  *entity.Trip
}

func (f *fakeTripRepository) Create(
	ctx context.Context,
	trip *entity.Trip,
) error {
	f.createCalled = true
	f.createdTrip = trip

	return f.createError
}

func (f *fakeTripRepository) List(
	ctx context.Context,
) ([]*entity.Trip, error) {
	return nil, nil
}

func (f *fakeTripRepository) GetByID(
	ctx context.Context,
	id common.TripID,
) (*entity.Trip, error) {
	return nil, nil
}

func TestCreateTripService_CreateTrip_Success(t *testing.T) {
	repo := &fakeTripRepository{}

	service := NewCreateTripService(repo)

	cmd := command.CreateTripCommand{
		Destination:    "Goa",
		StartDate:      "2026-08-20",
		EndDate:        "2026-08-25",
		TravelersCount: 2,
	}

	result, err := service.CreateTrip(
		context.Background(),
		cmd,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !repo.createCalled {
		t.Fatal("expected repository Create to be called")
	}

	if repo.createdTrip == nil {
		t.Fatal("expected trip to be passed to repository")
	}

	if result.Destination != "Goa" {
		t.Errorf(
			"expected destination Goa, got %s",
			result.Destination,
		)
	}

	if result.StartDate != "2026-08-20" {
		t.Errorf(
			"expected start date 2026-08-20, got %s",
			result.StartDate,
		)
	}

	if result.EndDate != "2026-08-25" {
		t.Errorf(
			"expected end date 2026-08-25, got %s",
			result.EndDate,
		)
	}

	if result.TravelersCount != 2 {
		t.Errorf(
			"expected travelers count 2, got %d",
			result.TravelersCount,
		)
	}
}

func TestCreateTripService_CreateTrip_InvalidCommand(t *testing.T) {
	repo := &fakeTripRepository{}

	service := NewCreateTripService(repo)

	cmd := command.CreateTripCommand{
		Destination:    "",
		StartDate:      "2026-08-20",
		EndDate:        "2026-08-25",
		TravelersCount: 2,
	}

	result, err := service.CreateTrip(
		context.Background(),
		cmd,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if result != (dto.TripDTO{}) {
		t.Fatal("expected empty TripDTO")
	}

	if repo.createCalled {
		t.Fatal("repository Create should not be called when validation fails")
	}
}

func TestCreateTripService_CreateTrip_RepositoryError(t *testing.T) {
	expectedError := errors.New("database error")

	repo := &fakeTripRepository{
		createError: expectedError,
	}

	service := NewCreateTripService(repo)

	cmd := command.CreateTripCommand{
		Destination:    "Goa",
		StartDate:      "2026-08-20",
		EndDate:        "2026-08-25",
		TravelersCount: 2,
	}

	result, err := service.CreateTrip(
		context.Background(),
		cmd,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, expectedError) {
		t.Errorf(
			"expected repository error, got %v",
			err,
		)
	}

	if result != (dto.TripDTO{}) {
		t.Fatal("expected empty TripDTO")
	}

	if !repo.createCalled {
		t.Fatal("expected repository Create to be called")
	}
}
