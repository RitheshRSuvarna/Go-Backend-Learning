package services

import (
	"context"
	"errors"
	"testing"

	"common"
	"plans/application/command"
	"plans/domain/entity"
)

// ------------------------------------------------------------
// Fake PlanVersion Repository
// ------------------------------------------------------------

type fakePlanVersionRepository struct {
	createCalled bool
	createdPlan  *entity.PlanVersion
	createErr    error
}

func (f *fakePlanVersionRepository) Create(
	ctx context.Context,
	planVersion *entity.PlanVersion,
) error {
	f.createCalled = true
	f.createdPlan = planVersion

	return f.createErr
}

// Add these if your actual PlanVersionRepository interface
// contains these methods.

func (f *fakePlanVersionRepository) GetActivePlan(
	ctx context.Context,
	id common.DaySessionID,
) (*entity.PlanVersion, error) {
	return nil, nil
}

func (f *fakePlanVersionRepository) GetByID(
	ctx context.Context,
	id common.PlanVersionID,
) (*entity.PlanVersion, error) {
	return nil, nil
}

func (f *fakePlanVersionRepository) ListPlanVersion(
	ctx context.Context,
	id common.DaySessionID,
) ([]*entity.PlanVersion, error) {
	return nil, nil
}

func (f *fakePlanVersionRepository) GetLatestVersion(
	ctx context.Context,
	id common.DaySessionID,
) (int, error) {
	return 0, nil
}

// ------------------------------------------------------------
// Helper: create DaySessionID
// ------------------------------------------------------------

func createTestDaySessionID(t *testing.T) common.DaySessionID {
	t.Helper()

	id, err := common.NewDaySessionID("day-session-1")
	if err != nil {
		t.Fatalf("failed to create day session ID: %v", err)
	}

	return id
}

// ------------------------------------------------------------
// SUCCESS
// ------------------------------------------------------------

func TestCreatePlanVersionService_CreatePlanVersion_Success(
	t *testing.T,
) {

	repo := &fakePlanVersionRepository{}

	service := NewCreatePlanVersionService(repo)

	daySessionID := createTestDaySessionID(t)

	cmd := command.CreatePlanVersionCommand{
		Version: 1,
		Note:    "Initial version",
	}

	result, err := service.CreatePlanVersion(
		context.Background(),
		daySessionID,
		cmd,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check returned DTO
	if result.Version != 1 {
		t.Errorf(
			"expected version 1, got %d",
			result.Version,
		)
	}

	if result.Note != "Initial version" {
		t.Errorf(
			"expected note 'Initial version', got %s",
			result.Note,
		)
	}

	// Repository should have been called.
	if !repo.createCalled {
		t.Fatal("expected repository Create to be called")
	}

	// Repository should receive the PlanVersion.
	if repo.createdPlan == nil {
		t.Fatal("expected created PlanVersion, got nil")
	}

	if repo.createdPlan.Version() != 1 {
		t.Errorf(
			"expected repository to receive version 1, got %d",
			repo.createdPlan.Version(),
		)
	}

	if repo.createdPlan.Note() != "Initial version" {
		t.Errorf(
			"expected repository to receive note 'Initial version', got %s",
			repo.createdPlan.Note(),
		)
	}
}

// ------------------------------------------------------------
// INVALID VERSION
// ------------------------------------------------------------

func TestCreatePlanVersionService_CreatePlanVersion_InvalidVersion(
	t *testing.T,
) {

	repo := &fakePlanVersionRepository{}

	service := NewCreatePlanVersionService(repo)

	daySessionID := createTestDaySessionID(t)

	cmd := command.CreatePlanVersionCommand{
		Version: 0, // invalid
		Note:    "Initial version",
	}

	_, err := service.CreatePlanVersion(
		context.Background(),
		daySessionID,
		cmd,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Entity creation failed,
	// so repository should NOT be called.
	if repo.createCalled {
		t.Fatal("repository Create should not be called")
	}
}

// ------------------------------------------------------------
// EMPTY NOTE
// ------------------------------------------------------------

func TestCreatePlanVersionService_CreatePlanVersion_EmptyNote(
	t *testing.T,
) {

	repo := &fakePlanVersionRepository{}

	service := NewCreatePlanVersionService(repo)

	daySessionID := createTestDaySessionID(t)

	cmd := command.CreatePlanVersionCommand{
		Version: 1,
		Note:    "", // invalid
	}

	_, err := service.CreatePlanVersion(
		context.Background(),
		daySessionID,
		cmd,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Entity creation failed,
	// so repository should NOT be called.
	if repo.createCalled {
		t.Fatal("repository Create should not be called")
	}
}

// ------------------------------------------------------------
// REPOSITORY ERROR
// ------------------------------------------------------------

func TestCreatePlanVersionService_CreatePlanVersion_RepositoryError(
	t *testing.T,
) {

	expectedError := errors.New("database error")

	repo := &fakePlanVersionRepository{
		createErr: expectedError,
	}

	service := NewCreatePlanVersionService(repo)

	daySessionID := createTestDaySessionID(t)

	cmd := command.CreatePlanVersionCommand{
		Version: 1,
		Note:    "Initial version",
	}

	_, err := service.CreatePlanVersion(
		context.Background(),
		daySessionID,
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

	// Repository was called, but it returned an error.
	if !repo.createCalled {
		t.Fatal("expected repository Create to be called")
	}
}