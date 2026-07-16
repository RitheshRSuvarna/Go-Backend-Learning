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

func (s *GetAssistantSuggestionService) GetAssistantSuggestions(ctx context.Context, id common.DaySessionID) ([]dto.AssistantSuggestionsDTO, error) {
	assis, err := s.sugrepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	out := make([]dto.AssistantSuggestionsDTO, 0, len(assis))

	for _, as := range assis {
		out = append(out, dto.ToAssistantSuggestionsDTO(as))
	}
	return out, nil
}