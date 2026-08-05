package service

import (
	"auth/domain/repository"
	"common"
	"context"
)

type DeleteSessionService struct {
	sessionrepo repository.AuthRepository
}

func (s *DeleteSessionService) deleteSession(ctx context.Context, id common.SessionID) error {
	if err := s.sessionrepo.DeleteSession(ctx, id); err != nil {
		return err
	}
	return nil
}
