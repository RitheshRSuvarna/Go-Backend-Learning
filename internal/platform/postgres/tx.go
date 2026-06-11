package postgress

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresTxManager struct {
	pool *pgxpool.Pool
}

func NewTXManager(pool *pgxpool.Pool) *PostgresTxManager {
	return &PostgresTxManager{pool: pool}
}

func (m *PostgresTxManager) WithinTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := m.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("Failed to Begin Transaction: %w", err)
	}

	txctx := context.WithValue(ctx, txkey{}, tx)
	if err := fn(txctx); err != nil {
		if Rollback := tx.Rollback(ctx); Rollback != nil {
			return fmt.Errorf("Failed to Rollback Transaction: %w", Rollback)
		}
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("Failed to Commit Transaction: %W", err)
	}
	return nil
}

func GetTx(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(txkey{}).(pgx.Tx)
	return tx, ok
}

type txkey struct{}
