package repository

import (
	"context"
	"auth/domain/entity"
	"common"
)

type AuthRepository interface {
	// User operations
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id common.UserID) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)

	// Session operations
	CreateSession(ctx context.Context, session *entity.Session) error
	GetSessionByTokenHash(ctx context.Context, tokenHash string) (*entity.Session, error)
	DeleteSession(ctx context.Context, id common.SessionID) error
}