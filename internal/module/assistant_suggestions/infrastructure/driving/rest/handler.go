package rest

import (
	"assistant_suggestions/application/command"
	"assistant_suggestions/application/services"
	"common"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	createsug *services.CreateAssistantSuggestionsService
	getsugg   *services.GetAssistantSuggestionService
	editsugg  *services.EditAssistantSuggestionService
}

func NewHandler(createsug *services.CreateAssistantSuggestionsService, getsugg *services.GetAssistantSuggestionService, editsugg *services.EditAssistantSuggestionService) *Handler {
	return &Handler{
		createsug: createsug,
		getsugg:   getsugg,
		editsugg:  editsugg,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		if r.URL.Path != "/assistant_suggestion" {
			writeError(w, r, http.StatusNotFound, "not_found", "not found")
			return
		}
		h.create(w, r)

	case http.MethodGet:

		// GET /day-sessions/{id}
		if id := r.PathValue("day_session_id"); id != "" {
			h.get(w, r)
			return
		}

	case http.MethodPut:
		if id := r.PathValue("assistant_suggestion_id"); id != "" {
			h.edit(w, r)
			return
		}
	}
}

type CreateAssistantSuggestionRequest struct {
	DaySessionID string `json: day_session_id`
	Message      string `json: message`
	Status       string `json: status`
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateAssistantSuggestionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("Decode error", err)
		writeError(w, r, http.StatusBadRequest, "Bad_request", "invalid json")
		return
	}
	fmt.Printf("%v\n", req)

	assistantsuggestion, err := h.createsug.CreateAssistantSuggestions(r.Context(), command.CreateAssistantSuggestionCommand{
		DaySessionID: req.DaySessionID,
		Message:      req.Message,
		Status:       req.Status,
	})
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(assistantsuggestion)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	daysessionid := r.URL.Query().Get("day_session_id")

	dayid, err := common.NewDaySessionID(daysessionid)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	assisSugg, err := h.getsugg.GetAssistantSuggestions(r.Context(), dayid)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(assisSugg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type EditAssistantSuggestionRequest struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (h *Handler) edit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("assistant_suggestion_id")

	sugid, err := common.NewAssistantSuggestionsID(id)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	var req EditAssistantSuggestionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	assistantsugg, err := h.editsugg.EditAssistantSuggestions(r.Context(), sugid, req.Message, req.Status)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(assistantsugg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
