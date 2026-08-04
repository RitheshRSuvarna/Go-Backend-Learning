package rest

import (
	"common"
	"day_session/application/command"
	"day_session/application/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	createDaysession *services.CreateDaySessionService
	getDaysession    *services.GetDaySessionService
	listDaysession   *services.ListDaySessionService
	setActivePlan *services.SetActivePlanService
}

func NewHandler(createds *services.CreateDaySessionService, getds *services.GetDaySessionService, listds *services.ListDaySessionService, updtactpln *services.SetActivePlanService) *Handler {
	return &Handler{
		createDaysession: createds,
		getDaysession:    getds,
		listDaysession:   listds,
		setActivePlan: updtactpln,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== DAY SESSION ServeHTTP ===", r.Method, r.URL.Path)

	switch r.Method {

	case http.MethodPost:
		if tripID := r.PathValue("trip_id"); tripID != "" {
			h.create(w, r)
			return
		}

		writeError(w, r, http.StatusNotFound, "not_found", "not found")

	case http.MethodGet:

		// GET /day-sessions/{id}
		if id := r.PathValue("trip_id"); id != "" {
			h.list(w, r)
			return
		}

		// GET /day-sessions?trip_id=...&date=...
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
	// Trip id
	TripID     string `json:"trip_id"`
	
	// Date
	Date       string `json:"date"`
	
	// Day session starting time
	StartTime  string `json:"start_time"`
	
	// Starting Place
	StartLabel string `json:"start_label"`
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("===== DAY SESSION CREATE HANDLER =====")
	tripid := r.PathValue("trip_id")
	var req CreateDaySessionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Decode error:", err)
		writeError(w, r, http.StatusBadRequest, "bad_request", "invalid json body")
		return
	}
	fmt.Printf("%+v\n", req)

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

func (h *Handler) getdaysession(w http.ResponseWriter, r *http.Request) {
	fmt.Println("===== ENTERED getdaysession HANDLER =====")
	daysessionID := r.PathValue("id")

	domaindaysessionID, err := common.NewDaySessionID(daysessionID)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}
	fmt.Println("1. DONE")

	daySession, err := h.getDaysession.GetDaySession(
		r.Context(),
		domaindaysessionID,
	)
	fmt.Printf("result = %+v\n", daySession)
	fmt.Printf("err = %v\n", err)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(daySession); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("trip_id")

	tripID, err := common.NewTripID(id)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	daySession, err := h.listDaysession.ListDaysession(
		r.Context(),
		tripID,
	)
	if err != nil {
		fmt.Println("List Handler Error:", err)
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(daySession); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func (h *Handler) updateActivePlan(
    w http.ResponseWriter,
    r *http.Request,
    daySessionID string,
    planVersionID string,
) {
    dsID, err := common.NewDaySessionID(daySessionID)
    if err != nil {
        // handle error
        return
    }

    pvID, err := common.NewPlanVersionID(planVersionID)
    if err != nil {
        // handle error
        return
    }

    err = h.setActivePlan.UpdateActivePlan(r.Context(), dsID, pvID)
    if err != nil {
        // handle error
        return
    }

    w.WriteHeader(http.StatusNoContent)
}