package repository

import (
	"context"
)

// Scanner ...
type Scanner interface {
	Scan(...interface{}) error
}

// MultiScanner ...
type MultiScanner interface {
	Scan(...interface{}) error
	Next() bool
	Close()
}

// QueryableRow ...
type QueryableRow interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) Scanner
}

// Queryable ...
type Queryable interface {
	Query(ctx context.Context, sql string, args ...interface{}) (MultiScanner, error)
}

// Executable ...
type Executable interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (int64, error)
}

// Connectionable ...
type Connectionable interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) Scanner
	Query(ctx context.Context, sql string, args ...interface{}) (MultiScanner, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (int64, error)

	Begin(ctx context.Context) (Transactional, error)
}

// Transactional ...
type Transactional interface {
	Connectionable
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
