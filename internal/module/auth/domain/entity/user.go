package entity

import (
	"common"
)

type User struct {
	id        common.UserID
	name      string
	email     string
	password  string
	createdAt common.Time
}

func NewUser(name, email, password string) (*User, error) {
	if name == "" {
		return nil, common.NewValidationError("Name is required", nil)
	}
	if email == "" {
		return nil, common.NewValidationError("Email is required", nil)
	}
	if password == "" {
		return nil, common.NewValidationError("Enter Password", nil)
	}

	now := common.Now()

	return &User{
		name:      name,
		email:     email,
		password:  password,
		createdAt: now,
	}, nil
}

func (s *User) ID() common.UserID      { return s.id }
func (s *User) UName() string          { return s.name }
func (s *User) Mail() string           { return s.email }
func (s *User) Pass() string           { return s.password }
func (s *User) CreatedAt() common.Time { return s.createdAt }

func (s *User) SetID(id common.UserID) {
	s.id = id
}

func (s *User) AssignPersistance(id common.UserID, createdAt common.Time) {
	s.id = id
	s.createdAt = createdAt
}

func RestoreUser(
	id common.UserID,
	name, email, password string,
	createdAt common.Time,
) *User {
	return &User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}
}
