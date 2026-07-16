package repository

import (
	"assistant_suggestions/domain/entity"
	queries "assistant_suggestions/infrastructure/driven/postgres/queries/assistant_suggestions"
	"common"
	"context"
	"fmt"
)

func (r *PostgresAssistantSuggestionRepository) Edit(ctx context.Context, id common.AssistantSuggestionsID, message, status string) (*entity.AssistantSuggestions, error) {
	pgid, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, err
	}

	row, err := r.getQueries(ctx).EditAssistantSuggestion(ctx, queries.EditAssistantSuggestionParams{
		ID: pgid,
		Message: message,
		Status: status,

	})
	if err != nil {
		fmt.Printf("Counldnot edit assistant suggestion:%v", err)
		return nil, err
	}

	editAssissugg, err := rowToDomainAssistantSuggestion(
		row.ID,
		row.DaySessionID,
		row.Message,
		row.Status,
		row.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return editAssissugg, nil
}