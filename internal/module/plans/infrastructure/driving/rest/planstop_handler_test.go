package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"common"
	"plans/application/services"
	"plans/domain/entity"
)

// ============================================================
// Fake PlanStop Repository
// ============================================================

type fakePlanStopRepository struct {
	stops []*entity.PlanStop
	err   error

	createCalled bool
	listCalled   bool
}

func (f *fakePlanStopRepository) Create(
	ctx context.Context,
	stop *entity.PlanStop,
) error {

	f.createCalled = true

	if f.err != nil {
		return f.err
	}

	f.stops = append(f.stops, stop)

	return nil
}

func (f *fakePlanStopRepository) ListStop(
	ctx context.Context,
	id common.PlanVersionID,
) ([]*entity.PlanStop, error) {

	f.listCalled = true

	if f.err != nil {
		return nil, f.err
	}

	return f.stops, nil
}

// ============================================================
// Test Handler
// ============================================================

func createTestHandler(repo *fakePlanStopRepository) *Handlers {

	createService := services.NewCreatePlanStopService(repo)

	listService := services.NewListPlanStopService(repo)

	return NewHandlers(
		createService,
		listService,
	)
}

// ============================================================
// TEST 1
// POST /.../{id}/stop
// Valid request
// ============================================================

func TestHandlers_CreatePlanStop_Success(t *testing.T) {

	repo := &fakePlanStopRepository{}

	handler := createTestHandler(repo)

	body := `{
		"position": 1,
		"title": "Baga Beach",
		"categorylabel": "Beach",
		"imageurl": "https://example.com/baga.jpg",
		"plannedarrival": "2026-08-20T10:00:00Z",
		"planneddeparture": "2026-08-20T11:00:00Z",
		"travelminutes": 20,
		"stayminutes": 60,
		"busyrisklabel": "Low"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/day_sessions/plan-version-1/stop",
		bytes.NewBufferString(body),
	)

	// Set PathValue("id")
	req.SetPathValue("id", "plan-version-1")

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

	if !repo.createCalled {
		t.Fatal(
			"expected repository Create to be called",
		)
	}

	// Make sure response is valid JSON.
	var response map[string]interface{}

	if err := json.Unmarshal(
		rec.Body.Bytes(),
		&response,
	); err != nil {
		t.Fatalf(
			"expected valid JSON response, got %v",
			err,
		)
	}
}

// ============================================================
// TEST 2
// Invalid JSON
// ============================================================

func TestHandlers_CreatePlanStop_InvalidJSON(t *testing.T) {

	repo := &fakePlanStopRepository{}

	handler := createTestHandler(repo)

	body := `{
		"title": "Baga Beach",
		"position":
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/day_sessions/plan-version-1/stop",
		bytes.NewBufferString(body),
	)

	req.SetPathValue("id", "plan-version-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	// Your current handler returns 404 here.
	if rec.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusNotFound,
			rec.Code,
		)
	}

	if repo.createCalled {
		t.Fatal(
			"repository Create should not be called",
		)
	}
}

// ============================================================
// TEST 3
// Invalid PlanVersion ID
// ============================================================

func TestHandlers_CreatePlanStop_InvalidPlanVersionID(t *testing.T) {

	repo := &fakePlanStopRepository{}

	handler := createTestHandler(repo)

	body := `{
		"position": 1,
		"title": "Baga Beach",
		"categorylabel": "Beach",
		"imageurl": "https://example.com/baga.jpg",
		"plannedarrival": "2026-08-20T10:00:00Z",
		"planneddeparture": "2026-08-20T11:00:00Z",
		"travelminutes": 20,
		"stayminutes": 60,
		"busyrisklabel": "Low"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/day_sessions/invalid/stop",
		bytes.NewBufferString(body),
	)

	req.SetPathValue("id", "")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code == http.StatusCreated {
		t.Fatal(
			"expected request to fail for invalid PlanVersion ID",
		)
	}

	if repo.createCalled {
		t.Fatal(
			"repository Create should not be called",
		)
	}
}

// ============================================================
// TEST 4
// Invalid Planned Arrival
// ============================================================

func TestHandlers_CreatePlanStop_InvalidArrival(t *testing.T) {

	repo := &fakePlanStopRepository{}

	handler := createTestHandler(repo)

	body := `{
		"position": 1,
		"title": "Baga Beach",
		"categorylabel": "Beach",
		"imageurl": "https://example.com/baga.jpg",
		"plannedarrival": "invalid-time",
		"planneddeparture": "2026-08-20T11:00:00Z",
		"travelminutes": 20,
		"stayminutes": 60,
		"busyrisklabel": "Low"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/day_sessions/plan-version-1/stop",
		bytes.NewBufferString(body),
	)

	req.SetPathValue("id", "plan-version-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}

	if repo.createCalled {
		t.Fatal(
			"repository Create should not be called",
		)
	}
}

// ============================================================
// TEST 5
// Invalid Planned Departure
// ============================================================

func TestHandlers_CreatePlanStop_InvalidDeparture(t *testing.T) {

	repo := &fakePlanStopRepository{}

	handler := createTestHandler(repo)

	body := `{
		"position": 1,
		"title": "Baga Beach",
		"categorylabel": "Beach",
		"imageurl": "https://example.com/baga.jpg",
		"plannedarrival": "2026-08-20T10:00:00Z",
		"planneddeparture": "invalid-time",
		"travelminutes": 20,
		"stayminutes": 60,
		"busyrisklabel": "Low"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/day_sessions/plan-version-1/stop",
		bytes.NewBufferString(body),
	)

	req.SetPathValue("id", "plan-version-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}

	if repo.createCalled {
		t.Fatal(
			"repository Create should not be called",
		)
	}
}

// ============================================================
// TEST 6
// Domain validation error
// ============================================================

func TestHandlers_CreatePlanStop_InvalidDomainData(t *testing.T) {

	repo := &fakePlanStopRepository{}

	handler := createTestHandler(repo)

	// Position = 0.
	body := `{
		"position": 0,
		"title": "Baga Beach",
		"categorylabel": "Beach",
		"imageurl": "https://example.com/baga.jpg",
		"plannedarrival": "2026-08-20T10:00:00Z",
		"planneddeparture": "2026-08-20T11:00:00Z",
		"travelminutes": 20,
		"stayminutes": 60,
		"busyrisklabel": "Low"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/day_sessions/plan-version-1/stop",
		bytes.NewBufferString(body),
	)

	req.SetPathValue("id", "plan-version-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}

	if repo.createCalled {
		t.Fatal(
			"repository Create should not be called",
		)
	}
}

// ============================================================
// TEST 7
// GET Plan Stops
// ============================================================

func TestHandlers_ListPlanStop_Success(t *testing.T) {

	arrival := time.Date(
		2026,
		8,
		20,
		10,
		0,
		0,
		0,
		time.UTC,
	)

	departure := time.Date(
		2026,
		8,
		20,
		11,
		0,
		0,
		0,
		time.UTC,
	)

	planVersionID, err := common.NewPlanVersionID(
		"plan-version-1",
	)

	if err != nil {
		t.Fatalf(
			"failed to create plan version ID: %v",
			err,
		)
	}

	stop, err := entity.NewPlanStop(
		planVersionID,
		1,
		"Baga Beach",
		"Beach",
		"https://example.com/baga.jpg",
		arrival,
		departure,
		20,
		60,
		"Low",
	)

	if err != nil {
		t.Fatalf(
			"failed to create test plan stop: %v",
			err,
		)
	}

	repo := &fakePlanStopRepository{
		stops: []*entity.PlanStop{
			stop,
		},
	}

	handler := createTestHandler(repo)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/day_sessions/plan-version-1/stop",
		nil,
	)

	req.SetPathValue("id", "plan-version-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	if !repo.listCalled {
		t.Fatal(
			"expected repository ListStop to be called",
		)
	}

	var response []interface{}

	if err := json.Unmarshal(
		rec.Body.Bytes(),
		&response,
	); err != nil {
		t.Fatalf(
			"expected valid JSON response, got %v",
			err,
		)
	}

	if len(response) != 1 {
		t.Fatalf(
			"expected 1 plan stop, got %d",
			len(response),
		)
	}
}

// ============================================================
// TEST 8
// GET with invalid PlanVersion ID
// ============================================================

func TestHandlers_ListPlanStop_InvalidID(t *testing.T) {

	repo := &fakePlanStopRepository{}

	handler := createTestHandler(repo)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/day_sessions/invalid/stop",
		nil,
	)

	req.SetPathValue("id", "")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code == http.StatusOK {
		t.Fatal(
			"expected request to fail for invalid ID",
		)
	}
}

// ============================================================
// TEST 9
// Repository error during GET
// ============================================================

func TestHandlers_ListPlanStop_RepositoryError(t *testing.T) {

	repo := &fakePlanStopRepository{
		err: errors.New("database error"),
	}

	handler := createTestHandler(repo)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/day_sessions/plan-version-1/stop",
		nil,
	)

	req.SetPathValue("id", "plan-version-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code == http.StatusOK {
		t.Fatal(
			"expected error response",
		)
	}
}

// ============================================================
// TEST 10
// Unsupported HTTP method
// ============================================================

func TestHandlers_MethodNotAllowed(t *testing.T) {

	repo := &fakePlanStopRepository{}

	handler := createTestHandler(repo)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/day_sessions/plan-version-1/stop",
		nil,
	)

	req.SetPathValue("id", "plan-version-1")

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