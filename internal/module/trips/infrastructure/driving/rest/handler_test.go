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
	"trip/application/services"
	"trip/domain/entity"
)

// ============================================================
// Fake Trip Repository
// ============================================================

type fakeTripRepository struct {
	trips []*entity.Trip
	err   error

	createCalled bool
	listCalled   bool
}

func (f *fakeTripRepository) Create(
	ctx context.Context,
	trip *entity.Trip,
) error {

	f.createCalled = true

	if f.err != nil {
		return f.err
	}

	f.trips = append(f.trips, trip)

	return nil
}

func (f *fakeTripRepository) List(
	ctx context.Context,
) ([]*entity.Trip, error) {

	f.listCalled = true

	if f.err != nil {
		return nil, f.err
	}

	return f.trips, nil
}

func (f *fakeTripRepository) GetByID(
	ctx context.Context,
	id common.TripID,
) (*entity.Trip, error) {

	if f.err != nil {
		return nil, f.err
	}

	if len(f.trips) == 0 {
		return nil, errors.New("trip not found")
	}

	return f.trips[0], nil
}

// ============================================================
// Helper: Create Handler
// ============================================================

func createTestHandler(repo *fakeTripRepository) *Handler {

	createService := services.NewCreateTripService(repo)

	listService := services.NewTripListService(repo)

	return NewHandler(
		createService,
		listService,
	)
}

// ============================================================
// TEST 1
// POST /trips with valid JSON
// ============================================================

func TestHandler_CreateTrip_Success(t *testing.T) {

	repo := &fakeTripRepository{}

	handler := createTestHandler(repo)

	body := `{
		"destination": "Goa",
		"start_date": "2026-08-20",
		"end_date": "2026-08-25",
		"travelers_count": 2
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/trips",
		bytes.NewBufferString(body),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	// Check status code.
	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusCreated,
			rec.Code,
		)
	}

	// Repository Create should have been called.
	if !repo.createCalled {
		t.Fatal("expected repository Create to be called")
	}

	// Check that response contains JSON.
	var response map[string]interface{}

	if err := json.Unmarshal(
		rec.Body.Bytes(),
		&response,
	); err != nil {
		t.Fatalf(
			"expected valid JSON response, got error: %v",
			err,
		)
	}
}

// ============================================================
// TEST 2
// POST /trips with invalid JSON
// ============================================================

func TestHandler_CreateTrip_InvalidJSON(t *testing.T) {

	repo := &fakeTripRepository{}

	handler := createTestHandler(repo)

	body := `{
		"destination": "Goa",
		"start_date": 
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/trips",
		bytes.NewBufferString(body),
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

	// Repository should not be called.
	if repo.createCalled {
		t.Fatal(
			"repository Create should not be called for invalid JSON",
		)
	}
}

// ============================================================
// TEST 3
// POST /trips with invalid domain data
// ============================================================

func TestHandler_CreateTrip_InvalidData(t *testing.T) {

	repo := &fakeTripRepository{}

	handler := createTestHandler(repo)

	// Destination is empty.
	body := `{
		"destination": "",
		"start_date": "2026-08-20",
		"end_date": "2026-08-25",
		"travelers_count": 2
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/trips",
		bytes.NewBufferString(body),
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	// NewTrip() should reject empty destination.
	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}

	// Repository should not be called.
	if repo.createCalled {
		t.Fatal(
			"repository Create should not be called when validation fails",
		)
	}
}

// ============================================================
// TEST 4
// GET /trips
// ============================================================

func TestHandler_ListTrips_Success(t *testing.T) {

	trip, err := entity.NewTrip(
		"Goa",
		"2026-08-20",
		"2026-08-25",
		2,
	)

	if err != nil {
		t.Fatalf(
			"failed to create test trip: %v",
			err,
		)
	}

	repo := &fakeTripRepository{
		trips: []*entity.Trip{
			trip,
		},
	}

	handler := createTestHandler(repo)

	req := httptest.NewRequest(
		http.MethodGet,
		"/trips",
		nil,
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

	if !repo.listCalled {
		t.Fatal(
			"expected repository List to be called",
		)
	}

	// Check response is valid JSON.
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
			"expected 1 trip, got %d",
			len(response),
		)
	}
}

// ============================================================
// TEST 5
// GET /trips with empty list
// ============================================================

func TestHandler_ListTrips_Empty(t *testing.T) {

	repo := &fakeTripRepository{
		trips: []*entity.Trip{},
	}

	handler := createTestHandler(repo)

	req := httptest.NewRequest(
		http.MethodGet,
		"/trips",
		nil,
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
			"expected valid JSON response, got %v",
			err,
		)
	}

	if len(response) != 0 {
		t.Fatalf(
			"expected 0 trips, got %d",
			len(response),
		)
	}
}

// ============================================================
// TEST 6
// Wrong URL
// ============================================================

func TestHandler_NotFound(t *testing.T) {

	repo := &fakeTripRepository{}

	handler := createTestHandler(repo)

	req := httptest.NewRequest(
		http.MethodGet,
		"/wrong-url",
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

// ============================================================
// TEST 7
// Unsupported HTTP method
// ============================================================

func TestHandler_MethodNotAllowed(t *testing.T) {

	repo := &fakeTripRepository{}

	handler := createTestHandler(repo)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/trips",
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