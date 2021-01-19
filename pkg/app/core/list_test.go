package core

import (
	"context"
	"testing"

	"github.com/purusah/nalipka/internal/config"
	"github.com/purusah/nalipka/pkg/app/db"
	"github.com/purusah/nalipka/pkg/repository"
)

func TestDeleteLastOnlyListNotRemovable(t *testing.T) {
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
	list := lists[0]
	err = DeleteList(ctx, db, list.ID)
	if err == ErrListNotRemovable {
		return
	}
	t.Fatalf("expected not removable list got %v", err)
}

func TestDeleteNotExistedList(t *testing.T) {
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
	list := lists[0]
	err = DeleteList(ctx, db, list.ID+1)
	if err == repository.ErrNoRowsFound {
		return
	}
	t.Fatalf("expected not removable list got %v", err)
}

func TestDeleteNotLastNotOnlyListOk(t *testing.T) {
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
	firstList := lists[0]
	secondList := List{Name: defaultListName, Position: 2}
	secondList.ID, err = CreateList(ctx, db, secondList, boardID)
	if err != nil {
		t.Fatal(err)
	}
	expectedTicket := Ticket{Name: "Ticket1", Position: 1}
	expectedTicket.ID, err = CreateTicket(ctx, db, expectedTicket, secondList.ID)
	if err != nil {
		t.Fatal(err)
	}
	err = DeleteList(ctx, db, secondList.ID)
	if err != nil {
		t.Fatal(err)
	}
	tickets, err := GetTicketsByListID(ctx, db, firstList.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(tickets) != 1 {
		t.Fatalf("expected 1 ticket got %d", len(tickets))
	}
	ticket := tickets[0]
	if ticket.ID != expectedTicket.ID || ticket.Name != expectedTicket.Name {
		t.Fatalf("tickets must be same")
	}
}

func TestDeleteLastNotOnlyListRemovable(t *testing.T) {
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
	firstList := lists[0]
	secondList := List{Name: defaultListName, Position: 2}
	secondList.ID, err = CreateList(ctx, db, secondList, boardID)
	if err != nil {
		t.Fatal(err)
	}
	expectedTicket := Ticket{Name: "Ticket1", Position: 1}
	expectedTicket.ID, err = CreateTicket(ctx, db, expectedTicket, firstList.ID)
	if err != nil {
		t.Fatal(err)
	}
	err = DeleteList(ctx, db, firstList.ID)
	if err == ErrListNotRemovable {
		// TODO Or change core logic
		return
	}
	t.Fatal(err)
}
