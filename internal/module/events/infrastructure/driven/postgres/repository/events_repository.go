package repository

import (
	"context"
	platformpostgres "postgres"
	"events/domain/repository"
	eventqueries "events/infrastructure/driven/postgres/queries/events"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresEventsRepository struct {
	pool *pgxpool.Pool
	queries *eventqueries.Queries
}

func NewEventsRepository(pool *pgxpool.Pool) *PostgresEventsRepository {
	return &PostgresEventsRepository{
		pool: pool,
		queries: eventqueries.New(pool),
	}
}

func (r *PostgresEventsRepository) getQueries(ctx context.Context) *eventqueries.Queries {
	if tx, ok := platformpostgres.GetTx(ctx); ok {
		return eventqueries.New(tx)
	}
	return r.queries
}

var _ repository.EventsRepository = (*PostgresEventsRepository)(nil)