package services

import (
	"assistant_suggestions/application/dto"
	"assistant_suggestions/domain/repository"
	"common"
	"context"
)

type EditAssistantSuggestionService struct {
	sugrepo repository.AssistantSuggestionRepository
}

func NewEditAssistantSuggestionService(sugrepo repository.AssistantSuggestionRepository) *EditAssistantSuggestionService {
	return &EditAssistantSuggestionService{sugrepo: sugrepo}
}

func (e *EditAssistantSuggestionService) EditAssistantSuggestions(ctx context.Context, id common.AssistantSuggestionsID, message, status string) (dto.AssistantSuggestionsDTO, error) {
	Sugg, err := e.sugrepo.Edit(ctx, id, message, status)
	if err != nil {
		return dto.AssistantSuggestionsDTO{}, err
	}
	return dto.ToAssistantSuggestionsDTO(Sugg), nil
}
