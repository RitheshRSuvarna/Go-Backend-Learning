package dto

import "day_session/drift"

type DaySessionResponseDTO struct {
	DaySession DaySessionDTO
	Status     drift.Status
}
