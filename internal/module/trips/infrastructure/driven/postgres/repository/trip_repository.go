package repository

import (
	"context"

	platformpostgres "postgres"
	"trip/domain/repository"
	tripqueries "trip/infrastructure/driven/postgres/queries/trip"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresTripRepository struct {
	pool    *pgxpool.Pool
	queries *tripqueries.Queries
}

func NewTripRepository(pool *pgxpool.Pool) *PostgresTripRepository {
	return &PostgresTripRepository{
		pool:    pool,
		queries: tripqueries.New(pool),
	}
}

func (r *PostgresTripRepository) getQueries(ctx context.Context) *tripqueries.Queries {
	if tx, ok := platformpostgres.GetTx(ctx); ok {
		return tripqueries.New(tx)
	}
	return r.queries
}

var _ repository.TripRepository = (*PostgresTripRepository)(nil)
