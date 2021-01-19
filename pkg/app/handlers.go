package app

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/purusah/nalipka/pkg/app/core"
	"github.com/purusah/nalipka/pkg/repository"
	"github.com/purusah/nalipka/pkg/router"
)

const maxLimit = 10

func boardRootHandler(app App, prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if prefix != r.URL.Path {
			router.HTTPErrorNotFound(w)
		}
		switch r.Method {
		case http.MethodGet:
			params := r.URL.Query()
			limit, err := strconv.Atoi(params.Get("limit"))
			if err != nil {
				router.HTTPErrorBadRequest(w)
				return
			}
			offset, err := strconv.Atoi(params.Get("offset"))
			if err != nil {
				router.HTTPErrorBadRequest(w)
				return
			}
			boards, err := core.ListBoards(r.Context(), app.db, limit, offset)
			if len(boards) == 0 {
				router.HTTPErrorNotFound(w)
				return
			}
			router.HTTPOk(w, boards)
		case http.MethodPost:
			// TODO
		default:
			router.HTTPErrorMethodNotAllowed(w)
		}
	}
}

func boardIDHandler(app App, prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		boardIDExpected := strings.TrimPrefix(r.URL.Path, prefix)
		boardID, err := strconv.Atoi(boardIDExpected)
		if err != nil {
			router.HTTPErrorBadRequest(w)
			return
		}
		switch r.Method {
		case http.MethodGet:
			board, err := core.GetBoard(r.Context(), app.db, boardID)
			if err == repository.ErrNoRowsFound {
				router.HTTPErrorNotFound(w)
				return
			}
			router.HTTPOk(w, board)
		case http.MethodDelete:
			err := core.DeleteBoard(r.Context(), app.db, boardID)
			if err == repository.ErrNoRowsAffected {
				router.HTTPErrorNotFound(w)
				return
			}
			router.HTTPOk(w, "")
		default:
			router.HTTPErrorMethodNotAllowed(w)
		}
	}
}
