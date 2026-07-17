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
	// fmt.Println("===== AssistantSuggestionHandler =====")
	// fmt.Println("Method:", r.Method)
	// fmt.Println("Path:", r.URL.Path)
	// fmt.Println("ID:", r.PathValue("id"))

	switch r.Method {

	case http.MethodPost:
		if r.PathValue("id") == "" {
			fmt.Println("404 returned from GET validation")
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

	case http.MethodPut:
		if r.PathValue("id") == "" {
			writeError(w, r, http.StatusNotFound, "not_found", "not found")
			return
		}
		h.edit(w, r)

	default:
		writeError(
			w,
			r,
			http.StatusMethodNotAllowed,
			"method_not_allowed",
			"method not allowed",
		)
	}
}

type CreateAssistantSuggestionRequest struct {
	DaySessionID string `json:"day_session_id"`
	Message      string `json:"message"`
	Status       string `json:"status"`
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateAssistantSuggestionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, http.StatusBadRequest, "Bad_request", "invalid json")
		return
	}
	// fmt.Printf("%+v\n", req)

	daySessionID := r.PathValue("id")

	assistantsuggestion, err := h.createsug.CreateAssistantSuggestions(r.Context(), command.CreateAssistantSuggestionCommand{
		DaySessionID: daySessionID,
		Message:      req.Message,
		Status:       req.Status,
	})
	// fmt.Printf("DTO returned: %+v\n", assistantsuggestion)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(assistantsuggestion)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	daysessionid := r.PathValue("id")

	dayid, err := common.NewDaySessionID(daysessionid)
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	assisSugg, err := h.getsugg.GetAssistantSuggestions(r.Context(), dayid)
	if err != nil {
		fmt.Printf("GET Handler Error: %v\n", err)
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
	id := r.PathValue("id")

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