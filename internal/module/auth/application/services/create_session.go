package service

import (
	"context"
	"common"
	"auth/domain/repository"
	"auth/application/command"
	"auth/domain/entity"
	"auth/application/dto"
)

type CreateSessionService struct {
	sessionrepo repository.AuthRepository
}

func (s *CreateSessionService) createSession(ctx context.Context, cmd command.CreateSessionCommand, id common.UserID) (dto.SessionDTO, error) {
	session, err := entity.NewSession(id, cmd.TokenHash)
	if err != nil {
		return dto.SessionDTO{}, err
	}

	if err := s.sessionrepo.CreateSession(ctx, session); err != nil {
		return dto.SessionDTO{}, err
	}
	return dto.ToSessionDTO(session), nil
}