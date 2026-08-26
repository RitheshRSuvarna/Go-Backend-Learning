package services

import (
	"context"
	"errors"
	"testing"

	"common"

	"trip/application/dto"
	"trip/domain/entity"
)

// Fake repository for testing GetTripByIDService.
type getTripByIDFakeTripRepository struct {
	trip       *entity.Trip
	getByIDErr error
}

func (f *getTripByIDFakeTripRepository) Create(
	ctx context.Context,
	trip *entity.Trip,
) error {
	return nil
}

func (f *getTripByIDFakeTripRepository) List(
	ctx context.Context,
) ([]*entity.Trip, error) {
	return nil, nil
}

func (f *getTripByIDFakeTripRepository) GetByID(
	ctx context.Context,
	id common.TripID,
) (*entity.Trip, error) {
	if f.getByIDErr != nil {
		return nil, f.getByIDErr
	}

	return f.trip, nil
}

func TestGetTripByIDService_Success(t *testing.T) {
	trip, err := entity.NewTrip(
		"Goa",
		"2026-08-20",
		"2026-08-25",
		2,
	)

	if err != nil {
		t.Fatalf("failed to create test trip: %v", err)
	}

	repo := &getTripByIDFakeTripRepository{
		trip: trip,
	}

	service := NewGetTripByIDService(repo)

	tripID, err := common.NewTripID("trip-1")
	if err != nil {
		t.Fatalf("failed to create trip ID: %v", err)
	}

	result, err := service.GetTripByID(
		context.Background(),
		tripID,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
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

func TestGetTripByIDService_RepositoryError(t *testing.T) {
	expectedError := errors.New("trip not found")

	repo := &getTripByIDFakeTripRepository{
		getByIDErr: expectedError,
	}

	service := NewGetTripByIDService(repo)

	tripID, err := common.NewTripID("trip-1")
	if err != nil {
		t.Fatalf("failed to create trip ID: %v", err)
	}

	result, err := service.GetTripByID(
		context.Background(),
		tripID,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, expectedError) {
		t.Errorf(
			"expected error %v, got %v",
			expectedError,
			err,
		)
	}

	if result != (dto.TripDTO{}) {
		t.Fatal("expected empty TripDTO")
	}
}
