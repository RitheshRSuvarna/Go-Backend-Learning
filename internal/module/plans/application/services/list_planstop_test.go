package services

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	"common"
// 	"plans/domain/entity"
// )

// // ============================================================
// // Fake PlanStopRepository
// // ============================================================

// type fakePlanStopRepository struct {
// 	stops []*entity.PlanStop
// 	err   error

// 	listStopCalled bool
// }

// func (f *fakePlanStopRepository) Create(
// 	ctx context.Context,
// 	stop *entity.PlanStop,
// ) error {
// 	return nil
// }

// func (f *fakePlanStopRepository) ListStop(
// 	ctx context.Context,
// 	id common.PlanVersionID,
// ) ([]*entity.PlanStop, error) {

// 	f.listStopCalled = true

// 	if f.err != nil {
// 		return nil, f.err
// 	}

// 	return f.stops, nil
// }

// // ============================================================
// // Helper: PlanVersionID
// // ============================================================

// func testPlanVersionIDForList(t *testing.T) common.PlanVersionID {
// 	t.Helper()

// 	id, err := common.NewPlanVersionID("plan-version-1")
// 	if err != nil {
// 		t.Fatalf("failed to create plan version ID: %v", err)
// 	}

// 	return id
// }

// // ============================================================
// // Helper: Create PlanStop
// // ============================================================

// func createTestPlanStop(t *testing.T) *entity.PlanStop {
// 	t.Helper()

// 	planVersionID := testPlanVersionIDForList(t)

// 	stop, err := entity.NewPlanStop(
// 		planVersionID,
// 		1,
// 		"Baga Beach",
// 		"beach",
// 		"https://example.com/baga.jpg",
// 		time.Date(
// 			2026,
// 			8,
// 			20,
// 			10,
// 			0,
// 			0,
// 			0,
// 			time.UTC,
// 		),
// 		time.Date(
// 			2026,
// 			8,
// 			20,
// 			12,
// 			0,
// 			0,
// 			0,
// 			time.UTC,
// 		),
// 		30,
// 		120,
// 		"low",
// 	)

// 	if err != nil {
// 		t.Fatalf("failed to create plan stop: %v", err)
// 	}

// 	return stop
// }

// // ============================================================
// // TEST 1
// // Successfully returns plan stops
// // ============================================================

// func TestListPlanStop_Success(t *testing.T) {

// 	stop := createTestPlanStop(t)

// 	repo := &fakePlanStopRepository{
// 		stops: []*entity.PlanStop{
// 			stop,
// 		},
// 	}

// 	service := NewListPlanStopService(repo)

// 	result, err := service.ListPlanStop(
// 		context.Background(),
// 		testPlanVersionIDForList(t),
// 	)

// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	if !repo.listStopCalled {
// 		t.Fatal("expected ListStop to be called")
// 	}

// 	if len(result) != 1 {
// 		t.Fatalf(
// 			"expected 1 plan stop, got %d",
// 			len(result),
// 		)
// 	}

// 	if result[0].Title != "Baga Beach" {
// 		t.Errorf(
// 			"expected title Baga Beach, got %s",
// 			result[0].Title,
// 		)
// 	}

// 	if result[0].CategoryLabel != "beach" {
// 		t.Errorf(
// 			"expected category beach, got %s",
// 			result[0].CategoryLabel,
// 		)
// 	}
// }

// // ============================================================
// // TEST 2
// // Repository returns an error
// // ============================================================

// func TestListPlanStop_RepositoryError(t *testing.T) {

// 	expectedError := errors.New(
// 		"database error",
// 	)

// 	repo := &fakePlanStopRepository{
// 		err: expectedError,
// 	}

// 	service := NewListPlanStopService(repo)

// 	result, err := service.ListPlanStop(
// 		context.Background(),
// 		testPlanVersionIDForList(t),
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

// 	if result != nil {
// 		t.Fatalf(
// 			"expected nil result, got %+v",
// 			result,
// 		)
// 	}
// }

// // ============================================================
// // TEST 3
// // Empty repository result
// // ============================================================

// func TestListPlanStop_EmptyList(t *testing.T) {

// 	repo := &fakePlanStopRepository{
// 		stops: []*entity.PlanStop{},
// 	}

// 	service := NewListPlanStopService(repo)

// 	result, err := service.ListPlanStop(
// 		context.Background(),
// 		testPlanVersionIDForList(t),
// 	)

// 	if err != nil {
// 		t.Fatalf(
// 			"expected no error, got %v",
// 			err,
// 		)
// 	}

// 	if result == nil {
// 		t.Fatal("expected empty slice, got nil")
// 	}

// 	if len(result) != 0 {
// 		t.Fatalf(
// 			"expected 0 plan stops, got %d",
// 			len(result),
// 		)
// 	}
// }

// // ============================================================
// // TEST 4
// // Multiple plan stops
// // ============================================================

// func TestListPlanStop_MultipleStops(t *testing.T) {

// 	planVersionID := testPlanVersionIDForList(t)

// 	stop1, err := entity.NewPlanStop(
// 		planVersionID,
// 		1,
// 		"Baga Beach",
// 		"beach",
// 		"https://example.com/baga.jpg",
// 		time.Date(
// 			2026,
// 			8,
// 			20,
// 			10,
// 			0,
// 			0,
// 			0,
// 			time.UTC,
// 		),
// 		time.Date(
// 			2026,
// 			8,
// 			20,
// 			12,
// 			0,
// 			0,
// 			0,
// 			time.UTC,
// 		),
// 		30,
// 		120,
// 		"low",
// 	)

// 	if err != nil {
// 		t.Fatalf("failed to create stop1: %v", err)
// 	}

// 	stop2, err := entity.NewPlanStop(
// 		planVersionID,
// 		2,
// 		"Fort Aguada",
// 		"historical",
// 		"https://example.com/fort.jpg",
// 		time.Date(
// 			2026,
// 			8,
// 			20,
// 			13,
// 			0,
// 			0,
// 			0,
// 			time.UTC,
// 		),
// 		time.Date(
// 			2026,
// 			8,
// 			20,
// 			15,
// 			0,
// 			0,
// 			0,
// 			time.UTC,
// 		),
// 		20,
// 		120,
// 		"low",
// 	)

// 	if err != nil {
// 		t.Fatalf("failed to create stop2: %v", err)
// 	}

// 	repo := &fakePlanStopRepository{
// 		stops: []*entity.PlanStop{
// 			stop1,
// 			stop2,
// 		},
// 	}

// 	service := NewListPlanStopService(repo)

// 	result, err := service.ListPlanStop(
// 		context.Background(),
// 		planVersionID,
// 	)

// 	if err != nil {
// 		t.Fatalf(
// 			"expected no error, got %v",
// 			err,
// 		)
// 	}

// 	if len(result) != 2 {
// 		t.Fatalf(
// 			"expected 2 plan stops, got %d",
// 			len(result),
// 		)
// 	}

// 	if result[0].Title != "Baga Beach" {
// 		t.Errorf(
// 			"expected first stop Baga Beach, got %s",
// 			result[0].Title,
// 		)
// 	}

// 	if result[1].Title != "Fort Aguada" {
// 		t.Errorf(
// 			"expected second stop Fort Aguada, got %s",
// 			result[1].Title,
// 		)
// 	}
// }