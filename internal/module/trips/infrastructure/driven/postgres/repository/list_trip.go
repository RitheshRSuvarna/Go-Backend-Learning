package repository

import (
	"context"
	"fmt"
	"trip/domain/entity"
)

func (r *PostgresTripRepository) List(ctx context.Context) ([]*entity.Trip, error) {

	rows, err := r.getQueries(ctx).ListTrips(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to get trip:%w", err)
	}

	out := make([]*entity.Trip, 0, len(rows))
	for _, row := range rows {
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
		out = append(out, trip)
	}
	return out, nil
}
