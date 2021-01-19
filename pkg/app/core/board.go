package core

import (
	"context"

	"github.com/purusah/nalipka/pkg/repository"
)

const boardListLimit = 10

// Board ...
type Board struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetBoard ...
func GetBoard(ctx context.Context, queryable repository.QueryableRow, boardID int) (board Board, err error) {
	r := queryable.QueryRow(ctx, getBoard, boardID)
	err = r.Scan(&board.ID, &board.Name)
	if err != nil {
		return board, err
	}
	return board, nil
}

// CreateBoard ...
func CreateBoard(ctx context.Context, queryable repository.QueryableRow, boardName string) (boardID int, err error) {
	r := queryable.QueryRow(ctx, createBoard, boardName, defaultListName)
	err = r.Scan(&boardID)
	if err != nil {
		return 0, err
	}
	return boardID, nil
}

// ListBoards ...
func ListBoards(ctx context.Context, queryable repository.Queryable, limit int, offset int) (boards []Board, err error) {
	// TODO
	return boards, nil
}

// DeleteBoard ...
func DeleteBoard(ctx context.Context, queryable repository.Executable, boardID int) error {
	rowsAffected, err := queryable.Exec(ctx, deleteBoard, boardID)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNoRowsAffected
	}
	return nil
}
