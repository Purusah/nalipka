package core

import (
	"context"

	"github.com/purusah/nalipka/pkg/repository"
)

const defaultListName = "Your TODOs"

// List ...
type List struct {
	ID       int
	Name     string
	Position int
}

// GetListByID ...
func GetListByID(ctx context.Context, queryable repository.QueryableRow, listID int) (list List, err error) {
	r := queryable.QueryRow(ctx, getListByID, listID)
	err = r.Scan(&list.ID, &list.Name, &list.Position)
	if err != nil {
		return list, err
	}
	return list, nil
}

// GetListsByBoardID ...
func GetListsByBoardID(ctx context.Context, queryable repository.Queryable, boardID int) (lists []List, err error) {
	var list List
	rows, err := queryable.Query(ctx, getBoardLists, boardID)
	if err != nil {
		return lists, err
	}
	for rows.Next() {
		err = rows.Scan(&list.ID, &list.Name, &list.Position)
		if err != nil {
			return lists, err
		}
		lists = append(lists, list)
	}
	if len(lists) == 0 {
		return lists, repository.ErrNoRowsFound
	}
	return lists, nil
}

// CreateList ...
func CreateList(ctx context.Context, q repository.QueryableRow, list List, boardID int) (listID int, err error) {
	r := q.QueryRow(ctx, createList, list.Name, list.Position, boardID)
	err = r.Scan(&listID)
	if err != nil {
		return listID, err
	}
	return listID, nil
}

// DeleteList ...
func DeleteList(ctx context.Context, txs repository.Connectionable, listID int) error {
	var iterList List
	var currentList, previousList List
	amount := 0

	tx, err := txs.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return err
	}
	// TODO Maybe rewrite to batched request
	rows, err := tx.Query(ctx, getListForDelete, listID)
	if err != nil {
		return err
	}
	for rows.Next() {
		err = rows.Scan(&iterList.ID, &iterList.Position, &amount)
		if err != nil {
			return err
		}
		if iterList.ID == listID {
			currentList.ID = iterList.ID
			currentList.Position = iterList.Position
		} else {
			previousList.ID = iterList.ID
			previousList.Position = iterList.Position
		}
	}
	if currentList.ID == 0 {
		return repository.ErrNoRowsFound
	}
	if amount == 1 {
		return ErrListNotRemovable
	}
	if previousList.ID == 0 {
		// TODO What to do with list that most left but other lists exist
		return ErrListNotRemovable
	}
	_, err = tx.Exec(ctx, updateTicketsList, previousList.ID, currentList.ID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, deleteList, currentList.ID)
	if err != nil {
		return err
	}
	tx.Commit(ctx)
	return nil
}
