package rest

import (
	"common"
	"encoding/json"
	"fmt"
	"net/http"
	"plans/application/command"
	"plans/application/services"
	"time"
)

type Handlers struct {
	createPlanStop *services.CreatePlanStopService

	listPlanStop *services.ListPlanStopService
}

func NewHandlers(createps *services.CreatePlanStopService, listps *services.ListPlanStopService) *Handlers {
	return &Handlers{
		createPlanStop: createps,
		listPlanStop:   listps,
	}
}

func (h *Handlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== PlanStop ServeHTTP ===", r.Method, r.URL.Path)
	switch r.Method {
	case http.MethodPost:
		h.create(w, r)

	case http.MethodGet:
		h.listPlanstop(w, r)

	default:
		writeError(w, r, http.StatusMethodNotAllowed, "bad_request", "Method not allowed")
	}
}

type CreatePlanStopRequest struct {
	// Starting Place number
	Position         int    `json:"position"`
	
	// Place name
	Title            string `json:"title"`
	
	// Category of the Place
	CategoryLabel    string `json:"categorylabel"`
	
	// Image URL of the Place
	ImageURL         string `json:"imageurl"`
	
	// Arrival Time
	PlannedArrival   string `json:"plannedarrival"`
	
	// Departure Time
	PlannedDeparture string `json:"planneddeparture"`
	
	// Travel Time
	TravelMinutes    int    `json:"travelminutes"`
	
	// Spendable time in the Place
	StayMinutes      int    `json:"stayminutes"`
	
	// Busy Risk Label
	BusyRiskLabel    string `json:"busyrisklabel"`
}

func (h *Handlers) create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">>> ENTERED PLAN STOP CREATE")
	fmt.Println("===== PLAN STOP CREATE HANDLER =====")
	var req CreatePlanStopRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, http.StatusNotFound, "bad_request", "invalid json body")
		return
	}

	planversionID, err := common.NewPlanVersionID(r.PathValue("id"))
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

	Planneddeparture, err := time.Parse(time.RFC3339, req.PlannedDeparture)
	if err != nil {
		writeError(w, r, http.StatusBadRequest,
			"bad_request",
			"invalid planned_arrival")
		return
	}

	fmt.Println("Before CreateStop")

	planstop, err := h.createPlanStop.CreateStop(r.Context(), planversionID, command.CreatePlanStopCommand{
		PlanVersionID:    planversionID,
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
	fmt.Println("After CreateStop")
	if err != nil {
		fmt.Printf("ERROR TYPE: %T\n", err)
		fmt.Printf("ERROR: %+v\n", err)
		writeDomainError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(planstop)
}

func (h *Handlers) listPlanstop(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

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
		fmt.Printf("ERROR TYPE: %T\n", err)
		fmt.Printf("ERROR: %+v\n", err)
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(planstop); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
