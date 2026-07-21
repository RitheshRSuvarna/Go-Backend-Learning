package rest

import (
	"common"
	"encoding/json"
	"events/application/command"
	"events/application/services"
	"net/http"
)

type Handler struct {
	createvent *services.CreateEventService
	getevent   *services.GetEventService
}

func NewHandler(createvent *services.CreateEventService, getevent *services.GetEventService) *Handler {
	return &Handler{
		createvent: createvent,
		getevent:   getevent,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		if r.PathValue("id") == "" {
			writeError(w, r, http.StatusNotFound, "not_found", "not found")
			return
		}
		h.create(w, r)

	case http.MethodGet:
		if r.PathValue("id") == "" {
			writeError(w, r, http.StatusNotFound, "not_found", "not found")
			return
		}
		h.get(w, r)

	default:
		writeError(w, r, http.StatusMethodNotAllowed, "method_not_allowed", "method not allowed")
	}
}

type CreateEventRequest struct {
	EventType string          `json:"eventType"`
	Payload   json.RawMessage `json:"payload"`
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateEventRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, http.StatusBadRequest, "Bad_Request", "bad request")
		return
	}

	daysessionid := r.PathValue("id")

	event, err := h.createvent.CreateEvents(r.Context(), command.CreateEventsCommand{
		DaySessionID: daysessionid,
		EventType:    req.EventType,
		Payload:      req.Payload,
	})
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(event)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	daysessionid := r.PathValue("id")

	daysid, err := common.NewDaySessionID(daysessionid)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	event, err := h.getevent.GetEvents(r.Context(), daysid)
	if err != nil {
		writeDomainError(w, r, err)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(event); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
