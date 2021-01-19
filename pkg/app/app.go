package app

import (
	"context"
	"log"
	"net/http"

	"github.com/purusah/nalipka/internal/config"
	"github.com/purusah/nalipka/pkg/app/db"
	"github.com/purusah/nalipka/pkg/repository"
)

type App struct {
	db *repository.Repository
}

func DefaultHandler(_ http.ResponseWriter, _ *http.Request) {}

// StartApp ...
func StartApp() {
	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	app := App{db: db.OpenDBPool(context.Background(), conf.Storage.Url)}

	srv := http.NewServeMux()
	srv.Handle("/api/v1/board", boardRootHandler(app, "/api/v1/board")) // create, list (sorted id, name)
	srv.Handle("/api/v1/board/", boardIDHandler(app, "/api/v1/board/"))
	srv.Handle("/api/v1/list/:id", http.HandlerFunc(DefaultHandler)) // list
	srv.Handle("/api/v1/ticket/:id", http.HandlerFunc(DefaultHandler))
	srv.Handle("/api/v1/ticket/:id/comment/:id", http.HandlerFunc(DefaultHandler))

	if err = http.ListenAndServe(conf.Service.Url, srv); err != nil {
		log.Fatal(err)
	}
}
