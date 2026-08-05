package rest

import (
	"common"
	"encoding/json"
	"fmt"
	"net/http"
	"plans/application/command"
	"plans/application/services"
	_ "plans/application/dto"
	"strings"
)

type Handler struct {
	createPlanVersion *services.CreatePlanVersionService
	listplanversion   *services.ListPlanVerionservice
	getPlanversion    *services.GetByIDPlanVersionService
}

type ActivePlanHandler struct {
	getPlanversion *services.GetByIDPlanVersionService
}

func NewHandler(createpv *services.CreatePlanVersionService, getpln *services.GetByIDPlanVersionService, listpv *services.ListPlanVerionservice) *Handler {
	return &Handler{
		createPlanVersion: createpv,
		getPlanversion:    getpln,
		listplanversion:   listpv,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== PlanVersion ServeHTTP ===", r.Method, r.URL.Path)

	switch {

	// GET/POST /day-sessions/{id}/plan-versions
	case strings.HasSuffix(r.URL.Path, "/plan-versions"):

		switch r.Method {

		case http.MethodPost:
			h.create(w, r)

		case http.MethodGet:
			h.list(w, r)

		default:
			writeError(
				w,
				r,
				http.StatusMethodNotAllowed,
				"bad_request",
				"method not allowed",
			)
		}

	// GET /day-sessions/{id}/active-plan
	case strings.HasSuffix(r.URL.Path, "/active-plan"):

		switch r.Method {

		case http.MethodGet:
			h.getactiveplan(w, r)

		default:
			writeError(
				w,
				r,
				http.StatusMethodNotAllowed,
				"bad_request",
				"method not allowed",
			)
		}

	default:
		writeError(
			w,
			r,
			http.StatusNotFound,
			"not_found",
			"not found",
		)
	}
}

type CreatePlanVersionRequest struct {

	// Version number
	Version int `json:"version"`

	// Note
	Note string `json:"note"`
}

// CreatePlanVersion godoc
// @Summary Create Plan Version
// @Description Create planversions for a daysessions
// @Tags Plan
// @Accept json
// @Produce json
// @Param id header string true "daysession ID"
// @Param id header string true "daysession ID"
// @Success 201 {object} dto.PlanVersionDTO
// @Failure 400 {object} apiError
// @Failure 404 {object} apiError
// @Failure 500 {object} apiError
// @Router /api/day_sessions/{id}/plan-versions [post]
func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PlanVersion create handler called")
	var req CreatePlanVersionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, http.StatusBadRequest, "bad_request", "invalid json body")
		return
	}

	daySessionID := r.PathValue("id")

	daysessionID, err := common.NewDaySessionID(daySessionID)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	planversion, err := h.createPlanVersion.CreatePlanVersion(
		r.Context(),
		daysessionID,
		command.CreatePlanVersionCommand{
			DaysessionID: daySessionID,
			Version:      req.Version,
			Note:         req.Note,
		},
	)
	if err != nil {
		fmt.Printf("ERROR TYPE: %T\n", err)
		fmt.Printf("ERROR: %+v\n", err)
		writeDomainError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(planversion)
}

// CreatePlanVersion godoc
// @Summary Create Plan Version
// @Description Create planversions for a daysessions
// @Tags Plan
// @Accept json
// @Produce json
// @Param id header string true "daysession ID"
// @Param request body CreatePlanVersionRequest true "PlanVersion details"
// @Success 201 {object} dto.PlanVersionDTO
// @Failure 400 {object} apiError
// @Failure 404 {object} apiError
// @Failure 500 {object} apiError
// @Router /api/day_sessions/{id}/plan-versions [get]
func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	daysessionid, err := common.NewDaySessionID(id)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	daysession, err := h.listplanversion.ListVersion(
		r.Context(),
		daysessionid,
	)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("content-type", "application/json")

	if err := json.NewEncoder(w).Encode(daysession); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CreatePlanVersion godoc
// @Summary Create Plan Version
// @Description Create planversions for a daysessions
// @Tags Plan
// @Accept json
// @Produce json
// @Param id header string true "daysession ID"
// @Success 201 {object} dto.PlanVersionDTO
// @Failure 400 {object} apiError
// @Failure 404 {object} apiError
// @Failure 500 {object} apiError
// @Router /api/day_sessions/{id}/plan-versions [get]
func (h *Handler) getactiveplan(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	activeplan, err := common.NewDaySessionID(id)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	planversion, err := h.getPlanversion.GetActivePlan(
		r.Context(),
		activeplan,
	)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(planversion); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
