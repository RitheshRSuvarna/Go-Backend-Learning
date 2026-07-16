package repository

import (
	"context"
	"common"
	"assistant_suggestions/domain/entity"
)
type AssistantSuggestionRepository interface {
	Create(ctx context.Context, assistantsuggestion *entity.AssistantSuggestions) error
	Get(ctx context.Context, id common.DaySessionID) ([]*entity.AssistantSuggestions, error)
	Edit(ctx context.Context, id common.AssistantSuggestionsID, message, status string) (*entity.AssistantSuggestions, error)
}
