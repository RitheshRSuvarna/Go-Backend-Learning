package entity

import (
	"common"
)

type AssistantSuggestions struct {
	id           common.AssistantSuggestionsID
	daysessionID common.DaySessionID
	message      string
	status       string
	createdAt    common.Time
}

func NewAssistantSuggestions(DaysessionID, message, status string) (*AssistantSuggestions, error) {
	// id, err := common.NewAssistantSuggestionsID()
	// if err != nil {
	// 	return nil, common.NewValidationError("Invalid assistant suggestion id", err)
	// }
	daysessionID, err := common.NewDaySessionID(DaysessionID)
	if err != nil {
		return nil, common.NewValidationError("Invalid daysession id", err)
	}
	if message == "" {
		return nil, common.NewValidationError("Message cannot be empty", err)
	}
	if status == "" {
		return nil, common.NewValidationError("Status cannot be empty", nil)
	}

	switch status {
	case "pending", "accepted", "snoozed":
	default:
		return nil, common.NewValidationError("Invalid status", nil)
	}

	now := common.Now()

	return &AssistantSuggestions{
		// id:           id,
		daysessionID: daysessionID,
		message:      message,
		status:       status,
		createdAt:    now,
	}, nil
}

func (a *AssistantSuggestions) ID() common.AssistantSuggestionsID { return a.id }
func (a *AssistantSuggestions) DaysessionID() common.DaySessionID { return a.daysessionID }
func (a *AssistantSuggestions) Message() string                   { return a.message }
func (a *AssistantSuggestions) Status() string                    { return a.status }
func (a *AssistantSuggestions) CreatedAt() common.Time            { return a.createdAt }

func (a *AssistantSuggestions) SetID(id common.AssistantSuggestionsID) {
	a.id = id
}

func (a *AssistantSuggestions) AssignPersistance(id common.AssistantSuggestionsID, createdAt common.Time) {
	a.id = id
	a.createdAt = createdAt
}

func RestoreAssistantSuggestions(
	id common.AssistantSuggestionsID,
	daysessionID common.DaySessionID,
	message, status string,
	createdAt common.Time,
) *AssistantSuggestions {
	return &AssistantSuggestions{
		id:           id,
		daysessionID: daysessionID,
		message:      message,
		status:       status,
		createdAt:    createdAt,
	}
}
