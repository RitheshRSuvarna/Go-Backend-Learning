package services

import (
	"common"
	"context"
	"day_session/application/dto"
	"day_session/domain/entity"
	"day_session/domain/repository"
	"day_session/application/port"
	trip "trip/domain/repository"


)
type GenerateDaySessionsService struct {
	generator      port.DaySessionGenerator
	tripRepo       trip.TripRepository
	daySessionRepo repository.DaySessionRepository
}
func NewGenerateDaySessionsService(
	generator port.DaySessionGenerator,
	tripRepo trip.TripRepository,
	daySessionRepo repository.DaySessionRepository,
) *GenerateDaySessionsService {
	return &GenerateDaySessionsService{
		generator:      generator,
		tripRepo:       tripRepo,
		daySessionRepo: daySessionRepo,
	}
}

func (s *GenerateDaySessionsService) Generate(
    ctx context.Context,
    tripID common.TripID,
) ([]dto.DaySessionDTO, error) {

    trip, err := s.tripRepo.GetByID(ctx, tripID)
    if err != nil {
        return nil, err
    }

    generated, err := s.generator.Generate(
        ctx,
        port.GenerateDaySessionsInput{
            TripID:      trip.ID().String(),
            Destination: trip.Destination(),
            StartDate:   trip.StartDate(),
            EndDate:     trip.EndDate(),
            Travelers:   trip.TravelersCount(),
        },
    )

    if err != nil {
        return nil, err
    }

    result := make([]dto.DaySessionDTO, 0)

    for _, day := range generated {

        ds, err := entity.NewDaySession(
            trip.ID().String(),
            day.Date,
            day.StartTime,
            day.StartLabel,
        )
        if err != nil {
            return nil, err
        }

        if err := s.daySessionRepo.CreateDaySession(
            ctx,
            ds,
        ); err != nil {
            return nil, err
        }

        result = append(
            result,
            dto.ToDaySessionDTO(ds),
        )
    }

    return result, nil
}