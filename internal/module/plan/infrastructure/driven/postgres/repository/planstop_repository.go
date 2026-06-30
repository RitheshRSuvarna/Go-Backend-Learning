package repository

import (
	"context"
	platformpostgres "postgres"
	"plan/domain/repository"
	psqueries "plan/infrastructure/driven/postgres/queries/plans"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPlanStopRepository struct {
	pool *pgxpool.Pool
	queries *psqueries.Queries
}

func NewPostgresPlanStopRepository(pool *pgxpool.Pool) *PostgresPlanStopRepository{
	return &PostgresPlanStopRepository{
		pool: pool,
		queries: psqueries.New(pool),
	}
}

func (r *PostgresPlanStopRepository) getQueries(ctx context.Context) *psqueries.Queries {
	if tx, ok := platformpostgres.GetTx(ctx); ok {
		return psqueries.New(tx)
	} 
	return r.queries
}

var _ repository.PlanStopRepository = (*PostgresPlanStopRepsitory)(nil)