package dto

import (
	"assistant_suggestions/domain/entity"
	"time"
)


// AssistantSuggestionDTO represents a assistant suggestion returned by the API.
type AssistantSuggestionsDTO struct {
	ID           string
	DaysessionID string
	Message      string
	Status       string
	CreatedAt    string
}

func ToAssistantSuggestionsDTO(a *entity.AssistantSuggestions) AssistantSuggestionsDTO {
	return AssistantSuggestionsDTO{
		ID:           a.ID().String(),
		DaysessionID: a.DaysessionID().String(),
		Message:      a.Message(),
		Status:       a.Status(),
		CreatedAt:    a.CreatedAt().Format(time.RFC3339),
	}
}
