package service

import (
	"context"
	"auth/domain/repository"
	"auth/application/dto"
)

type GetSessionByTokenHashService struct {
	sessionrepo repository.AuthRepository
}

func (s *GetSessionByTokenHashService) getBySessionTokenhash(ctx context.Context, token string) (dto.SessionDTO, error) {
	session, err := s.sessionrepo.GetSessionByTokenHash(ctx, token)
	if err != nil {
		return dto.SessionDTO{}, err
	}
	return dto.ToSessionDTO(session), nil
}