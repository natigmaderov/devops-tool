package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type PostgresStorage struct {
	Pool *pgxpool.Pool
}

func NewPostgresStorage(cfg PostgresConfig) (*PostgresStorage, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	err = pool.Ping(context.Background())

	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("error pinging the database: %w", err)
	}

	log.Println("Connected to the database")
	return &PostgresStorage{Pool: pool}, nil
}
