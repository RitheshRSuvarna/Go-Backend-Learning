package dto

import (
	"time"
	"auth/domain/entity"
)

type SessionDTO struct {
	ID string
	UserID string
	TokenHash string
	CreatedAt string
}

func ToSessionDTO(t *entity.Session) SessionDTO {
	return SessionDTO {
		ID: t.ID().String(),
		UserID: t.UID().String(),
		TokenHash: t.Token(),
		CreatedAt: t.CreatedAt().Format(time.RFC3339),
	}
}