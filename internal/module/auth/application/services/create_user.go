package service

import (
	"context"
	"auth/domain/entity"
	"auth/domain/repository"  
	"auth/application/command"  
	"auth/application/dto"
)

type CreateUserService struct {
	userrepo repository.AuthRepository
}

func NewCreateUser(userrepo repository.AuthRepository) *CreateUserService {
	return &CreateUserService{userrepo: userrepo}
}

func (u *CreateUserService) createuser(ctx context.Context, cmd command.CreateUserCommand) (dto.UserDTO, error) {
	user, err := entity.NewUser(cmd.Name, cmd.Email, cmd.Password)
	if err != nil {
		return dto.UserDTO{}, err
	}
	if err := u.userrepo.CreateUser(ctx, user); err != nil {
		return dto.UserDTO{}, err
	}
	return dto.ToUserDTO(user), nil
}