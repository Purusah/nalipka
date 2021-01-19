package repository

import "errors"

// ErrNoRowsAffected ...
var ErrNoRowsAffected = errors.New("no row affected")

// ErrNoRowsFound ...
var ErrNoRowsFound = errors.New("no rows found")

// ErrTxClosed ...
var ErrTxClosed = errors.New("transaction closed")

// ErrTxCommitRollback ...
var ErrTxCommitRollback = errors.New("transaction commit rollback")
