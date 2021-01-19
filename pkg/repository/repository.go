package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Repository ...
type Repository struct {
	p *pgxpool.Pool
}

// Query ...
func (r *Repository) Query(ctx context.Context, sql string, args ...interface{}) (MultiScanner, error) {
	rs, err := r.p.Query(ctx, sql, args...)
	return &Rows{r: rs}, err
}

// QueryRow ...
func (r *Repository) QueryRow(ctx context.Context, sql string, args ...interface{}) Scanner {
	return &Row{r: r.p.QueryRow(ctx, sql, args...)}
}

// Exec ...
func (r *Repository) Exec(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	t, err := r.p.Exec(ctx, sql, args...)
	return t.RowsAffected(), err
}

// Begin ...
func (r *Repository) Begin(ctx context.Context) (Transactional, error) {
	tx, err := r.p.Begin(ctx)
	return &Tx{t: tx}, err
}

// NewRepository ...
func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{p: pool}
}
