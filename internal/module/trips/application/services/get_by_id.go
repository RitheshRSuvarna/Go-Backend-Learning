package services

import (
	"context"
	"common"
	"trip/application/dto"
	"trip/domain/repository"
)

type GetTripByIDService struct {
	tripRepo repository.TripRepository
}

func NewGetTripByIDService(tripRepo repository.TripRepository) *GetTripByIDService {
	return &GetTripByIDService{tripRepo: tripRepo}
}

func (s *GetTripByIDService) GetTripByID(ctx context.Context, id common.TripID) (dto.TripDTO, error) {
	trip, err := s.tripRepo.GetByID(ctx, id)
	if err != nil {
		return dto.TripDTO{}, err
	}
	return dto.ToTripDTO(trip), nil
}