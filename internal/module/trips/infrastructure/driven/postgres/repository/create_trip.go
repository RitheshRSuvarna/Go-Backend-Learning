package repository

import (
	"common"
	"context"
	"fmt"
	"trip/domain/entity"
	tripqueries "trip/infrastructure/driven/postgres/queries/trip"
)

func (r *PostgresTripRepository) Create(ctx context.Context, trip *entity.Trip) error {
	startDate, err := dateStringToPGDate(trip.StartDate())
	if err != nil {
		return common.NewValidationError("Invalid start_date", err)
	}

	endDate, err := dateStringToPGDate(trip.EndDate())
	if err != nil {
		return common.NewValidationError("Invalid end_date", err)
	}

	row, err := r.getQueries(ctx).CreateTrip(ctx, tripqueries.CreateTripParams{
		Destination:    trip.Destination(),
		StartDate:      startDate,
		EndDate:        endDate,
		TravelersCount: int32(trip.TravelersCount()),
	})
	if err != nil {
		fmt.Println("CreateTrip DB Error:", err)
		return fmt.Errorf("failed to create trip: %w", err)
	}

	created, err := rowToDomainTrip(
		row.ID,
		row.Destination,
		row.StartDate,
		row.EndDate,
		row.TravelersCount,
		row.CreatedAt,
	)
	if err != nil {
		return err
	}
	trip.SetID(created.ID())
	return nil
}
