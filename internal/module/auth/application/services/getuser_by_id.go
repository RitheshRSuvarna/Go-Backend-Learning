package service

import (
	"auth/application/dto"
	"auth/domain/repository"
	"common"
	"context"
)

type GetUserByIDService struct {
	userrepo repository.AuthRepository
}

func (t *GetUserByIDService) getuserByID(ctx context.Context, id common.UserID) (dto.UserDTO, error) {
	user, err := t.userrepo.GetUserByID(ctx, id)
	if err != nil {
		return dto.UserDTO{}, err
	}
	return dto.ToUserDTO(user), nil
}
