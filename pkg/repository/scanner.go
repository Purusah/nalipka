package repository

import (
	"github.com/jackc/pgx/v4"
)

type Row struct {
	r pgx.Row
}

func (s *Row) Scan(src ...interface{}) error {
	err := s.r.Scan(src...)
	if err == pgx.ErrNoRows {
		return ErrNoRowsFound
	}
	return err
}

type Rows struct {
	r pgx.Rows
}

func (s *Rows) Scan(src ...interface{}) error {
	err := s.r.Scan(src...)
	if err == pgx.ErrNoRows {
		return ErrNoRowsFound
	}
	return err
}

func (s *Rows) Next() bool {
	return s.r.Next()
}

func (s *Rows) Close() {
	s.r.Close()
}
