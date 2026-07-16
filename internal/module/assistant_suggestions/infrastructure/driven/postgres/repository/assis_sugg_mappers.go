package repository

import (
	"common"
	"assistant_suggestions/domain/entity"

	"github.com/jackc/pgx/v5/pgtype"
)

func rowToDomainAssistantSuggestion(
	id pgtype.UUID,
	daysessionID pgtype.UUID,
	message, status string,
	createdAt pgtype.Timestamptz,
) (*entity.AssistantSuggestions, error) {
	suggid, err := common.NewAssistantSuggestionsID(pgTypeToString(id))
	if err != nil {
		return nil, common.NewValidationError("Invalid assistant suggestion id:%v", err)
	}
	daysessionid, err := common.NewDaySessionID(pgTypeToString(daysessionID))
	if err != nil {
		return nil, common.NewValidationError("Invalid daysession id:%v", err)
	}
	if !createdAt.Valid {
		return nil, common.NewValidationError("invalid createdAt:%v", err)
	}
	
	return entity.RestoreAssistantSuggestions(
		suggid,
		daysessionid,
		message,
		status,
		common.NewTime(createdAt.Time),
	), nil
}