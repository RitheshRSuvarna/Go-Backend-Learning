package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"common"
	"day_session/domain/entity"
	// daysessionRepo "day_session/domain/repository"

	"plans/application/services"
	planEntity "plans/domain/entity"
	// planRepo "plans/domain/repository"
)

// ============================================================
// FAKE PLAN VERSION REPOSITORY
// ============================================================

type fakePlanVersionRepository struct {
	versions []*planEntity.PlanVersion
	err      error

	createCalled bool
}



func (f *fakePlanVersionRepository) Create(
	ctx context.Context,
	version *planEntity.PlanVersion,
) error {

	f.createCalled = true

	if f.err != nil {
		return f.err
	}

	f.versions = append(f.versions, version)

	return nil
}

func (f *fakePlanVersionRepository) GetActivePlan(
	ctx context.Context,
	id common.DaySessionID,
) (*planEntity.PlanVersion, error) {

	if f.err != nil {
		return nil, f.err
	}

	if len(f.versions) == 0 {
		return nil, errors.New("no active plan")
	}

	return f.versions[0], nil
}

func (f *fakePlanVersionRepository) GetByID(
	ctx context.Context,
	id common.PlanVersionID,
) (*planEntity.PlanVersion, error) {

	if f.err != nil {
		return nil, f.err
	}

	if len(f.versions) == 0 {
		return nil, errors.New("plan version not found")
	}

	return f.versions[0], nil
}

func (f *fakePlanVersionRepository) ListPlanVersion(
	ctx context.Context,
	id common.DaySessionID,
) ([]*planEntity.PlanVersion, error) {

	if f.err != nil {
		return nil, f.err
	}

	return f.versions, nil
}

func (f *fakePlanVersionRepository) GetLatestVersion(
	ctx context.Context,
	id common.DaySessionID,
) (int, error) {

	if f.err != nil {
		return 0, f.err
	}

	if len(f.versions) == 0 {
		return 0, nil
	}

	return f.versions[len(f.versions)-1].Version(), nil
}

// ============================================================
// FAKE DAY SESSION REPOSITORY
// ============================================================

type fakeDaySessionRepository struct {
	daySession *entity.DaySession
	err        error
}

func (f *fakeDaySessionRepository) CreateDaySession(
	ctx context.Context,
	daySession *entity.DaySession,
) error {

	if f.err != nil {
		return f.err
	}

	f.daySession = daySession

	return nil
}

func (f *fakeDaySessionRepository) GetDaySession(
	ctx context.Context,
	id common.DaySessionID,
) (*entity.DaySession, error) {

	if f.err != nil {
		return nil, f.err
	}

	if f.daySession == nil {
		return nil, errors.New("day session not found")
	}

	return f.daySession, nil
}

func (f *fakeDaySessionRepository) ListDaySession(
	ctx context.Context,
	id common.TripID,
) ([]*entity.DaySession, error) {

	if f.err != nil {
		return nil, f.err
	}

	if f.daySession == nil {
		return []*entity.DaySession{}, nil
	}

	return []*entity.DaySession{f.daySession}, nil
}

func (f *fakeDaySessionRepository) UpdateActivePlan(
	ctx context.Context,
	daySessionID common.DaySessionID,
	planVersionID common.PlanVersionID,
) error {

	if f.err != nil {
		return f.err
	}

	return nil
}

// ============================================================
// HELPER
// ============================================================

func createPlanVersionTestHandler(
	versionRepo *fakePlanVersionRepository,
	dayRepo *fakeDaySessionRepository,
) *Handler {

	createService :=
		services.NewCreatePlanVersionService(versionRepo)

	listService :=
		services.NewListPlanVersionService(versionRepo)

	getActiveService :=
		services.NewGetByIDPlanVersionService(
			versionRepo,
			dayRepo,
		)

	return NewHandler(
		createService,
		getActiveService,
		listService,
	)
}

// ============================================================
// TEST 1: CREATE SUCCESS
// ============================================================

func TestHandler_CreatePlanVersion_Success(t *testing.T) {

	versionRepo := &fakePlanVersionRepository{}

	dayRepo := &fakeDaySessionRepository{}

	handler := createPlanVersionTestHandler(
		versionRepo,
		dayRepo,
	)

	body := `{
		"version": 1,
		"note": "Initial version"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/day_sessions/day-session-1/plan-versions",
		bytes.NewBufferString(body),
	)

	req.SetPathValue(
		"id",
		"day-session-1",
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusCreated,
			rec.Code,
		)
	}

	if !versionRepo.createCalled {
		t.Fatal(
			"expected repository Create to be called",
		)
	}

	var response map[string]interface{}

	if err := json.Unmarshal(
		rec.Body.Bytes(),
		&response,
	); err != nil {
		t.Fatalf(
			"expected valid JSON response: %v",
			err,
		)
	}
}

// ============================================================
// TEST 2: INVALID JSON
// ============================================================

func TestHandler_CreatePlanVersion_InvalidJSON(t *testing.T) {

	versionRepo := &fakePlanVersionRepository{}

	dayRepo := &fakeDaySessionRepository{}

	handler := createPlanVersionTestHandler(
		versionRepo,
		dayRepo,
	)

	body := `{
		"version":
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/day_sessions/day-session-1/plan-versions",
		bytes.NewBufferString(body),
	)

	req.SetPathValue(
		"id",
		"day-session-1",
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}

	if versionRepo.createCalled {
		t.Fatal(
			"repository Create should not be called",
		)
	}
}

// ============================================================
// TEST 3: INVALID DAY SESSION ID
// ============================================================

func TestHandler_CreatePlanVersion_InvalidID(t *testing.T) {

	versionRepo := &fakePlanVersionRepository{}

	dayRepo := &fakeDaySessionRepository{}

	handler := createPlanVersionTestHandler(
		versionRepo,
		dayRepo,
	)

	body := `{
		"version": 1,
		"note": "Initial version"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/day_sessions//plan-versions",
		bytes.NewBufferString(body),
	)

	req.SetPathValue(
		"id",
		"",
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code == http.StatusCreated {
		t.Fatal(
			"expected invalid ID request to fail",
		)
	}

	if versionRepo.createCalled {
		t.Fatal(
			"repository Create should not be called",
		)
	}
}

// ============================================================
// TEST 4: INVALID VERSION
// ============================================================

func TestHandler_CreatePlanVersion_InvalidVersion(t *testing.T) {

	versionRepo := &fakePlanVersionRepository{}

	dayRepo := &fakeDaySessionRepository{}

	handler := createPlanVersionTestHandler(
		versionRepo,
		dayRepo,
	)

	body := `{
		"version": 0,
		"note": "Initial version"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/day_sessions/day-session-1/plan-versions",
		bytes.NewBufferString(body),
	)

	req.SetPathValue(
		"id",
		"day-session-1",
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}

	if versionRepo.createCalled {
		t.Fatal(
			"repository Create should not be called",
		)
	}
}

// ============================================================
// TEST 5: LIST SUCCESS
// ============================================================

func TestHandler_ListPlanVersion_Success(t *testing.T) {

	daySessionID, err :=
		common.NewDaySessionID("day-session-1")

	if err != nil {
		t.Fatalf("failed to create ID: %v", err)
	}

	version, err :=
		planEntity.NewPlanVersion(
			daySessionID,
			1,
			"Initial version",
		)

	if err != nil {
		t.Fatalf("failed to create version: %v", err)
	}

	versionRepo := &fakePlanVersionRepository{
		versions: []*planEntity.PlanVersion{
			version,
		},
	}

	dayRepo := &fakeDaySessionRepository{}

	handler := createPlanVersionTestHandler(
		versionRepo,
		dayRepo,
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/day_sessions/day-session-1/plan-versions",
		nil,
	)

	req.SetPathValue(
		"id",
		"day-session-1",
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	var response []interface{}

	if err := json.Unmarshal(
		rec.Body.Bytes(),
		&response,
	); err != nil {
		t.Fatalf(
			"expected valid JSON: %v",
			err,
		)
	}

	if len(response) != 1 {
		t.Fatalf(
			"expected 1 plan version, got %d",
			len(response),
		)
	}
}

// ============================================================
// TEST 6: LIST INVALID ID
// ============================================================

func TestHandler_ListPlanVersion_InvalidID(t *testing.T) {

	versionRepo := &fakePlanVersionRepository{}

	dayRepo := &fakeDaySessionRepository{}

	handler := createPlanVersionTestHandler(
		versionRepo,
		dayRepo,
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/day_sessions//plan-versions",
		nil,
	)

	req.SetPathValue(
		"id",
		"",
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code == http.StatusOK {
		t.Fatal(
			"expected invalid ID to fail",
		)
	}
}

// ============================================================
// TEST 7: ACTIVE PLAN
// ============================================================

func TestHandler_GetActivePlan_Success(t *testing.T) {

	daySessionID, err :=
		common.NewDaySessionID("day-session-1")

	if err != nil {
		t.Fatalf("failed to create day session ID: %v", err)
	}

	version, err :=
		planEntity.NewPlanVersion(
			daySessionID,
			1,
			"Initial version",
		)

	if err != nil {
		t.Fatalf("failed to create version: %v", err)
	}

	versionRepo := &fakePlanVersionRepository{
		versions: []*planEntity.PlanVersion{
			version,
		},
	}

	daySession, err :=
		entity.NewDaySession(
			"trip-1",
			"2026-08-20",
			"09:00",
			"Hotel",
		)

	if err != nil {
		t.Fatalf(
			"failed to create day session: %v",
			err,
		)
	}

	dayRepo := &fakeDaySessionRepository{
		daySession: daySession,
	}

	handler := createPlanVersionTestHandler(
		versionRepo,
		dayRepo,
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/day_sessions/day-session-1/active-plan",
		nil,
	)

	req.SetPathValue(
		"id",
		"day-session-1",
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	var response map[string]interface{}

	if err := json.Unmarshal(
		rec.Body.Bytes(),
		&response,
	); err != nil {
		t.Fatalf(
			"expected valid JSON: %v",
			err,
		)
	}
}

// ============================================================
// TEST 8: ACTIVE PLAN INVALID ID
// ============================================================

func TestHandler_GetActivePlan_InvalidID(t *testing.T) {

	versionRepo := &fakePlanVersionRepository{}

	dayRepo := &fakeDaySessionRepository{}

	handler := createPlanVersionTestHandler(
		versionRepo,
		dayRepo,
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/day_sessions//active-plan",
		nil,
	)

	req.SetPathValue(
		"id",
		"",
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code == http.StatusOK {
		t.Fatal(
			"expected invalid ID to fail",
		)
	}
}

// ============================================================
// TEST 9: METHOD NOT ALLOWED
// ============================================================

func TestHandler_MethodNotAllowed(t *testing.T) {

	versionRepo := &fakePlanVersionRepository{}

	dayRepo := &fakeDaySessionRepository{}

	handler := createPlanVersionTestHandler(
		versionRepo,
		dayRepo,
	)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/day_sessions/day-session-1/plan-versions",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusMethodNotAllowed,
			rec.Code,
		)
	}
}

// ============================================================
// TEST 10: NOT FOUND
// ============================================================

func TestHandler_NotFound(t *testing.T) {

	versionRepo := &fakePlanVersionRepository{}

	dayRepo := &fakeDaySessionRepository{}

	handler := createPlanVersionTestHandler(
		versionRepo,
		dayRepo,
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/wrong-path",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusNotFound,
			rec.Code,
		)
	}
}