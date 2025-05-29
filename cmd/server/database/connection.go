package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type DBConfig struct {
	DSN string
}

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	config := DBConfig{
		"host=chat_service_db port=5432 dbname=chat-service user=admin password=Egalam47 sslmode=disable",
	}

	pool, err := pgxpool.New(ctx, config.DSN)
	if err != nil {
		log.Fatalf("pgxpool.New: %v", err)
	}

	return pool, nil
}
