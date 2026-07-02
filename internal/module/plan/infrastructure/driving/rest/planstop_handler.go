package rest

import (
	"common"
	"encoding/json"
	"net/http"
	"plan/application/command"
	"plan/application/services"
	"time"
)

type Handlers struct {
	createPlanStop *services.CreatePlanStopService
	getPlanStop    *services.GetStopByIDService
	listPlanStop   *services.ListPlanStopService
}

func NewHandlers(createps *services.CreatePlanStopService, getps *services.GetStopByIDService, listps *services.ListPlanStopService) *Handlers {
	return &Handlers{
		createPlanStop: createps,
		getPlanStop:    getps,
		listPlanStop:   listps,
	}
}

func (h *Handlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/plan-stop" {
		writeError(w, r, http.StatusNotFound, "not_found", "not found")
		return
	}

	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	case http.MethodGet:
		if r.PathValue("id") != "" {
			h.getbyid(w, r)
		} else if r.PathValue("planverisonid") != "" {
			h.listPlanstop(w, r)
		} else {
			writeError(w, r, http.StatusBadRequest, "bad_request", "Missing identifier")
		}
	default:
		writeError(w, r, http.StatusMethodNotAllowed, "bad_request", "Method not allowed")
	}
}

type CreatePlanStopRequest struct {
	PlanVersionID    string `json:"planversionid"`
	Position         int    `json:"position"`
	Title            string `json:"title"`
	CategoryLabel    string `json:"categorylabel"`
	ImageURL         string `json:"imageurl"`
	PlannedArrival   string `json:"plannedarrival"`
	PlannedDeparture string `json:"planneddeparture"`
	TravelMinutes    int    `json:"travelminutes"`
	StayMinutes      int    `json:"stayminutes"`
	BusyRiskLabel    string `json:"busyrisklabel"`
}

func (h *Handlers) create(w http.ResponseWriter, r *http.Request) {
	var req CreatePlanStopRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, http.StatusNotFound, "bad_request", "invalid json body")
		return
	}

	planversionID, err := common.NewPlanVersionID(req.PlanVersionID)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}
	Plannedarrival, err := time.Parse(time.RFC3339, req.PlannedArrival)
	if err != nil {
		writeError(w, r, http.StatusBadRequest,
			"bad_request",
			"invalid planned_arrival")
		return
	}

	Planneddeparture, err := time.Parse(time.RFC3339, req.PlannedArrival)
	if err != nil {
		writeError(w, r, http.StatusBadRequest,
			"bad_request",
			"invalid planned_arrival")
		return
	}

	planstop, err := h.createPlanStop.CreateStop(r.Context(), planversionID, command.CreatePlanStopCommand{
		Position:         req.Position,
		Title:            req.Title,
		CategoryLabel:    req.CategoryLabel,
		ImageURL:         req.ImageURL,
		PlannedArrival:   Plannedarrival,
		PlannedDeparture: Planneddeparture,
		TravelMinutes:    req.TravelMinutes,
		StayMinutes:      req.StayMinutes,
		BusyRiskLabel:    req.BusyRiskLabel,
	})
	if err != nil {
		writeDomainError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(planstop)
}

func (h *Handlers) getbyid(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	planstopID, err := common.NewPlanStopID(id)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	planstop, err := h.getPlanStop.GetByID(
		r.Context(),
		planstopID,
	)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	if err := json.NewEncoder(w).Encode(planstop); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) listPlanstop(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("planversionid")

	planversionID, err := common.NewPlanVersionID(id)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	planstop, err := h.listPlanStop.ListPlanStop(
		r.Context(),
		planversionID,
	)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(planstop); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
