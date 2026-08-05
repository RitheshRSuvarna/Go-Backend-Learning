package entity

import (
	"common"
)

type PasswordReset struct {
	id         common.PasswordResetID
	userID     common.UserID
	tokenHash  string
	expiresAt  common.Time
	usedAt     *common.Time
	createdAt  common.Time
}

func NewPasswordReset(
	userID common.UserID,
	tokenHash string,
	expiresAt common.Time,
) (*PasswordReset, error) {

	if userID.String() == "" {
		return nil, common.NewValidationError("user id is required", nil)
	}

	if tokenHash == "" {
		return nil, common.NewValidationError("token hash is required", nil)
	}

	now := common.Now()

	return &PasswordReset{
		userID:    userID,
		tokenHash: tokenHash,
		expiresAt: expiresAt,
		createdAt: now,
	}, nil
}

func (s *PasswordReset) SetID(id common.PasswordResetID) {
	s.id = id
}

func (s *PasswordReset) AssignPersistance(id common.PasswordResetID, createdat common.Time) {
	s.id = id
	s.createdAt = createdat
}

func RestorePasswordReset(
	id common.PasswordResetID,
	userID common.UserID,
	tokenHash string,
	expiresAt common.Time,
	usedAt *common.Time,
	createdAt common.Time,
) *PasswordReset {

	return &PasswordReset{
		id:         id,
		userID:     userID,
		tokenHash:  tokenHash,
		expiresAt:  expiresAt,
		usedAt:     usedAt,
		createdAt:  createdAt,
	}
}