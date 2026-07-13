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
	listDaysession   *services.ListDaySessionService
	listDaysessionID *services.ListDaySessionServiceID
}

func NewHandler(createds *services.CreateDaySessionService, listds *services.ListDaySessionService, listdsid *services.ListDaySessionServiceID) *Handler {
	return &Handler{
		createDaysession: createds,
		listDaysession:   listds,
		listDaysessionID: listdsid,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/day-sessions" {
		writeError(w, r, http.StatusNotFound, "not_found", "not found")
		return
	}

	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	case http.MethodGet:
		switch {
		case r.URL.Query().Get("trip_id") != "":
			h.list(w, r)

		case r.URL.Query().Get("day-session_id") != "" &&
			r.URL.Query().Get("date") != "":
			h.getByTripIDAndDate(w, r)
		}
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

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateDaySessionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Decode error:", err)
		writeError(w, r, http.StatusBadRequest, "bad_request", "invalid json body")
		return
	}
	fmt.Printf("%+v\n", req)

	daysession, err := h.createDaysession.CreateDaySession(r.Context(), command.CreateDaySessionCommand{
		TripID:     req.TripID,
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

func (h *Handler) getByTripIDAndDate(
	w http.ResponseWriter,
	r *http.Request,
) {
	tripID := r.URL.Query().Get("trip_id")
	date := r.URL.Query().Get("date")

	domainTripID, err := common.NewTripID(tripID)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	daySession, err := h.listDaysession.GetByTripIDAndDate(
		r.Context(),
		domainTripID,
		date,
	)
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
	id := r.PathValue("id")

	daysessionID, err := common.NewDaySessionID(id)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	daySession, err := h.listDaysessionID.GetByID(
		r.Context(),
		daysessionID,
	)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(daySession); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
