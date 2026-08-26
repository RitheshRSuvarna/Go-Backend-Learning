package services

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"common"

// 	"day_session/application/command"
// 	"day_session/application/dto"
// 	"day_session/domain/entity"
// )

// // Fake repository for testing.
// type fakeDaySessionRepository struct {
// 	createCalled bool
// 	createError  error
// 	created      *entity.DaySession
// }

// func (f *fakeDaySessionRepository) CreateDaySession(
// 	ctx context.Context,
// 	daySession *entity.DaySession,
// ) error {
// 	f.createCalled = true
// 	f.created = daySession

// 	return f.createError
// }

// func (f *fakeDaySessionRepository) GetDaySession(
//     ctx context.Context,
//     id common.DaySessionID,
// ) (*entity.DaySession, error) {
//     return nil, nil
// }

// // Implement other methods required by
// // DaySessionRepository here.
// // Add them if your interface contains more methods.

// func TestCreateDaySessionService_Success(t *testing.T) {
// 	repo := &fakeDaySessionRepository{}

// 	service := NewDaySessionService(repo)

// 	cmd := command.CreateDaySessionCommand{
// 		TripID:     "trip-1",
// 		Date:       "2026-08-20",
// 		StartTime:  "09:00",
// 		StartLabel: "Hotel",
// 	}

// 	result, err := service.CreateDaySession(
// 		context.Background(),
// 		cmd,
// 	)

// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	if !repo.createCalled {
// 		t.Fatal("expected repository CreateDaySession to be called")
// 	}

// 	if repo.created == nil {
// 		t.Fatal("expected DaySession to be passed to repository")
// 	}

// 	if result.TripID.String() != "trip-1" {
// 		t.Errorf(
// 			"expected trip ID trip-1, got %s",
// 			result.TripID.String(),
// 		)
// 	}

// 	if result.Date != "2026-08-20" {
// 		t.Errorf(
// 			"expected date 2026-08-20, got %s",
// 			result.Date,
// 		)
// 	}

// 	if result.StartTime != "09:00" {
// 		t.Errorf(
// 			"expected start time 09:00, got %s",
// 			result.StartTime,
// 		)
// 	}

// 	if result.StartLabel != "Hotel" {
// 		t.Errorf(
// 			"expected start label Hotel, got %s",
// 			result.StartLabel,
// 		)
// 	}
// }

// func TestCreateDaySessionService_InvalidCommand(t *testing.T) {
// 	repo := &fakeDaySessionRepository{}

// 	service := NewDaySessionService(repo)

// 	cmd := command.CreateDaySessionCommand{
// 		TripID:     "",
// 		Date:       "2026-08-20",
// 		StartTime:  "09:00",
// 		StartLabel: "Hotel",
// 	}

// 	result, err := service.CreateDaySession(
// 		context.Background(),
// 		cmd,
// 	)

// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}

// 	if repo.createCalled {
// 		t.Fatal(
// 			"repository CreateDaySession should not be called when validation fails",
// 		)
// 	}

// 	if result != (dto.DaySessionDTO{}) {
// 		t.Fatal("expected empty DaySessionDTO")
// 	}
// }

// func TestCreateDaySessionService_RepositoryError(t *testing.T) {
// 	expectedError := errors.New("database error")

// 	repo := &fakeDaySessionRepository{
// 		createError: expectedError,
// 	}

// 	service := NewDaySessionService(repo)

// 	cmd := command.CreateDaySessionCommand{
// 		TripID:     "trip-1",
// 		Date:       "2026-08-20",
// 		StartTime:  "09:00",
// 		StartLabel: "Hotel",
// 	}

// 	result, err := service.CreateDaySession(
// 		context.Background(),
// 		cmd,
// 	)

// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}

// 	if !errors.Is(err, expectedError) {
// 		t.Errorf(
// 			"expected error %v, got %v",
// 			expectedError,
// 			err,
// 		)
// 	}

// 	if result != (dto.DaySessionDTO{}) {
// 		t.Fatal("expected empty DaySessionDTO")
// 	}

// 	if !repo.createCalled {
// 		t.Fatal(
// 			"expected repository CreateDaySession to be called",
// 		)
// 	}
// }