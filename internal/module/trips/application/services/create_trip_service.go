package services

import (
	"context"
	"trip/application/command"
	"trip/application/dto"
	"trip/domain/entity"
	"trip/domain/repository"
)

type CreateTripService struct {
	tripRepo repository.TripRepository
}

func NewCreateTripService(tripRepo repository.TripRepository) *CreateTripService {
	return &CreateTripService{tripRepo: tripRepo}
}

func (s *CreateTripService) CreateTrip(ctx context.Context, cmd command.CreateTripCommand) (dto.TripDTO, error) {
	trip, err := entity.NewTrip(cmd.Destination, cmd.StartDate, cmd.EndDate, cmd.TravelersCount)
	if err != nil {
		return dto.TripDTO{}, err
	}

	if err := s.tripRepo.Create(ctx, trip); err != nil {
		return dto.TripDTO{}, err
	}

	return dto.ToTripDTO(trip), nil
}
