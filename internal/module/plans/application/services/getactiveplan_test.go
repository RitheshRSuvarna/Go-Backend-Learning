package services

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"common"
// 	daySessionEntity "day_session/domain/entity"
// 	"plans/domain/entity"
// )

// // ============================================================
// // Fake PlanVersionRepository
// // ============================================================

// type fakePlanVersionRepository struct {
// 	planVersion *entity.PlanVersion

// 	getByIDError   error
// 	getActiveError error

// 	getByIDCalled   bool
// 	getActiveCalled bool
// }

// func (f *fakePlanVersionRepository) Create(
// 	ctx context.Context,
// 	planVersion *entity.PlanVersion,
// ) error {
// 	return nil
// }

// func (f *fakePlanVersionRepository) GetActivePlan(
// 	ctx context.Context,
// 	id common.DaySessionID,
// ) (*entity.PlanVersion, error) {

// 	f.getActiveCalled = true

// 	if f.getActiveError != nil {
// 		return nil, f.getActiveError
// 	}

// 	return f.planVersion, nil
// }

// func (f *fakePlanVersionRepository) GetByID(
// 	ctx context.Context,
// 	id common.PlanVersionID,
// ) (*entity.PlanVersion, error) {

// 	f.getByIDCalled = true

// 	if f.getByIDError != nil {
// 		return nil, f.getByIDError
// 	}

// 	return f.planVersion, nil
// }

// func (f *fakePlanVersionRepository) ListPlanVersion(
// 	ctx context.Context,
// 	id common.DaySessionID,
// ) ([]*entity.PlanVersion, error) {
// 	return []*entity.PlanVersion{f.planVersion}, nil
// }

// func (f *fakePlanVersionRepository) GetLatestVersion(
// 	ctx context.Context,
// 	id common.DaySessionID,
// ) (int, error) {
// 	return 1, nil
// }

// // ============================================================
// // Fake DaySessionRepository
// // ============================================================

// type fakeDaySessionRepository struct {
// 	daySession *daySessionEntity.DaySession
// 	getError   error
// }

// func (f *fakeDaySessionRepository) CreateDaySession(
// 	ctx context.Context,
// 	daySession *daySessionEntity.DaySession,
// ) error {
// 	return nil
// }

// func (f *fakeDaySessionRepository) GetDaySession(
// 	ctx context.Context,
// 	id common.DaySessionID,
// ) (*daySessionEntity.DaySession, error) {

// 	if f.getError != nil {
// 		return nil, f.getError
// 	}

// 	return f.daySession, nil
// }

// func (f *fakeDaySessionRepository) ListDaySession(
// 	ctx context.Context,
// 	id common.TripID,
// ) ([]*daySessionEntity.DaySession, error) {
// 	return []*daySessionEntity.DaySession{f.daySession}, nil
// }

// func (f *fakeDaySessionRepository) UpdateActivePlan(
// 	ctx context.Context,
// 	daySessionID common.DaySessionID,
// 	planVersionID common.PlanVersionID,
// ) error {
// 	return nil
// }

// // ============================================================
// // Helpers
// // ============================================================

// func createTestDaySessionID(t *testing.T) common.DaySessionID {
// 	t.Helper()

// 	id, err := common.NewDaySessionID("day-session-1")
// 	if err != nil {
// 		t.Fatalf("failed to create day session ID: %v", err)
// 	}

// 	return id
// }

// func createTestPlanVersionID(t *testing.T) common.PlanVersionID {
// 	t.Helper()

// 	id, err := common.NewPlanVersionID("plan-version-1")
// 	if err != nil {
// 		t.Fatalf("failed to create plan version ID: %v", err)
// 	}

// 	return id
// }

// // ============================================================
// // TEST 1
// // ActivePlanVersionID exists
// // ============================================================

// func TestGetActivePlan_WithActivePlanVersionID(t *testing.T) {

// 	daySessionID := createTestDaySessionID(t)
// 	planVersionID := createTestPlanVersionID(t)

// 	daySession, err := daySessionEntity.NewDaySession(
// 		"trip-1",
// 		"2026-08-20",
// 		"09:00",
// 		"Hotel",
// 	)

// 	if err != nil {
// 		t.Fatalf("failed to create day session: %v", err)
// 	}

// 	// Give the day session an active plan version.
// 	daySession.SetActivePlanVersionID(planVersionID)

// 	planVersion, err := entity.NewPlanVersion(
// 		daySessionID,
// 		1,
// 		"Initial version",
// 	)

// 	if err != nil {
// 		t.Fatalf("failed to create plan version: %v", err)
// 	}

// 	planRepo := &fakePlanVersionRepository{
// 		planVersion: planVersion,
// 	}

// 	dayRepo := &fakeDaySessionRepository{
// 		daySession: daySession,
// 	}

// 	service := NewGetByIDPlanVersionService(
// 		planRepo,
// 		dayRepo,
// 	)

// 	result, err := service.GetActivePlan(
// 		context.Background(),
// 		daySessionID,
// 	)

// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	// Because ActivePlanVersionID exists,
// 	// GetByID should have been called.
// 	if !planRepo.getByIDCalled {
// 		t.Fatal("expected GetByID to be called")
// 	}

// 	// GetActivePlan should NOT have been called.
// 	if planRepo.getActiveCalled {
// 		t.Fatal("GetActivePlan should not be called")
// 	}

// 	if result.Version != 1 {
// 		t.Errorf(
// 			"expected version 1, got %d",
// 			result.Version,
// 		)
// 	}

// 	if result.Note != "Initial version" {
// 		t.Errorf(
// 			"expected note 'Initial version', got %s",
// 			result.Note,
// 		)
// 	}
// }

// // ============================================================
// // TEST 2
// // ActivePlanVersionID does NOT exist
// // ============================================================

// func TestGetActivePlan_WithoutActivePlanVersionID(t *testing.T) {

// 	daySessionID := createTestDaySessionID(t)

// 	daySession, err := daySessionEntity.NewDaySession(
// 		"trip-1",
// 		"2026-08-20",
// 		"09:00",
// 		"Hotel",
// 	)

// 	if err != nil {
// 		t.Fatalf("failed to create day session: %v", err)
// 	}

// 	// We do NOT call SetActivePlanVersionID().
// 	// Therefore ActivePlanVersionID() == nil.

// 	planVersion, err := entity.NewPlanVersion(
// 		daySessionID,
// 		1,
// 		"Initial version",
// 	)

// 	if err != nil {
// 		t.Fatalf("failed to create plan version: %v", err)
// 	}

// 	planRepo := &fakePlanVersionRepository{
// 		planVersion: planVersion,
// 	}

// 	dayRepo := &fakeDaySessionRepository{
// 		daySession: daySession,
// 	}

// 	service := NewGetByIDPlanVersionService(
// 		planRepo,
// 		dayRepo,
// 	)

// 	result, err := service.GetActivePlan(
// 		context.Background(),
// 		daySessionID,
// 	)

// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	// Since there is no active plan ID,
// 	// service should call GetActivePlan().
// 	if !planRepo.getActiveCalled {
// 		t.Fatal("expected GetActivePlan to be called")
// 	}

// 	// GetByID should NOT be called.
// 	if planRepo.getByIDCalled {
// 		t.Fatal("GetByID should not be called")
// 	}

// 	if result.Version != 1 {
// 		t.Errorf(
// 			"expected version 1, got %d",
// 			result.Version,
// 		)
// 	}
// }

// // ============================================================
// // TEST 3
// // DaySessionRepository returns an error
// // ============================================================

// func TestGetActivePlan_DaySessionError(t *testing.T) {

// 	expectedError := errors.New("day session not found")

// 	planRepo := &fakePlanVersionRepository{}

// 	dayRepo := &fakeDaySessionRepository{
// 		getError: expectedError,
// 	}

// 	service := NewGetByIDPlanVersionService(
// 		planRepo,
// 		dayRepo,
// 	)

// 	daySessionID := createTestDaySessionID(t)

// 	_, err := service.GetActivePlan(
// 		context.Background(),
// 		daySessionID,
// 	)

// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}

// 	if !errors.Is(err, expectedError) {
// 		t.Errorf(
// 			"expected %v, got %v",
// 			expectedError,
// 			err,
// 		)
// 	}

// 	// Since getting the DaySession failed,
// 	// no PlanVersion repository method should be called.
// 	if planRepo.getByIDCalled {
// 		t.Fatal("GetByID should not be called")
// 	}

// 	if planRepo.getActiveCalled {
// 		t.Fatal("GetActivePlan should not be called")
// 	}
// }

// // ============================================================
// // TEST 4
// // ActivePlanVersionID exists,
// // but GetByID returns an error.
// // ============================================================

// func TestGetActivePlan_GetByIDError(t *testing.T) {

// 	daySessionID := createTestDaySessionID(t)
// 	planVersionID := createTestPlanVersionID(t)

// 	daySession, err := daySessionEntity.NewDaySession(
// 		"trip-1",
// 		"2026-08-20",
// 		"09:00",
// 		"Hotel",
// 	)

// 	if err != nil {
// 		t.Fatalf("failed to create day session: %v", err)
// 	}

// 	daySession.SetActivePlanVersionID(planVersionID)

// 	expectedError := errors.New("plan version not found")

// 	planRepo := &fakePlanVersionRepository{
// 		getByIDError: expectedError,
// 	}

// 	dayRepo := &fakeDaySessionRepository{
// 		daySession: daySession,
// 	}

// 	service := NewGetByIDPlanVersionService(
// 		planRepo,
// 		dayRepo,
// 	)

// 	_, err = service.GetActivePlan(
// 		context.Background(),
// 		daySessionID,
// 	)

// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}

// 	if !errors.Is(err, expectedError) {
// 		t.Errorf(
// 			"expected %v, got %v",
// 			expectedError,
// 			err,
// 		)
// 	}
// }

// // ============================================================
// // TEST 5
// // No ActivePlanVersionID,
// // but GetActivePlan returns an error.
// // ============================================================

// func TestGetActivePlan_GetActivePlanError(t *testing.T) {

// 	daySessionID := createTestDaySessionID(t)

// 	daySession, err := daySessionEntity.NewDaySession(
// 		"trip-1",
// 		"2026-08-20",
// 		"09:00",
// 		"Hotel",
// 	)

// 	if err != nil {
// 		t.Fatalf("failed to create day session: %v", err)
// 	}

// 	expectedError := errors.New("active plan not found")

// 	planRepo := &fakePlanVersionRepository{
// 		getActiveError: expectedError,
// 	}

// 	dayRepo := &fakeDaySessionRepository{
// 		daySession: daySession,
// 	}

// 	service := NewGetByIDPlanVersionService(
// 		planRepo,
// 		dayRepo,
// 	)

// 	_, err = service.GetActivePlan(
// 		context.Background(),
// 		daySessionID,
// 	)

// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}

// 	if !errors.Is(err, expectedError) {
// 		t.Errorf(
// 			"expected %v, got %v",
// 			expectedError,
// 			err,
// 		)
// 	}
// }