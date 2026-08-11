package rest

import (
	"common"
	"day_session/application/command"
	_ "day_session/application/dto"
	"day_session/application/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Handler struct {
	createDaysession       *services.CreateDaySessionService
	getDaysession          *services.GetDaySessionService
	listDaysession         *services.ListDaySessionService
	setActivePlan          *services.SetActivePlanService
	generateLLMPlanService *services.GenerateLLMPlanService
}

func NewHandler(createds *services.CreateDaySessionService, getds *services.GetDaySessionService, listds *services.ListDaySessionService, updtactpln *services.SetActivePlanService, generateLLMPlanSvc *services.GenerateLLMPlanService) *Handler {
	return &Handler{
		createDaysession:       createds,
		getDaysession:          getds,
		listDaysession:         listds,
		setActivePlan:          updtactpln,
		generateLLMPlanService: generateLLMPlanSvc,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== DAY SESSION ServeHTTP ===", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodPost:
		if r.PathValue("id") != "" && strings.HasSuffix(r.URL.Path, "/llm/plan") {
			h.generateLLMPlan(w, r)
			return
		}

		if tripID := r.PathValue("trip_id"); tripID != "" {
			h.create(w, r)
			return
		}

		writeError(w, r, http.StatusNotFound, "not_found", "not found")

	case http.MethodGet:
		if id := r.PathValue("trip_id"); id != "" {
			h.list(w, r)
			return
		}

		if id := r.PathValue("id"); id != "" {
			h.getdaysession(w, r)
			return
		}

		writeError(w, r, http.StatusNotFound, "not_found", "not found")

	case http.MethodPut:
		daySessionID := r.PathValue("id")
		planVersionID := r.PathValue("planVersionId")

		if daySessionID == "" || planVersionID == "" {
			writeError(w, r, http.StatusBadRequest, "bad_request", "missing path parameter")
			return
		}
		h.updateActivePlan(w, r, daySessionID, planVersionID)

	default:
		writeError(w, r, http.StatusMethodNotAllowed, "bad_request", "method not allowed")
	}
}

type CreateDaySessionRequest struct {
	TripID     string `json:"trip_id"`
	Date       string `json:"date"`
	StartTime  string `json:"start_time"`
	StartLabel string `json:"start_label"`
}

// CreateDaySessions godoc
// @Summary Create a Daysessions
// @Description Creates daysessions for the specified Trip.
// @Tags DaySession
// @Accept json
// @Produce json
// @Param x-trip_id header string true "Trip ID"
// @Param request body CreateDaySessionRequest true "DaySession details"
// @Success 201 {object} dto.DaySessionDTO
// @Failure 400 {object} apiError
// @Failure 404 {object} apiError
// @Failure 500 {object} apiError
// @Router /api/day_sessions/{trip_id} [post]
func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	tripid := r.PathValue("trip_id")
	var req CreateDaySessionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, http.StatusBadRequest, "bad_request", "invalid json body")
		return
	}

	daysession, err := h.createDaysession.CreateDaySession(r.Context(), command.CreateDaySessionCommand{
		TripID:     tripid,
		Date:       req.Date,
		StartTime:  req.StartTime,
		StartLabel: req.StartLabel,
	})
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(daysession)
}

// GenerateLLMPlan godoc
// @Summary Generate a day-session plan using the LLM service
// @Description Calls the deterministic LLM contract service to generate a plan for the day session.
// @Tags DaySession
// @Accept json
// @Produce json
// @Param id path string true "Day session ID"
// @Success 200 {object} map[string]any
// @Failure 400 {object} apiError
// @Failure 502 {object} apiError
// @Failure 504 {object} apiError
// @Router /api/day-sessions/{id}/llm/plan [post]
func (h *Handler) generateLLMPlan(w http.ResponseWriter, r *http.Request) {
	daySessionID := r.PathValue("id")
	if _, err := common.NewDaySessionID(daySessionID); err != nil {
		writeDomainError(w, r, err)
		return
	}

	plan, err := h.generateLLMPlanService.GeneratePlan(r.Context(), daySessionID)
	if err != nil {
		writeError(w, r, http.StatusBadGateway, "llm_error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(plan)
}

// GetDaySessions godoc
// @Summary Gets a specfic Daysession by daysession id
// @Description Gets a daysession from the specified Trip.
// @Tags DaySession
// @Accept json
// @Produce json
// @Param id header string true "daysession ID"
// @Success 201 {object} dto.DaySessionDTO
// @Failure 400 {object} apiError
// @Failure 404 {object} apiError
// @Failure 500 {object} apiError
// @Router /api/day_sessions/{id} [get]
func (h *Handler) getdaysession(w http.ResponseWriter, r *http.Request) {
	daysessionID := r.PathValue("id")

	domaindaysessionID, err := common.NewDaySessionID(daysessionID)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	daySession, err := h.getDaysession.GetDaySession(r.Context(), domaindaysessionID)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(daySession); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ListDaySessions godoc
// @Summary List all the Daysessions
// @Description List daysessions for the specified Trip.
// @Tags DaySession
// @Accept json
// @Produce json
// @Param x-trip_id header string true "Trip ID"
// @Param request body CreateDaySessionRequest true "DaySession details"
// @Success 201 {object} dto.DaySessionDTO
// @Failure 400 {object} apiError
// @Failure 404 {object} apiError
// @Failure 500 {object} apiError
// @Router /api/day_sessions/{trip_id} [get]
func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("trip_id")

	tripID, err := common.NewTripID(id)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	daySession, err := h.listDaysession.ListDaysession(r.Context(), tripID)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(daySession); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdatesActivePlan godoc
// @Summary Updates active plan
// @Description updates the active plan of the day session.
// @Tags DaySession
// @Accept json
// @Produce json
// @Param x-d header string true "Trip ID"
// @Param x-vid header string true "Planversion ID"
// @Success 201 {object} dto.DaySessionDTO
// @Failure 400 {object} apiError
// @Failure 404 {object} apiError
// @Failure 500 {object} apiError
// @Router /api/day_sessions/{id}/active-plan/{vid} [put]
func (h *Handler) updateActivePlan(w http.ResponseWriter, r *http.Request, daySessionID string, planVersionID string) {
	dsID, err := common.NewDaySessionID(daySessionID)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	pvID, err := common.NewPlanVersionID(planVersionID)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	if err := h.setActivePlan.UpdateActivePlan(r.Context(), dsID, pvID); err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
