package rest

import (
	"common"
	"encoding/json"
	"net/http"
	"plan/application/command"
	"plan/application/services"
)

type Handler struct {
	createPlanVersion *services.CreatePlanVersionService
	getPlanversion    *services.GetByIDPlanVersionService
	listplanversion   *services.ListPlanVerionservice
}

func NewHandler(createpv *services.CreatePlanVersionService, getpv *services.GetByIDPlanVersionService, listpv *services.ListPlanVerionservice) *Handler {
	return &Handler{
		createPlanVersion: createpv,
		getPlanversion:    getpv,
		listplanversion:   listpv,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/planversions" {
		writeError(w, r, http.StatusNotFound, "not_found", "not found")
		return
	}

	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	case http.MethodGet:
    if r.PathValue("id") != "" {
        h.getbyid(w, r)
    } else if r.PathValue("daySessionID") != "" {
        h.list(w, r)
    } else {
        writeError(w, r, http.StatusBadRequest,
            "bad_request",
            "missing identifier",
        )
    }

	default:
		writeError(w, r, http.StatusMethodNotAllowed, "bad_request", "method not allowed")
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

func (h *Handler) getbyid(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	planversionID, err := common.NewPlanVersionID(id)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	planversion, err := h.getPlanversion.GetPlanVersionByID(
		r.Context(),
		planversionID,
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