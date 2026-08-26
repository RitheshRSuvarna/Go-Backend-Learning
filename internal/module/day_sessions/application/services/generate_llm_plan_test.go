package services

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	"common"

// 	"day_session/domain/entity"
// 	"day_session/domain/port"

// 	plansentity "plans/domain/entity"
// 	plansrepo "plans/domain/repository"

// 	tripentity "trip/domain/entity"
// 	triprepo "trip/domain/repository"
// )

// // ============================================================
// // Fake Plan Generator
// // ============================================================

// type fakePlanGenerator struct {
// 	generatePlanCalled bool
// 	replanCalled       bool

// 	generatePlanResult port.PlanResponse
// 	generatePlanError  error

// 	replanResult port.PlanResponse
// 	replanError  error

// 	lastPlanRequest   port.PlanRequest
// 	lastReplanRequest port.RePlanRequest
// }

// func (f *fakePlanGenerator) GeneratePlan(
// 	ctx context.Context,
// 	req port.PlanRequest,
// ) (port.PlanResponse, error) {

// 	f.generatePlanCalled = true
// 	f.lastPlanRequest = req

// 	return f.generatePlanResult, f.generatePlanError
// }

// func (f *fakePlanGenerator) Replan(
// 	ctx context.Context,
// 	req port.RePlanRequest,
// ) (port.PlanResponse, error) {

// 	f.replanCalled = true
// 	f.lastReplanRequest = req

// 	return f.replanResult, f.replanError
// }

// // ============================================================
// // Fake DaySession Repository
// // ============================================================

// type fakeDaySessionRepository struct {
// 	daySession *entity.DaySession
// 	getError   error
// }

// func (f *fakeDaySessionRepository) GetDaySession(
// 	ctx context.Context,
// 	id common.DaySessionID,
// ) (*entity.DaySession, error) {
// 	if f.getError != nil {
// 		return nil, f.getError
// 	}

// 	return f.daySession, nil
// }

// // Add your other DaySessionRepository methods here if required
// // by your actual interface.

// // ============================================================
// // Fake Trip Repository
// // ============================================================

// type fakeTripRepository struct {
// 	trip      *tripentity.Trip
// 	getByIDErr error
// }

// func (f *fakeTripRepository) GetByID(
// 	ctx context.Context,
// 	id common.TripID,
// ) (*tripentity.Trip, error) {
// 	if f.getByIDErr != nil {
// 		return nil, f.getByIDErr
// 	}

// 	return f.trip, nil
// }

// // Add Create/List if required by your actual TripRepository.

// // ============================================================
// // Fake PlanVersion Repository
// // ============================================================

// type fakePlanVersionRepository struct {
// 	latestVersion int
// 	getLatestErr  error

// 	activePlan *plansentity.PlanVersion
// 	getActiveErr error

// 	createCalled bool
// 	createdPlan  *plansentity.PlanVersion
// 	createError  error
// }

// func (f *fakePlanVersionRepository) GetLatestVersion(
// 	ctx context.Context,
// 	id common.DaySessionID,
// ) (int, error) {

// 	if f.getLatestErr != nil {
// 		return 0, f.getLatestErr
// 	}

// 	return f.latestVersion, nil
// }

// func (f *fakePlanVersionRepository) GetActivePlan(
// 	ctx context.Context,
// 	id common.DaySessionID,
// ) (*plansentity.PlanVersion, error) {

// 	if f.getActiveErr != nil {
// 		return nil, f.getActiveErr
// 	}

// 	return f.activePlan, nil
// }

// func (f *fakePlanVersionRepository) Create(
// 	ctx context.Context,
// 	planVersion *plansentity.PlanVersion,
// ) error {

// 	f.createCalled = true
// 	f.createdPlan = planVersion

// 	return f.createError
// }

// // Add GetByID/ListPlanVersion if required by your actual interface.

// // ============================================================
// // Fake PlanStop Repository
// // ============================================================

// type fakePlanStopRepository struct {
// 	createCalled bool
// 	createdStops []*plansentity.PlanStop
// 	createError  error

// 	existingStops []*plansentity.PlanStop
// 	listError     error
// }

// func (f *fakePlanStopRepository) Create(
// 	ctx context.Context,
// 	planStop *plansentity.PlanStop,
// ) error {

// 	f.createCalled = true
// 	f.createdStops = append(f.createdStops, planStop)

// 	return f.createError
// }

// func (f *fakePlanStopRepository) ListStop(
// 	ctx context.Context,
// 	id common.PlanVersionID,
// ) ([]*plansentity.PlanStop, error) {

// 	if f.listError != nil {
// 		return nil, f.listError
// 	}

// 	return f.existingStops, nil
// }

// // ============================================================
// // Helper: create DaySession
// // ============================================================

// func createTestDaySession(t *testing.T) *entity.DaySession {
// 	t.Helper()

// 	daySession, err := entity.NewDaySession(
// 		"trip-1",
// 		"2026-08-20",
// 		"09:00",
// 		"Hotel",
// 	)

// 	if err != nil {
// 		t.Fatalf("failed to create test day session: %v", err)
// 	}

// 	return daySession
// }

// // ============================================================
// // Helper: create Trip
// // ============================================================

// func createTestTrip(t *testing.T) *tripentity.Trip {
// 	t.Helper()

// 	trip, err := tripentity.NewTrip(
// 		"Goa",
// 		"2026-08-20",
// 		"2026-08-25",
// 		2,
// 	)

// 	if err != nil {
// 		t.Fatalf("failed to create test trip: %v", err)
// 	}

// 	return trip
// }

// // ============================================================
// // GeneratePlan - Success
// // ============================================================

// func TestGenerateLLMPlanService_GeneratePlan_Success(t *testing.T) {

// 	daySession := createTestDaySession(t)
// 	trip := createTestTrip(t)

// 	generator := &fakePlanGenerator{
// 		generatePlanResult: port.PlanResponse{
// 			Stops: []port.PlanStopResponse{
// 				{
// 					Position:         1,
// 					Title:            "Baga Beach",
// 					CategoryLabel:    "Beach",
// 					ImageURL:         "https://example.com/baga.jpg",
// 					PlannedArrival:   time.Date(2026, 8, 20, 10, 0, 0, 0, time.UTC),
// 					PlannedDeparture: time.Date(2026, 8, 20, 12, 0, 0, 0, time.UTC),
// 					TravelMinutes:    20,
// 					StayMinutes:      120,
// 				},
// 			},
// 		},
// 	}

// 	dayRepo := &fakeDaySessionRepository{
// 		daySession: daySession,
// 	}

// 	tripRepo := &fakeTripRepository{
// 		trip: trip,
// 	}

// 	versionRepo := &fakePlanVersionRepository{
// 		latestVersion: 0,
// 	}

// 	stopRepo := &fakePlanStopRepository{}

// 	service := NewGenerateLLMPlanService(
// 		generator,
// 		dayRepo,
// 		tripRepo,
// 		versionRepo,
// 		stopRepo,
// 	)

// 	result, err := service.GeneratePlan(
// 		context.Background(),
// 		"day-session-1",
// 	)

// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	if !generator.generatePlanCalled {
// 		t.Fatal("expected GeneratePlan to be called")
// 	}

// 	if !versionRepo.createCalled {
// 		t.Fatal("expected PlanVersion.Create to be called")
// 	}

// 	if versionRepo.createdPlan.Version() != 1 {
// 		t.Errorf(
// 			"expected version 1, got %d",
// 			versionRepo.createdPlan.Version(),
// 		)
// 	}

// 	if !stopRepo.createCalled {
// 		t.Fatal("expected PlanStop.Create to be called")
// 	}

// 	if len(stopRepo.createdStops) != 1 {
// 		t.Fatalf(
// 			"expected 1 plan stop, got %d",
// 			len(stopRepo.createdStops),
// 		)
// 	}

// 	if result.Stops[0].Title != "Baga Beach" {
// 		t.Errorf(
// 			"expected Baga Beach, got %s",
// 			result.Stops[0].Title,
// 		)
// 	}
// }

// // ============================================================
// // GeneratePlan - Empty DaySession ID
// // ============================================================

// func TestGenerateLLMPlanService_GeneratePlan_EmptyDaySessionID(
// 	t *testing.T,
// ) {

// 	service := NewGenerateLLMPlanService(
// 		&fakePlanGenerator{},
// 		&fakeDaySessionRepository{},
// 		&fakeTripRepository{},
// 		&fakePlanVersionRepository{},
// 		&fakePlanStopRepository{},
// 	)

// 	result, err := service.GeneratePlan(
// 		context.Background(),
// 		"",
// 	)

// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}

// 	if result != (port.PlanResponse{}) {
// 		t.Fatal("expected empty PlanResponse")
// 	}

// 	expected := "day session id is required"

// 	if err.Error() != expected {
// 		t.Errorf(
// 			"expected error %q, got %q",
// 			expected,
// 			err.Error(),
// 		)
// 	}
// }

// // ============================================================
// // GeneratePlan - DaySession repository error
// // ============================================================

// func TestGenerateLLMPlanService_GeneratePlan_DaySessionError(
// 	t *testing.T,
// ) {

// 	expectedError := errors.New("day session not found")

// 	service := NewGenerateLLMPlanService(
// 		&fakePlanGenerator{},
// 		&fakeDaySessionRepository{
// 			getError: expectedError,
// 		},
// 		&fakeTripRepository{},
// 		&fakePlanVersionRepository{},
// 		&fakePlanStopRepository{},
// 	)

// 	result, err := service.GeneratePlan(
// 		context.Background(),
// 		"day-session-1",
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

// 	if result != (port.PlanResponse{}) {
// 		t.Fatal("expected empty PlanResponse")
// 	}
// }

// // ============================================================
// // GeneratePlan - Trip repository error
// // ============================================================

// func TestGenerateLLMPlanService_GeneratePlan_TripError(
// 	t *testing.T,
// ) {

// 	expectedError := errors.New("trip not found")

// 	service := NewGenerateLLMPlanService(
// 		&fakePlanGenerator{},
// 		&fakeDaySessionRepository{
// 			daySession: createTestDaySession(t),
// 		},
// 		&fakeTripRepository{
// 			getByIDErr: expectedError,
// 		},
// 		&fakePlanVersionRepository{},
// 		&fakePlanStopRepository{},
// 	)

// 	_, err := service.GeneratePlan(
// 		context.Background(),
// 		"day-session-1",
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
// // GeneratePlan - LLM error
// // ============================================================

// func TestGenerateLLMPlanService_GeneratePlan_GeneratorError(
// 	t *testing.T,
// ) {

// 	expectedError := errors.New("LLM unavailable")

// 	service := NewGenerateLLMPlanService(
// 		&fakePlanGenerator{
// 			generatePlanError: expectedError,
// 		},
// 		&fakeDaySessionRepository{
// 			daySession: createTestDaySession(t),
// 		},
// 		&fakeTripRepository{
// 			trip: createTestTrip(t),
// 		},
// 		&fakePlanVersionRepository{},
// 		&fakePlanStopRepository{},
// 	)

// 	_, err := service.GeneratePlan(
// 		context.Background(),
// 		"day-session-1",
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
// // GeneratePlan - First plan when no previous version exists
// // ============================================================

// func TestGenerateLLMPlanService_GeneratePlan_FirstVersion(
// 	t *testing.T,
// ) {

// 	daySession := createTestDaySession(t)
// 	trip := createTestTrip(t)

// 	generator := &fakePlanGenerator{
// 		generatePlanResult: port.PlanResponse{},
// 	}

// 	versionRepo := &fakePlanVersionRepository{
// 		getLatestErr: pgx.ErrNoRows,
// 	}

// 	service := NewGenerateLLMPlanService(
// 		generator,
// 		&fakeDaySessionRepository{
// 			daySession: daySession,
// 		},
// 		&fakeTripRepository{
// 			trip: trip,
// 		},
// 		versionRepo,
// 		&fakePlanStopRepository{},
// 	)

// 	_, err := service.GeneratePlan(
// 		context.Background(),
// 		"day-session-1",
// 	)

// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	if versionRepo.createdPlan == nil {
// 		t.Fatal("expected plan version to be created")
// 	}

// 	if versionRepo.createdPlan.Version() != 1 {
// 		t.Errorf(
// 			"expected first version to be 1, got %d",
// 			versionRepo.createdPlan.Version(),
// 		)
// 	}
// }

// // ============================================================
// // Replan - Empty ID
// // ============================================================

// func TestGenerateLLMPlanService_Replan_EmptyDaySessionID(
// 	t *testing.T,
// ) {

// 	service := NewGenerateLLMPlanService(
// 		&fakePlanGenerator{},
// 		&fakeDaySessionRepository{},
// 		&fakeTripRepository{},
// 		&fakePlanVersionRepository{},
// 		&fakePlanStopRepository{},
// 	)

// 	result, err := service.Replan(
// 		context.Background(),
// 		"",
// 	)

// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}

// 	if result != (port.PlanResponse{}) {
// 		t.Fatal("expected empty PlanResponse")
// 	}
// }

// // ============================================================
// // Replan - Success
// // ============================================================

// func TestGenerateLLMPlanService_Replan_Success(t *testing.T) {

// 	daySession := createTestDaySession(t)
// 	trip := createTestTrip(t)

// 	activeVersion, err := plansentity.NewPlanVersion(
// 		common.MustDaySessionID("day-session-1"),
// 		1,
// 		"Initial version",
// 	)

// 	if err != nil {
// 		t.Fatalf(
// 			"failed to create plan version: %v",
// 			err,
// 		)
// 	}

// 	stop, err := plansentity.NewPlanStop(
// 		activeVersion.ID(),
// 		1,
// 		"Baga Beach",
// 		"Beach",
// 		"https://example.com/baga.jpg",
// 		time.Date(2026, 8, 20, 10, 0, 0, 0, time.UTC),
// 		time.Date(2026, 8, 20, 12, 0, 0, 0, time.UTC),
// 		20,
// 		120,
// 		"Low",
// 	)

// 	if err != nil {
// 		t.Fatalf(
// 			"failed to create plan stop: %v",
// 			err,
// 		)
// 	}

// 	generator := &fakePlanGenerator{
// 		replanResult: port.PlanResponse{
// 			Stops: []port.PlanStopResponse{
// 				{
// 					Position:         1,
// 					Title:            "Fort Aguada",
// 					CategoryLabel:    "Sightseeing",
// 					ImageURL:         "https://example.com/fort.jpg",
// 					PlannedArrival:   time.Date(2026, 8, 20, 10, 0, 0, 0, time.UTC),
// 					PlannedDeparture: time.Date(2026, 8, 20, 12, 0, 0, 0, time.UTC),
// 					TravelMinutes:    30,
// 					StayMinutes:      120,
// 				},
// 			},
// 		},
// 	}

// 	versionRepo := &fakePlanVersionRepository{
// 		activePlan:    activeVersion,
// 		latestVersion: 1,
// 	}

// 	stopRepo := &fakePlanStopRepository{
// 		existingStops: []*plansentity.PlanStop{
// 			stop,
// 		},
// 	}

// 	service := NewGenerateLLMPlanService(
// 		generator,
// 		&fakeDaySessionRepository{
// 			daySession: daySession,
// 		},
// 		&fakeTripRepository{
// 			trip: trip,
// 		},
// 		versionRepo,
// 		stopRepo,
// 	)

// 	result, err := service.Replan(
// 		context.Background(),
// 		"day-session-1",
// 	)

// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	if !generator.replanCalled {
// 		t.Fatal("expected Replan to be called")
// 	}

// 	if len(generator.lastReplanRequest.ExistingPlanStops) != 1 {
// 		t.Fatalf(
// 			"expected 1 existing stop, got %d",
// 			len(generator.lastReplanRequest.ExistingPlanStops),
// 		)
// 	}

// 	if generator.lastReplanRequest.ExistingPlanStops[0].Title != "Baga Beach" {
// 		t.Errorf(
// 			"expected Baga Beach, got %s",
// 			generator.lastReplanRequest.ExistingPlanStops[0].Title,
// 		)
// 	}

// 	if !versionRepo.createCalled {
// 		t.Fatal("expected new PlanVersion to be created")
// 	}

// 	if versionRepo.createdPlan.Version() != 2 {
// 		t.Errorf(
// 			"expected new version 2, got %d",
// 			versionRepo.createdPlan.Version(),
// 		)
// 	}

// 	if len(stopRepo.createdStops) != 1 {
// 		t.Fatalf(
// 			"expected 1 new stop, got %d",
// 			len(stopRepo.createdStops),
// 		)
// 	}

// 	if result.Stops[0].Title != "Fort Aguada" {
// 		t.Errorf(
// 			"expected Fort Aguada, got %s",
// 			result.Stops[0].Title,
// 		)
// 	}
// }

// // ============================================================
// // Replan - GetActivePlan error
// // ============================================================

// func TestGenerateLLMPlanService_Replan_ActivePlanError(
// 	t *testing.T,
// ) {

// 	expectedError := errors.New("active plan not found")

// 	service := NewGenerateLLMPlanService(
// 		&fakePlanGenerator{},
// 		&fakeDaySessionRepository{
// 			daySession: createTestDaySession(t),
// 		},
// 		&fakeTripRepository{
// 			trip: createTestTrip(t),
// 		},
// 		&fakePlanVersionRepository{
// 			getActiveErr: expectedError,
// 		},
// 		&fakePlanStopRepository{},
// 	)

// 	_, err := service.Replan(
// 		context.Background(),
// 		"day-session-1",
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
// // Replan - ListStop error
// // ============================================================

// func TestGenerateLLMPlanService_Replan_ListStopError(
// 	t *testing.T,
// ) {

// 	expectedError := errors.New("failed to get plan stops")

// 	activeVersion, err := plansentity.NewPlanVersion(
// 		common.MustDaySessionID("day-session-1"),
// 		1,
// 		"Initial version",
// 	)

// 	if err != nil {
// 		t.Fatalf("failed to create plan version: %v", err)
// 	}

// 	service := NewGenerateLLMPlanService(
// 		&fakePlanGenerator{},
// 		&fakeDaySessionRepository{
// 			daySession: createTestDaySession(t),
// 		},
// 		&fakeTripRepository{
// 			trip: createTestTrip(t),
// 		},
// 		&fakePlanVersionRepository{
// 			activePlan: activeVersion,
// 		},
// 		&fakePlanStopRepository{
// 			listError: expectedError,
// 		},
// 	)

// 	_, err = service.Replan(
// 		context.Background(),
// 		"day-session-1",
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
// // Replan - Generator error
// // ============================================================

// func TestGenerateLLMPlanService_Replan_GeneratorError(
// 	t *testing.T,
// ) {

// 	expectedError := errors.New("LLM replan failed")

// 	activeVersion, err := plansentity.NewPlanVersion(
// 		common.MustDaySessionID("day-session-1"),
// 		1,
// 		"Initial version",
// 	)

// 	if err != nil {
// 		t.Fatalf("failed to create plan version: %v", err)
// 	}

// 	service := NewGenerateLLMPlanService(
// 		&fakePlanGenerator{
// 			replanError: expectedError,
// 		},
// 		&fakeDaySessionRepository{
// 			daySession: createTestDaySession(t),
// 		},
// 		&fakeTripRepository{
// 			trip: createTestTrip(t),
// 		},
// 		&fakePlanVersionRepository{
// 			activePlan:    activeVersion,
// 			latestVersion: 1,
// 		},
// 		&fakePlanStopRepository{},
// 	)

// 	_, err = service.Replan(
// 		context.Background(),
// 		"day-session-1",
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