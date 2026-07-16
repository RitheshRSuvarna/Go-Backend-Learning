package repository

import (
	"assistant_suggestions/domain/repository"
	asqueries "assistant_suggestions/infrastructure/driven/postgres/queries/assistant_suggestions"
	"context"
	platformpostgres "postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresAssistantSuggestionRepository struct {
	pool    *pgxpool.Pool
	queries *asqueries.Queries
}

func NewAssistantSuggestionRepository(pool *pgxpool.Pool) *PostgresAssistantSuggestionRepository {
	return &PostgresAssistantSuggestionRepository{
		pool:    pool,
		queries: asqueries.New(pool),
	}
}

func (r *PostgresAssistantSuggestionRepository) getQueries(ctx context.Context) *asqueries.Queries {
	if tx, ok := platformpostgres.GetTx(ctx); ok {
		return asqueries.New(tx)
	}
	return r.queries
}

var _ repository.AssistantSuggestionRepository = (*PostgresAssistantSuggestionRepository)(nil)
