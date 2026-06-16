package rest

import (
	"day_session/application/command"
	"day_session/application/services"
	"encoding/json"
	"net/http"
)

type Handler struct {
	createDaysession *services.CreateDaySessionService
	listDaysession   *services.ListDaySessionService
}

func NewHandler(createds *services.CreateDaySessionService, listds *services.ListDaySessionService) *Handler {
	return &Handler{
		createDaysession: createds,
		listDaysession:   listds,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/day_session" {
		writeError(w, r, http.StatusNotFound, "not_found", "not found")
		return
	}

	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	case http.MethodGet:
		h.list(w, r)
	default:
		writeError(w, r, http.StatusMethodNotAllowed, "bad_request", "method not allowed")
	}
}

type CreateDaySessionRequest struct {
	TripID     string `json:"tripid"`
	Date       string `json:"date"`
	StartTime  string `json:"start time"`
	StartLabel string `json:"start label"`
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateDaySessionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, http.StatusBadRequest, "bad_request", "invalid json body")
		return
	}

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

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	daysession, err := h.listDaysession.GetByID(r.Context(), r.)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(daysession)
}
