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

func (e *EditAssistantSuggestionService) EditAssistantSugg(ctx context.Context, id common.AssistantSuggestionsID) (dto.AssistantSuggestionsDTO, error) {
	Sugg, err := e.sugrepo.Edit(ctx, id)
	if err != nil {
		return dto.AssistantSuggestionsDTO{}, err
	}
	return dto.ToAssistantSuggestionsDTO(Sugg), nil
}
