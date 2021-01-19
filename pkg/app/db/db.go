package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/purusah/nalipka/pkg/repository"
)

// OpenDBPool ...
func OpenDBPool(ctx context.Context, url string) *repository.Repository {
	confDb, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Panicf("can't parse db config %e", err)
	}
	confDb.MaxConns = 2
	confDb.MinConns = 1
	conn, err := pgxpool.ConnectConfig(ctx, confDb)
	if err != nil {
		log.Panicf("can't connect db %e", err)
	}

	return repository.NewRepository(conn)
}
