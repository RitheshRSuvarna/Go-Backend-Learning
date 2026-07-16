package repository

import (
	"assistant_suggestions/domain/entity"
	"common"
	"context"
	"fmt"
)

func (r *PostgresAssistantSuggestionRepository) Get(ctx context.Context, id common.DaySessionID) ([]*entity.AssistantSuggestions, error) {
	pgid, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, err
	}
	rows, err := r.getQueries(ctx).GetAssistantSuggestions(ctx, pgid)
	if err != nil {
		fmt.Printf("Could not get assistant suggestiond:%v", err)
		return nil, err
	}

	assistantsugg := make([]*entity.AssistantSuggestions,0,len(rows))

	for _, row := range rows {
		assistantsug, err := rowToDomainAssistantSuggestion(
			row.ID,
			row.DaySessionID,
			row.Message,
			row.Status,
			row.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		assistantsugg = append(assistantsugg, assistantsug)
	}
	return assistantsugg, nil
}