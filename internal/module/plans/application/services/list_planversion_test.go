package services

import (
	"context"
	"errors"
	"testing"

	"common"
	"plans/domain/entity"
)

// ----------------------------------------------------
// Fake PlanVersionRepository
// ----------------------------------------------------

type fakeListPlanVersionRepository struct {
	versions []*entity.PlanVersion
	err      error

	listCalled bool
}

func (f *fakeListPlanVersionRepository) Create(
	ctx context.Context,
	planversion *entity.PlanVersion,
) error {
	return nil
}

func (f *fakeListPlanVersionRepository) GetActivePlan(
	ctx context.Context,
	id common.DaySessionID,
) (*entity.PlanVersion, error) {
	return nil, nil
}

func (f *fakeListPlanVersionRepository) GetByID(
	ctx context.Context,
	id common.PlanVersionID,
) (*entity.PlanVersion, error) {
	return nil, nil
}

func (f *fakeListPlanVersionRepository) ListPlanVersion(
	ctx context.Context,
	id common.DaySessionID,
) ([]*entity.PlanVersion, error) {

	f.listCalled = true

	if f.err != nil {
		return nil, f.err
	}

	return f.versions, nil
}

func (f *fakeListPlanVersionRepository) GetLatestVersion(
	ctx context.Context,
	id common.DaySessionID,
) (int, error) {
	return 0, nil
}

// ----------------------------------------------------
// Helper
// ----------------------------------------------------

func createTestDaySessionIDForList(t *testing.T) common.DaySessionID {
	t.Helper()

	id, err := common.NewDaySessionID("day-session-1")
	if err != nil {
		t.Fatalf("failed to create day session ID: %v", err)
	}

	return id
}

// ----------------------------------------------------
// Test 1: Successful list
// ----------------------------------------------------

func TestListPlanVersion_Success(t *testing.T) {

	daySessionID := createTestDaySessionIDForList(t)

	version1, err := entity.NewPlanVersion(
		daySessionID,
		1,
		"Initial version",
	)

	if err != nil {
		t.Fatalf("failed to create version1: %v", err)
	}

	version2, err := entity.NewPlanVersion(
		daySessionID,
		2,
		"Replanned version",
	)

	if err != nil {
		t.Fatalf("failed to create version2: %v", err)
	}

	repo := &fakeListPlanVersionRepository{
		versions: []*entity.PlanVersion{
			version1,
			version2,
		},
	}

	service := NewListPlanVersionService(repo)

	result, err := service.ListVersion(
		context.Background(),
		daySessionID,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !repo.listCalled {
		t.Fatal("expected ListPlanVersion to be called")
	}

	if len(result) != 2 {
		t.Fatalf(
			"expected 2 plan versions, got %d",
			len(result),
		)
	}

	if result[0].Version != 1 {
		t.Errorf(
			"expected first version to be 1, got %d",
			result[0].Version,
		)
	}

	if result[0].Note != "Initial version" {
		t.Errorf(
			"expected first note 'Initial version', got %s",
			result[0].Note,
		)
	}

	if result[1].Version != 2 {
		t.Errorf(
			"expected second version to be 2, got %d",
			result[1].Version,
		)
	}

	if result[1].Note != "Replanned version" {
		t.Errorf(
			"expected second note 'Replanned version', got %s",
			result[1].Note,
		)
	}
}

// ----------------------------------------------------
// Test 2: Repository error
// ----------------------------------------------------

func TestListPlanVersion_RepositoryError(t *testing.T) {

	expectedError := errors.New("database error")

	repo := &fakeListPlanVersionRepository{
		err: expectedError,
	}

	service := NewListPlanVersionService(repo)

	daySessionID := createTestDaySessionIDForList(t)

	result, err := service.ListVersion(
		context.Background(),
		daySessionID,
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
		t.Errorf(
			"expected nil result, got %v",
			result,
		)
	}
}

// ----------------------------------------------------
// Test 3: Empty list
// ----------------------------------------------------

func TestListPlanVersion_EmptyList(t *testing.T) {

	repo := &fakeListPlanVersionRepository{
		versions: []*entity.PlanVersion{},
	}

	service := NewListPlanVersionService(repo)

	daySessionID := createTestDaySessionIDForList(t)

	result, err := service.ListVersion(
		context.Background(),
		daySessionID,
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if result == nil {
		t.Fatal("expected empty slice, got nil")
	}

	if len(result) != 0 {
		t.Fatalf(
			"expected 0 versions, got %d",
			len(result),
		)
	}
}