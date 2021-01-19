package core

import (
	"context"
	"testing"

	"github.com/purusah/nalipka/internal/config"
	"github.com/purusah/nalipka/pkg/app/db"
	"github.com/purusah/nalipka/pkg/repository"
)

func TestCreateBoard(t *testing.T) {
	conf, _ := config.GetConfig()
	boardNameExpected := "Test1"
	ctx := context.Background()
	db := db.OpenDBPool(ctx, conf.Storage.Url)
	boardID, err := CreateBoard(ctx, db, boardNameExpected)
	if err != nil {
		t.Fatal(err)
	}
	if boardID == 0 {
		t.Fatal("no error with zero id")
	}
	board, err := GetBoard(ctx, db, boardID)
	if err != nil {
		t.Fatal(err)
	}
	if board.Name != boardNameExpected || board.ID != boardID {
		t.Fatalf("board name %s not equal to excpected", board.Name)
	}
}

func TestCreateBoardDeleteBoardAndList(t *testing.T) {
	conf, _ := config.GetConfig()
	boardNameExpected := "Test1"
	ctx := context.Background()
	db := db.OpenDBPool(ctx, conf.Storage.Url)
	boardID, err := CreateBoard(ctx, db, boardNameExpected)
	if err != nil {
		t.Fatal(err)
	}
	if boardID == 0 {
		t.Fatal("no error with zero id")
	}
	lists, err := GetListsByBoardID(ctx, db, boardID)
	if err != nil {
		t.Fatal(err)
	}
	if len(lists) != 1 {
		t.Fatal("expected 1 list at new board, got", len(lists))
	}
	list := lists[0]
	err = DeleteBoard(ctx, db, boardID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = GetBoard(ctx, db, boardID)
	if err != repository.ErrNoRowsFound {
		t.Fatal("expected no rows found got", err)
	}
	_, err = GetListByID(ctx, db, list.ID)
	if err != repository.ErrNoRowsFound {
		t.Fatal("expected no rows found got", err)
	}
}

func TestNewBoardCreateList(t *testing.T) {
	conf, _ := config.GetConfig()
	boardNameExpected := "Test1"
	ctx := context.Background()
	db := db.OpenDBPool(ctx, conf.Storage.Url)
	boardID, err := CreateBoard(ctx, db, boardNameExpected)
	if err != nil {
		t.Fatal(err)
	}
	if boardID == 0 {
		t.Fatal("no error with zero id")
	}
	lists, err := GetListsByBoardID(ctx, db, boardID)
	if err != nil {
		t.Fatal(err)
	}
	if len(lists) != 1 {
		t.Fatal("expected 1 list at new board, got", len(lists))
	}
	list := lists[0]
	if list.Position != 1 {
		t.Fatalf("must be 1st list")
	}
	listByID, err := GetListByID(ctx, db, list.ID)
	if err != nil {
		t.Fatal(err)
	}
	if list.Name != listByID.Name || list.ID != listByID.ID {
		t.Fatal("lists must be same expected", list, "got", listByID)
	}
}

func TestRemovingBoardRemoveList(t *testing.T) {
	conf, _ := config.GetConfig()
	boardNameExpected := "Test1"
	ctx := context.Background()
	db := db.OpenDBPool(ctx, conf.Storage.Url)
	boardID, err := CreateBoard(ctx, db, boardNameExpected)
	if err != nil {
		t.Fatal(err)
	}
	lists, err := GetListsByBoardID(ctx, db, boardID)
	if err != nil {
		t.Fatal(err)
	}
	list := lists[0]
	err = DeleteBoard(ctx, db, boardID)
	if err != nil {
		t.Fatalf("board must be deleted got %v", err)
	}
	_, err = GetListByID(ctx, db, list.ID)
	if err != repository.ErrNoRowsFound {
		t.Fatalf("must not find any list got %v", err)
	}
	lists, err = GetListsByBoardID(ctx, db, boardID)
	if err != repository.ErrNoRowsFound {
		t.Fatalf("must not find any list got %v", lists)
	}
}

func TestDeleteBoardNotExist(t *testing.T) {
	conf, _ := config.GetConfig()
	db := db.OpenDBPool(context.Background(), conf.Storage.Url)
	err := DeleteBoard(context.Background(), db, 6789)
	if err == repository.ErrNoRowsAffected {
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal("row must not exist")
}
