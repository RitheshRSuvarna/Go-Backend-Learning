package service

import (
	"context"
	"auth/domain/repository"
	"auth/application/dto"
)

type GetUserByEmailService struct {
	userrepo repository.AuthRepository
}

func (u *GetUserByEmailService) getUserByEmail(ctx context.Context, email string) (dto.UserDTO, error) {
	user, err := u.userrepo.GetUserByEmail(ctx, email)
	if err != nil {
		return dto.UserDTO{}, err
	}
	return dto.ToUserDTO(user), nil
}