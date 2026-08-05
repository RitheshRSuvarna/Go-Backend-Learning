package dto

import (
	"auth/domain/entity"
	"time"
)

type UserDTO struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt string
}

func ToUserDTO(t *entity.User) UserDTO {
	return UserDTO{
		ID:        t.ID().String(),
		Name:      t.UName(),
		Email:     t.Mail(),
		Password:  t.Pass(),
		CreatedAt: t.CreatedAt().Format(time.RFC3339),
	}
}
