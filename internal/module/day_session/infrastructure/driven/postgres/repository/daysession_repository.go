package repository

import (
	"context"
	platformpostgres "postgres"
	"day_session/domain/repository"
	dsqueries "day_session/infrastructure/driven/postgres/queries/day_session"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDaySessionRepository struct {
	pool *pgxpool.Pool
	queries *dsqueries.Queries
}

func NewDaySessionRepository(pool *pgxpool.Pool) *PostgresDaySessionRepository {
	return &PostgresDaySessionRepository{
		pool: pool,
		queries: dsqueries.New(pool),
	}
}

func (r *PostgresDaySessionRepository) getQueries(ctx context.Context) *dsqueries.Queries {
	if tx, ok := platformpostgres.GetTx(ctx); ok {
		return dsqueries.New(tx)
	}
	return r.queries
}

var _ repository.DaySessionRepository = (*PostgresDaySessionRepository)(nil)
