package services

// import (
// 	"context"
// "errors"
// 	"testing"

// 	"common"
// 	"day_session/domain/entity"
// )

// type fakeSetActivePlanRepository struct {
// 	updateActivePlanCalled bool
// 	updateActivePlanError  error
// }

// func (f *fakeSetActivePlanRepository) CreateDaySession(
// 	ctx context.Context,
// 	daySession *entity.DaySession,
// ) error {
// 	return nil
// }

// func (f *fakeSetActivePlanRepository) GetDaySession(
// 	ctx context.Context,
// 	id common.DaySessionID,
// ) (*entity.DaySession, error) {
// 	return nil, nil
// }

// func (f *fakeSetActivePlanRepository) UpdateActivePlan(
// 	ctx context.Context,
// 	daySessionID common.DaySessionID,
// 	planVersionID common.PlanVersionID,
// ) error {
// 	f.updateActivePlanCalled = true

// 	return f.updateActivePlanError
// }

// func TestSetActivePlanService_Success(t *testing.T) {
// 	repo := &fakeSetActivePlanRepository{}

// 	service := NewSetActivePlanService(repo)

// 	daySessionID, err := common.NewDaySessionID("day-session-1")
// 	if err != nil {
// 		t.Fatalf("failed to create day session ID: %v", err)
// 	}

// 	planVersionID, err := common.NewPlanVersionID("plan-version-1")
// 	if err != nil {
// 		t.Fatalf("failed to create plan version ID: %v", err)
// 	}

// 	err = service.UpdateActivePlan(
// 		context.Background(),
// 		daySessionID,
// 		planVersionID,
// 	)

// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	if !repo.updateActivePlanCalled {
// 		t.Fatal("expected UpdateActivePlan to be called")
// 	}
// }

// func TestSetActivePlanService_RepositoryError(t *testing.T) {
// 	expectedError := errors.New("database error")

// 	repo := &fakeSetActivePlanRepository{
// 		updateActivePlanError: expectedError,
// 	}

// 	service := NewSetActivePlanService(repo)

// 	daySessionID, err := common.NewDaySessionID("day-session-1")
// 	if err != nil {
// 		t.Fatalf("failed to create day session ID: %v", err)
// 	}

// 	planVersionID, err := common.NewPlanVersionID("plan-version-1")
// 	if err != nil {
// 		t.Fatalf("failed to create plan version ID: %v", err)
// 	}

// 	err = service.UpdateActivePlan(
// 		context.Background(),
// 		daySessionID,
// 		planVersionID,
// 	)

// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}

// 	if !errors.Is(err, expectedError) {
// 		t.Errorf(
// 			"expected error %v, got %v",
// 			expectedError,
// 			err,
// 		)
// 	}

// 	if !repo.updateActivePlanCalled {
// 		t.Fatal("expected UpdateActivePlan to be called")
// 	}
// }