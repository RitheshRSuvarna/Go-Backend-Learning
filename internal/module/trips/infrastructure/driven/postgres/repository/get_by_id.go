package repository

import(
	"context"
	"fmt"
	"common"
	"trip/domain/entity"
)

func (r *PostgresTripRepository) GetByID(ctx context.Context, id common.TripID) (*entity.Trip, error) {
	tripid, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return nil, fmt.Errorf("Failed to convert trip id to pg uuid:%w", err)
	}
	row, err := r.getQueries(ctx).GetTripByID(ctx, tripid)
	if err != nil {
		return nil, fmt.Errorf("Failed to get trip by id:%w", err)
	}

	trip, err := rowToDomainTrip (
		row.ID,
		row.Destination,
		row.StartDate,
		row.EndDate,
		row.TravelersCount,
		row.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return trip, nil
}