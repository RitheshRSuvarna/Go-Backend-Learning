package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"common"

	"day_session/application/services"
	"day_session/domain/entity"

	eventsEntity "events/domain/entity"
	plansEntity "plans/domain/entity"
)

// ============================================================
// FAKE DAY SESSION REPOSITORY
// ============================================================

type fakeDaySessionRepository struct {
	daySessions map[string]*entity.DaySession

	createErr error
	getErr    error
	listErr   error
	updateErr error

	updatedDaySessionID common.DaySessionID
	updatedPlanVersion  common.PlanVersionID
}

func newFakeDaySessionRepository() *fakeDaySessionRepository {
	return &fakeDaySessionRepository{
		daySessions: make(map[string]*entity.DaySession),
	}
}

func (f *fakeDaySessionRepository) CreateDaySession(
	ctx context.Context,
	ds *entity.DaySession,
) error {
	if f.createErr != nil {
		return f.createErr
	}

	id, err := common.NewDaySessionID("day-session-1")
	if err != nil {
		return err
	}

	ds.SetID(id)
	f.daySessions[id.String()] = ds

	return nil
}

func (f *fakeDaySessionRepository) GetDaySession(
	ctx context.Context,
	id common.DaySessionID,
) (*entity.DaySession, error) {

	if f.getErr != nil {
		return nil, f.getErr
	}

	ds, ok := f.daySessions[id.String()]
	if !ok {
		return nil, errors.New("day session not found")
	}

	return ds, nil
}

func (f *fakeDaySessionRepository) ListDaySession(
	ctx context.Context,
	tripID common.TripID,
) ([]*entity.DaySession, error) {

	if f.listErr != nil {
		return nil, f.listErr
	}

	result := make([]*entity.DaySession, 0)

	for _, ds := range f.daySessions {
		if ds.TripID().String() == tripID.String() {
			result = append(result, ds)
		}
	}

	return result, nil
}

func (f *fakeDaySessionRepository) UpdateActivePlan(
	ctx context.Context,
	daySessionID common.DaySessionID,
	planVersionID common.PlanVersionID,
) error {

	if f.updateErr != nil {
		return f.updateErr
	}

	f.updatedDaySessionID = daySessionID
	f.updatedPlanVersion = planVersionID

	ds, ok := f.daySessions[daySessionID.String()]
	if !ok {
		return errors.New("day session not found")
	}

	ds.SetActivePlanVersionID(planVersionID)

	return nil
}

// ============================================================
// FAKE PLAN VERSION REPOSITORY
// ============================================================

type fakePlanVersionRepository struct {
	planVersions []*plansEntity.PlanVersion
}

func newFakePlanVersionRepository() *fakePlanVersionRepository {
	return &fakePlanVersionRepository{
		planVersions: make([]*plansEntity.PlanVersion, 0),
	}
}

func (f *fakePlanVersionRepository) Create(
	ctx context.Context,
	planversion *plansEntity.PlanVersion,
) error {
	f.planVersions = append(f.planVersions, planversion)
	return nil
}

func (f *fakePlanVersionRepository) GetActivePlan(
	ctx context.Context,
	id common.DaySessionID,
) (*plansEntity.PlanVersion, error) {

	if len(f.planVersions) == 0 {
		return nil, errors.New("plan version not found")
	}

	return f.planVersions[len(f.planVersions)-1], nil
}

func (f *fakePlanVersionRepository) GetByID(
	ctx context.Context,
	id common.PlanVersionID,
) (*plansEntity.PlanVersion, error) {

	for _, pv := range f.planVersions {
		if pv.ID().String() == id.String() {
			return pv, nil
		}
	}

	return nil, errors.New("plan version not found")
}

func (f *fakePlanVersionRepository) ListPlanVersion(
	ctx context.Context,
	id common.DaySessionID,
) ([]*plansEntity.PlanVersion, error) {
	return f.planVersions, nil
}

func (f *fakePlanVersionRepository) GetLatestVersion(
	ctx context.Context,
	id common.DaySessionID,
) (int, error) {

	if len(f.planVersions) == 0 {
		return 0, nil
	}

	return f.planVersions[len(f.planVersions)-1].Version(), nil
}

// ============================================================
// FAKE PLAN STOP REPOSITORY
// ============================================================
type fakePlanStopRepository struct {
	planStops []*plansEntity.PlanStop
}

func newFakePlanStopRepository() *fakePlanStopRepository {
	return &fakePlanStopRepository{
		planStops: make([]*plansEntity.PlanStop, 0),
	}
}

func (f *fakePlanStopRepository) Create(
	ctx context.Context,
	planstop *plansEntity.PlanStop,
) error {
	f.planStops = append(f.planStops, planstop)
	return nil
}

func (f *fakePlanStopRepository) GetByID(
	ctx context.Context,
	id common.PlanStopID,
) (*plansEntity.PlanStop, error) {
	for _, stop := range f.planStops {
		if stop.ID().String() == id.String() {
			return stop, nil
		}
	}

	return nil, errors.New("plan stop not found")
}

func (f *fakePlanStopRepository) ListStop(
	ctx context.Context,
	id common.PlanVersionID,
) ([]*plansEntity.PlanStop, error) {
	return f.planStops, nil
}

// ============================================================
// FAKE EVENTS REPOSITORY
// ============================================================
type fakeEventsRepository struct {
	events []*eventsEntity.Events
}

func newFakeEventsRepository() *fakeEventsRepository {
	return &fakeEventsRepository{
		events: make([]*eventsEntity.Events, 0),
	}
}

func (f *fakeEventsRepository) CreateEvents(
	ctx context.Context,
	event *eventsEntity.Events,
) error {
	f.events = append(f.events, event)
	return nil
}

func (f *fakeEventsRepository) GetEvents(
	ctx context.Context,
	id common.DaySessionID,
) ([]*eventsEntity.Events, error) {
	return f.events, nil
}

func (f *fakeEventsRepository) GetByID(
	ctx context.Context,
	id common.EventsID,
) (*eventsEntity.Events, error) {
	for _, event := range f.events {
		if event.ID().String() == id.String() {
			return event, nil
		}
	}

	return nil, errors.New("event not found")
}

func (f *fakeEventsRepository) List(
	ctx context.Context,
	id common.DaySessionID,
) ([]*eventsEntity.Events, error) {
	return f.events, nil
}

// ============================================================
// HELPER
// ============================================================

func createTestHandler() (
	*Handler,
	*fakeDaySessionRepository,
) {
	dayRepo := newFakeDaySessionRepository()

	planVersionRepo := newFakePlanVersionRepository()
	planStopRepo := newFakePlanStopRepository()
	eventsRepo := newFakeEventsRepository()

	createService := services.NewDaySessionService(dayRepo)

	getService := services.NewGetDaySessionService(
		dayRepo,
		planVersionRepo,
		planStopRepo,
		eventsRepo,
	)

	listService := services.NewListDaySessionService(dayRepo)

	setActivePlanService := services.NewSetActivePlanService(dayRepo)

	var generateLLMService *services.GenerateLLMPlanService

	handler := NewHandler(
		createService,
		getService,
		listService,
		setActivePlanService,
		generateLLMService,
	)

	return handler, dayRepo
}

// ============================================================
// CREATE DAY SESSION
// ============================================================

func TestHandler_CreateDaySession(t *testing.T) {

	handler, repo := createTestHandler()

	body := `{
		"date": "2026-08-26",
		"start_time": "09:00",
		"start_label": "Hotel"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/day_sessions/trip-1",
		strings.NewReader(body),
	)

	req.SetPathValue("trip_id", "trip-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"expected status %d, got %d, body: %s",
			http.StatusCreated,
			rec.Code,
			rec.Body.String(),
		)
	}

	if len(repo.daySessions) != 1 {
		t.Fatalf(
			"expected 1 day session, got %d",
			len(repo.daySessions),
		)
	}
}

// ============================================================
// CREATE - INVALID JSON
// ============================================================

func TestHandler_CreateDaySession_InvalidJSON(t *testing.T) {

	handler, _ := createTestHandler()

	body := `{
		"date": "2026-08-26",
		"start_time":
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/day_sessions/trip-1",
		strings.NewReader(body),
	)

	req.SetPathValue("trip_id", "trip-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

// ============================================================
// CREATE - INVALID DOMAIN DATA
// ============================================================

func TestHandler_CreateDaySession_InvalidDate(t *testing.T) {

	handler, _ := createTestHandler()

	body := `{
		"date": "26-08-2026",
		"start_time": "09:00",
		"start_label": "Hotel"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/day_sessions/trip-1",
		strings.NewReader(body),
	)

	req.SetPathValue("trip_id", "trip-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code == http.StatusCreated {
		t.Fatal("expected request to fail")
	}
}

// ============================================================
// GET DAY SESSION
// ============================================================

func TestHandler_GetDaySession(t *testing.T) {

	handler, repo := createTestHandler()

	tripID, err := common.NewTripID("trip-1")
	if err != nil {
		t.Fatal(err)
	}

	ds, err := entity.NewDaySession(
		tripID.String(),
		"2026-08-26",
		"09:00",
		"Hotel",
	)
	if err != nil {
		t.Fatal(err)
	}

	dsID, err := common.NewDaySessionID("day-session-1")
	if err != nil {
		t.Fatal(err)
	}

	ds.SetID(dsID)

	repo.daySessions[dsID.String()] = ds

	req := httptest.NewRequest(
		http.MethodGet,
		"/day_sessions/day-session-1",
		nil,
	)

	req.SetPathValue("id", "day-session-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d, body: %s",
			http.StatusOK,
			rec.Code,
			rec.Body.String(),
		)
	}

	var response map[string]interface{}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
}

// ============================================================
// GET DAY SESSION - INVALID ID
// ============================================================

func TestHandler_GetDaySession_InvalidID(t *testing.T) {

	handler, _ := createTestHandler()

	req := httptest.NewRequest(
		http.MethodGet,
		"/day_sessions/",
		nil,
	)

	req.SetPathValue("id", "")

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
// GET DAY SESSION - NOT FOUND
// ============================================================

func TestHandler_GetDaySession_NotFound(t *testing.T) {

	handler, _ := createTestHandler()

	req := httptest.NewRequest(
		http.MethodGet,
		"/day_sessions/does-not-exist",
		nil,
	)

	req.SetPathValue("id", "does-not-exist")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code == http.StatusOK {
		t.Fatal("expected error for non-existing day session")
	}
}

// ============================================================
// LIST DAY SESSIONS
// ============================================================

func TestHandler_ListDaySessions(t *testing.T) {

	handler, repo := createTestHandler()

	tripID, err := common.NewTripID("trip-1")
	if err != nil {
		t.Fatal(err)
	}

	ds1, err := entity.NewDaySession(
		tripID.String(),
		"2026-08-26",
		"09:00",
		"Hotel",
	)
	if err != nil {
		t.Fatal(err)
	}

	ds1ID, err := common.NewDaySessionID("day-session-1")
	if err != nil {
		t.Fatal(err)
	}

	ds1.SetID(ds1ID)

	repo.daySessions[ds1ID.String()] = ds1

	req := httptest.NewRequest(
		http.MethodGet,
		"/day_sessions/trip-1",
		nil,
	)

	req.SetPathValue("trip_id", "trip-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d, body: %s",
			http.StatusOK,
			rec.Code,
			rec.Body.String(),
		)
	}

	var response []interface{}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}

	if len(response) != 1 {
		t.Fatalf(
			"expected 1 day session, got %d",
			len(response),
		)
	}
}

// ============================================================
// LIST - INVALID TRIP ID
// ============================================================

func TestHandler_ListDaySessions_InvalidTripID(t *testing.T) {

	handler, _ := createTestHandler()

	req := httptest.NewRequest(
		http.MethodGet,
		"/day_sessions/",
		nil,
	)

	req.SetPathValue("trip_id", "")

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
// UPDATE ACTIVE PLAN
// ============================================================

func TestHandler_UpdateActivePlan(t *testing.T) {

	handler, repo := createTestHandler()

	tripID, err := common.NewTripID("trip-1")
	if err != nil {
		t.Fatal(err)
	}

	ds, err := entity.NewDaySession(
		tripID.String(),
		"2026-08-26",
		"09:00",
		"Hotel",
	)
	if err != nil {
		t.Fatal(err)
	}

	dsID, err := common.NewDaySessionID("day-session-1")
	if err != nil {
		t.Fatal(err)
	}

	ds.SetID(dsID)

	repo.daySessions[dsID.String()] = ds

	req := httptest.NewRequest(
		http.MethodPut,
		"/day_sessions/day-session-1/active-plan/plan-version-1",
		nil,
	)

	req.SetPathValue("id", "day-session-1")
	req.SetPathValue("planVersionId", "plan-version-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf(
			"expected status %d, got %d, body: %s",
			http.StatusNoContent,
			rec.Code,
			rec.Body.String(),
		)
	}

	if repo.updatedDaySessionID.String() != "day-session-1" {
		t.Fatalf("day session ID was not updated correctly")
	}

	if repo.updatedPlanVersion.String() != "plan-version-1" {
		t.Fatalf("plan version ID was not updated correctly")
	}

	activeID := ds.ActivePlanVersionID()

	if activeID == nil {
		t.Fatal("expected active plan version ID to be set")
	}

	if activeID.String() != "plan-version-1" {
		t.Fatalf(
			"expected active plan version plan-version-1, got %s",
			activeID.String(),
		)
	}
}

// ============================================================
// UPDATE ACTIVE PLAN - MISSING PARAMETERS
// ============================================================

func TestHandler_UpdateActivePlan_MissingParameters(t *testing.T) {

	handler, _ := createTestHandler()

	req := httptest.NewRequest(
		http.MethodPut,
		"/day_sessions//active-plan/",
		nil,
	)

	req.SetPathValue("id", "")
	req.SetPathValue("planVersionId", "")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

// ============================================================
// UPDATE ACTIVE PLAN - INVALID DAY SESSION ID
// ============================================================

func TestHandler_UpdateActivePlan_InvalidDaySessionID(t *testing.T) {

	handler, _ := createTestHandler()

	req := httptest.NewRequest(
		http.MethodPut,
		"/day_sessions/invalid/active-plan/plan-version-1",
		nil,
	)

	req.SetPathValue("id", "")
	req.SetPathValue("planVersionId", "plan-version-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

// ============================================================
// UPDATE ACTIVE PLAN - REPOSITORY ERROR
// ============================================================

func TestHandler_UpdateActivePlan_RepositoryError(t *testing.T) {

	handler, repo := createTestHandler()

	repo.updateErr = errors.New("database error")

	req := httptest.NewRequest(
		http.MethodPut,
		"/day_sessions/day-session-1/active-plan/plan-version-1",
		nil,
	)

	req.SetPathValue("id", "day-session-1")
	req.SetPathValue("planVersionId", "plan-version-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code == http.StatusNoContent {
		t.Fatal("expected repository error")
	}
}

// ============================================================
// UNSUPPORTED METHOD
// ============================================================

func TestHandler_MethodNotAllowed(t *testing.T) {

	handler, _ := createTestHandler()

	req := httptest.NewRequest(
		http.MethodDelete,
		"/day_sessions/day-session-1",
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
// UNKNOWN POST ROUTE
// ============================================================

func TestHandler_PostUnknownRoute(t *testing.T) {

	handler, _ := createTestHandler()

	req := httptest.NewRequest(
		http.MethodPost,
		"/something",
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
// UNKNOWN GET ROUTE
// ============================================================

func TestHandler_GetUnknownRoute(t *testing.T) {

	handler, _ := createTestHandler()

	req := httptest.NewRequest(
		http.MethodGet,
		"/something",
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
// CREATE - REPOSITORY ERROR
// ============================================================

func TestHandler_CreateDaySession_RepositoryError(t *testing.T) {

	handler, repo := createTestHandler()

	repo.createErr = errors.New("database error")

	body := `{
		"date": "2026-08-26",
		"start_time": "09:00",
		"start_label": "Hotel"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/day_sessions/trip-1",
		strings.NewReader(body),
	)

	req.SetPathValue("trip_id", "trip-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code == http.StatusCreated {
		t.Fatal("expected repository error")
	}
}

// ============================================================
// LIST - REPOSITORY ERROR
// ============================================================

func TestHandler_ListDaySessions_RepositoryError(t *testing.T) {

	handler, repo := createTestHandler()

	repo.listErr = errors.New("database error")

	req := httptest.NewRequest(
		http.MethodGet,
		"/day_sessions/trip-1",
		nil,
	)

	req.SetPathValue("trip_id", "trip-1")

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code == http.StatusOK {
		t.Fatal("expected repository error")
	}
}
