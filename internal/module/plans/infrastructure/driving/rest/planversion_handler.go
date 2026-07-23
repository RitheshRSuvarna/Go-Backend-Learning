package rest

import (
	"common"
	"encoding/json"
	"net/http"
	"plans/application/command"
	"plans/application/services"
	"strings"
)

type Handler struct {
	createPlanVersion *services.CreatePlanVersionService
	listplanversion   *services.ListPlanVerionservice
	getPlanversion    *services.GetByIDPlanVersionService
}

type ActivePlanHandler struct {
	getPlanversion    *services.GetByIDPlanVersionService
}

func NewHandler(createpv *services.CreatePlanVersionService, getpln *services.GetByIDPlanVersionService, listpv *services.ListPlanVerionservice) *Handler {
	return &Handler{
		createPlanVersion: createpv,
		getPlanversion: getpln,
		listplanversion:   listpv,
	}
}

// func NewActivePlanHandler(getpv *services.GetByIDPlanVersionService) *ActivePlanHandler {
// 	return &ActivePlanHandler{
		
// 	}
// }

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
	DaySessionID string `json:"daysessionid"`
	Version      int    `json:"version"`
	Note         string `json:"note"`
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req CreatePlanVersionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, http.StatusBadRequest, "bad_request", "invalid json body")
		return
	}

	daysessionID, err := common.NewDaySessionID(req.DaySessionID)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	planversion, err := h.createPlanVersion.CreatePlanVersion(r.Context(), daysessionID, command.CreatePlanVersionCommand{
		DaysessionID: req.DaySessionID,
		Version:      req.Version,
		Note:         req.Note,
	})
	if err != nil {
		writeDomainError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(planversion)
}


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