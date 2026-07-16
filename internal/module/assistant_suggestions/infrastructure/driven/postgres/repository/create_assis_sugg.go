package repository

import (
	"assistant_suggestions/domain/entity"
	asqueries "assistant_suggestions/infrastructure/driven/postgres/queries/assistant_suggestions"
	"common"
	"context"
	"fmt"
)

func (r *PostgresAssistantSuggestionRepository) Create(ctx context.Context, assistantsugg *entity.AssistantSuggestions) error {
	daysessionid, err := uuidStringToPgUUID(assistantsugg.DaysessionID().String())
	if err != nil {
		return  common.NewValidationError("Invlaid id", err)
	}

	row, err :=r.getQueries(ctx).CreateAssistantSuggestions(ctx, asqueries.CreateAssistantSuggestionsParams{
		DaySessionID: daysessionid,
		Message: assistantsugg.Message(),
		Status: assistantsugg.Status(),
	})
	if err != nil {
		return fmt.Errorf("Falied to create assistant suggestions:%v", err)
	}

	created, err := rowToDomainAssistantSuggestion(
		row.ID,
		row.DaySessionID,
		row.Message,
		row.Status,
		row.CreatedAt,
	)
	if err != nil {
		return err
	}
	assistantsugg.SetID(created.ID())
	return nil
}