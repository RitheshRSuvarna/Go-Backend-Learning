package services

import (
	"assistant_suggestions/application/command"
	"assistant_suggestions/application/dto"
	"assistant_suggestions/domain/entity"
	"assistant_suggestions/domain/repository"
	"context"
)

type CreateAssistantSuggestionsService struct {
	sugrepo repository.AssistantSuggestionRepository
}

func NewAssistantSuggestionService(sugrepo repository.AssistantSuggestionRepository) *CreateAssistantSuggestionsService {
	return &CreateAssistantSuggestionsService{sugrepo: sugrepo}
}

func (c *CreateAssistantSuggestionsService) CreateAssSug(ctx context.Context, cmd command.CreateAssistantSuggestionCommand) (dto.AssistantSuggestionsDTO, error) {
	AssSug, err := entity.NewAssistantSuggestions(cmd.DaySessionID, cmd.Message, cmd.Status)
	if err != nil {
		return dto.AssistantSuggestionsDTO{}, nil
	}
	if err := c.sugrepo.Create(ctx, AssSug); err != nil {
		return dto.AssistantSuggestionsDTO{}, err
	}
	return dto.ToAssistantSuggestionsDTO(AssSug), nil
}
