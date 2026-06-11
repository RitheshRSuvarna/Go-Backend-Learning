package services

import (
	"context"
	"trip/application/dto"
	"trip/domain/repository"
)

type ListTripService struct {
	tripRepo repository.TripRepository
}

func NewTripListService(tripRepo repository.TripRepository) *ListTripService {
	return &ListTripService{tripRepo: tripRepo}
}

func (s *ListTripService) ListTrips(ctx context.Context) ([]dto.TripDTO, error) {
	trips, err := s.tripRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]dto.TripDTO, 0, len(trips))
	for _, t := range trips {
		out = append(out, dto.ToTripDTO(t))
	}
	return out, nil

}
