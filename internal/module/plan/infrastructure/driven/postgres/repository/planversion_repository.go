package repository

import(
	"context"
	platformpostgres "postgres"
	"plan/domain/repository"
	pvqueries "plan/infrastructure/driven/postgres/queries/plans"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPlanVersionRepository struct {
	pool *pgxpool.Pool
	queries *pvqueries.Queries
}

func NewPostgresPlanVersionRepository(pool *pgxpool.Pool) *PostgresPlanVersionRepository {
	return &PostgresPlanVersionRepository{
		pool: pool,
		queries: pvqueries.New(pool),
	}
}

func (r *PostgresPlanVersionRepository) getQueries(ctx context.Context) *pvqueries.Queries {
	if tx, ok := platformpostgres.GetTx(ctx); ok {
		return pvqueries.New(tx)
	}
	return r.queries
}

var _ repository.PlanVersionRepository = (*PostgresPlanVersionRepository)(nil)