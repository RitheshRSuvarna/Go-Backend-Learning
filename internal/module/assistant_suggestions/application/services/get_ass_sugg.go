package services

import (
	"context"
	"common"
	"assistant_suggestions/domain/repository"
	"assistant_suggestions/application/dto"
)

type GetAssistantSuggestionService struct {
	sugrepo repository.AssistantSuggestionRepository
}

func NewGetAssistantSuggestionService( sugrepo repository.AssistantSuggestionRepository) *GetAssistantSuggestionService{
	return &GetAssistantSuggestionService{sugrepo: sugrepo}
}

func (s *GetAssistantSuggestionService) GetAssistantSuggestions(ctx context.Context, id common.DaySessionID) (dto.AssistantSuggestionsDTO, error) {
	sugg, err := s.sugrepo.Get(ctx, id)
	if err != nil {
		return dto.AssistantSuggestionsDTO{}, err
	}
	return dto.ToAssistantSuggestionsDTO(sugg), nil
}