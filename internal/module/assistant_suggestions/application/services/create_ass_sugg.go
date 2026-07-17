package services

import (
	"assistant_suggestions/application/command"
	"assistant_suggestions/application/dto"
	"assistant_suggestions/domain/entity"
	"assistant_suggestions/domain/repository"
	"context"
	"fmt"
)

type CreateAssistantSuggestionsService struct {
	sugrepo repository.AssistantSuggestionRepository
}

func NewAssistantSuggestionService(sugrepo repository.AssistantSuggestionRepository) *CreateAssistantSuggestionsService {
	return &CreateAssistantSuggestionsService{sugrepo: sugrepo}
}

func (c *CreateAssistantSuggestionsService) CreateAssistantSuggestions(ctx context.Context, cmd command.CreateAssistantSuggestionCommand) (dto.AssistantSuggestionsDTO, error) {
	AssSug, err := entity.NewAssistantSuggestions(cmd.DaySessionID, cmd.Message, cmd.Status)
	if err != nil {
		return dto.AssistantSuggestionsDTO{}, err
	}
	// fmt.Printf("Entity before DB: %+v\n", AssSug)
	if err := c.sugrepo.Create(ctx, AssSug); err != nil {
		fmt.Printf("Entity after DB: %v\n", err)
		return dto.AssistantSuggestionsDTO{}, err
	}
	// fmt.Printf("Entity after DB: %+v\n", AssSug)
	return dto.ToAssistantSuggestionsDTO(AssSug), nil
}
