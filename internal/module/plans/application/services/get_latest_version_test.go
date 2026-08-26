package services

import (
	"context"
	"errors"
	"testing"

	"common"
	"plans/domain/entity"
)

// ------------------------------------------------------------
// Fake PlanVersion Repository
// ------------------------------------------------------------

type fakeLatestVersionRepository struct {
	latestVersion int
	getError      error
}

func (f *fakeLatestVersionRepository) GetLatestVersion(
	ctx context.Context,
	id common.DaySessionID,
) (int, error) {
	if f.getError != nil {
		return 0, f.getError
	}

	return f.latestVersion, nil
}

// Add these methods if your actual PlanVersionRepository
// interface requires them.

func (f *fakeLatestVersionRepository) Create(
	ctx context.Context,
	planVersion *entity.PlanVersion,
) error {
	return nil
}

func (f *fakeLatestVersionRepository) GetActivePlan(
	ctx context.Context,
	id common.DaySessionID,
) (*entity.PlanVersion, error) {
	return nil, nil
}

func (f *fakeLatestVersionRepository) GetByID(
	ctx context.Context,
	id common.PlanVersionID,
) (*entity.PlanVersion, error) {
	return nil, nil
}

func (f *fakeLatestVersionRepository) ListPlanVersion(
	ctx context.Context,
	id common.DaySessionID,
) ([]*entity.PlanVersion, error) {
	return nil, nil
}

// ------------------------------------------------------------
// Helper
// ------------------------------------------------------------

func createTestDaySessionIDForLatestVersion(
	t *testing.T,
) common.DaySessionID {

	t.Helper()

	id, err := common.NewDaySessionID("day-session-1")

	if err != nil {
		t.Fatalf(
			"failed to create day session ID: %v",
			err,
		)
	}

	return id
}

// ------------------------------------------------------------
// SUCCESS
// ------------------------------------------------------------

func TestGetLatestVersionService_GetLatestVersion_Success(
	t *testing.T,
) {

	repo := &fakeLatestVersionRepository{
		latestVersion: 3,
	}

	service := NewGetLatestVersion(repo)

	daySessionID := createTestDaySessionIDForLatestVersion(t)

	result, err := service.GetLatestVersion(
		context.Background(),
		daySessionID,
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if result != 3 {
		t.Errorf(
			"expected latest version 3, got %d",
			result,
		)
	}
}

// ------------------------------------------------------------
// REPOSITORY ERROR
// ------------------------------------------------------------

func TestGetLatestVersionService_GetLatestVersion_RepositoryError(
	t *testing.T,
) {

	expectedError := errors.New(
		"database error",
	)

	repo := &fakeLatestVersionRepository{
		getError: expectedError,
	}

	service := NewGetLatestVersion(repo)

	daySessionID := createTestDaySessionIDForLatestVersion(t)

	result, err := service.GetLatestVersion(
		context.Background(),
		daySessionID,
	)

	if err == nil {
		t.Fatal(
			"expected error, got nil",
		)
	}

	if result != 0 {
		t.Errorf(
			"expected version 0, got %d",
			result,
		)
	}

	// errors.Is works because your service uses %w.
	if !errors.Is(err, expectedError) {
		t.Errorf(
			"expected wrapped error %v, got %v",
			expectedError,
			err,
		)
	}
}
