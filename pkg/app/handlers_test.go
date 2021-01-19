package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/purusah/nalipka/internal/config"
	"github.com/purusah/nalipka/pkg/app/core"
	"github.com/purusah/nalipka/pkg/app/db"
)

type boardResponse struct {
	Data core.Board `json:"data"`
}

func TestBoardIDHandlerGetNotFound(t *testing.T) {
	conf, _ := config.GetConfig()
	ctx := context.Background()
	app := App{db.OpenDBPool(ctx, conf.Storage.Url)}
	req, err := http.NewRequest("GET", "/api/v1/board/1234", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()

	http.HandlerFunc(boardIDHandler(app, "/api/v1/board/")).ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestBoardIDHandlerGetOk(t *testing.T) {
	conf, _ := config.GetConfig()
	ctx := context.Background()
	app := App{db.OpenDBPool(ctx, conf.Storage.Url)}
	expectedBoard := core.Board{Name: "test1"}
	boardID, err := core.CreateBoard(ctx, app.db, expectedBoard.Name)
	expectedBoard.ID = boardID
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/board/%d", expectedBoard.ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()

	http.HandlerFunc(boardIDHandler(app, "/api/v1/board/")).ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	respBody := boardResponse{}
	err = json.Unmarshal(resp.Body.Bytes(), &respBody)
	if err != nil {
		t.Fatal(err)
	}

	if respBody.Data.ID != expectedBoard.ID || respBody.Data.Name != expectedBoard.Name {
		t.Fatalf("expected same board")
	}
}
