package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Tx ...
type Tx struct {
	t pgx.Tx
}

// QueryRow ...
func (t *Tx) QueryRow(ctx context.Context, sql string, args ...interface{}) Scanner {
	return &Row{r: t.t.QueryRow(ctx, sql, args...)}
}

// Query ...
func (t *Tx) Query(ctx context.Context, sql string, args ...interface{}) (MultiScanner, error) {
	rs, err := t.t.Query(ctx, sql, args...)
	return &Rows{r: rs}, err
}

// Exec ...
func (t *Tx) Exec(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	e, err := t.t.Exec(ctx, sql, args...)
	return e.RowsAffected(), err
}

// Begin ...
func (t *Tx) Begin(ctx context.Context) (Transactional, error) {
	tx, err := t.t.Begin(ctx)
	return &Tx{t: tx}, err
}

// Commit ...
func (t *Tx) Commit(ctx context.Context) error {
	err := t.t.Commit(ctx)
	if err == pgx.ErrTxClosed {
		return ErrTxClosed
	}
	if err == pgx.ErrTxCommitRollback {
		return ErrTxCommitRollback
	}
	return err
}

// Rollback ...
func (t *Tx) Rollback(ctx context.Context) error {
	return t.t.Rollback(ctx)
}
