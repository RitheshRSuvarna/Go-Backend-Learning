package postgress

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	DATABASE_URL string
}

func NewDB(ctx context.Context, config PostgresConfig) (*pgxpool.Pool, error) {
	if config.DATABASE_URL == "" {
		return nil, fmt.Errorf("DATABASE_URL  is Required")
	}
	db, err := pgxpool.New(ctx, config.DATABASE_URL)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to  DB")
	}
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Failed to Ping DB")
	}
	return db, nil

}
