package services

import (
	"context"
	"errors"
	"testing"

	"common"

	"trip/application/dto"
	"trip/domain/entity"
)

type listTripFakeRepository struct {
	trips   []*entity.Trip
	listErr error
}

func (f *listTripFakeRepository) Create(
	ctx context.Context,
	trip *entity.Trip,
) error {
	return nil
}

func (f *listTripFakeRepository) List(
	ctx context.Context,
) ([]*entity.Trip, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}

	return f.trips, nil
}

func (f *listTripFakeRepository) GetByID(
	ctx context.Context,
	id common.TripID,
) (*entity.Trip, error) {
	return nil, nil
}

func TestListTripService_Success(t *testing.T) {
	trip1, err := entity.NewTrip(
		"Goa",
		"2026-08-20",
		"2026-08-25",
		2,
	)
	if err != nil {
		t.Fatalf("failed to create trip1: %v", err)
	}

	trip2, err := entity.NewTrip(
		"Mysore",
		"2026-09-01",
		"2026-09-05",
		3,
	)
	if err != nil {
		t.Fatalf("failed to create trip2: %v", err)
	}

	repo := &listTripFakeRepository{
		trips: []*entity.Trip{
			trip1,
			trip2,
		},
	}

	service := NewTripListService(repo)

	result, err := service.ListTrips(context.Background())

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected TripDTO slice, got nil")
	}

	if len(result) != 2 {
		t.Fatalf(
			"expected 2 trips, got %d",
			len(result),
		)
	}

	if result[0].Destination != "Goa" {
		t.Errorf(
			"expected first destination Goa, got %s",
			result[0].Destination,
		)
	}

	if result[0].StartDate != "2026-08-20" {
		t.Errorf(
			"expected first start date 2026-08-20, got %s",
			result[0].StartDate,
		)
	}

	if result[1].Destination != "Mysore" {
		t.Errorf(
			"expected second destination Mysore, got %s",
			result[1].Destination,
		)
	}

	if result[1].TravelersCount != 3 {
		t.Errorf(
			"expected second travelers count 3, got %d",
			result[1].TravelersCount,
		)
	}
}

func TestListTripService_RepositoryError(t *testing.T) {
	expectedError := errors.New("database error")

	repo := &listTripFakeRepository{
		listErr: expectedError,
	}

	service := NewTripListService(repo)

	result, err := service.ListTrips(
		context.Background(),
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

	if result != nil {
		t.Fatal("expected nil result, got result")
	}
}

func TestListTripService_EmptyList(t *testing.T) {
	repo := &listTripFakeRepository{
		trips: []*entity.Trip{},
	}

	service := NewTripListService(repo)

	result, err := service.ListTrips(
		context.Background(),
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected empty slice, got nil")
	}

	if len(result) != 0 {
		t.Fatalf(
			"expected 0 trips, got %d",
			len(result),
		)
	}

	// Verify the service returns an empty slice rather than nil.
	expected := []dto.TripDTO{}

	if len(result) != len(expected) {
		t.Fatalf("expected empty DTO slice")
	}
}
