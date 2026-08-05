package entity

import (
	"common"
)

type Session struct {
	id        common.SessionID
	userID    common.UserID
	tokenHash string
	createdAt common.Time
}

func NewSession(userid common.UserID, token string) (*Session, error) {
	if userid.String() == "" {
		return nil, common.NewValidationError("user id is required", nil)
	}

	if token == "" {
		return nil, common.NewValidationError("hashpassword is required", nil)
	}

	now := common.Now()

	return &Session{
		userID: userid,
		tokenHash: token,
		createdAt: now,
	}, nil
}

func (s *Session) SetID(id common.SessionID) {
	s.id = id
}

func (s *Session) AssignPersistance(id common.SessionID, createdat common.Time) {
	s.id = id
	s.createdAt = createdat
}

func (s *Session) ID() common.SessionID {return s.id}
func (s *Session) UID() common.UserID {return s.userID}
func (s *Session) Token() string {return s.tokenHash}
func (s *Session) CreatedAt() common.Time {return s.createdAt}

func RestoreSession(
	id common.SessionID,
	userID common.UserID,
	tokenHash string,
	createdAt common.Time,
) *Session{
	return &Session{
		id: id,
		userID: userID,
		tokenHash: tokenHash,
		createdAt: createdAt,
	}
}
