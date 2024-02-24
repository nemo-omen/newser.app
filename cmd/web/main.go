package main

import (
	"database/sql"
	"flag"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"newser.app/cmd/web/core"
	"newser.app/internal/repository"
	"newser.app/internal/service"
)

func main() {
	addr := flag.String("addr", "4321", "HTTP network address")
	dev := flag.Bool("dev", false, "Whether to run the server in development mode")
	dsn := flag.String("dsn", "/internal/data/newser.sqlite", "Sqlite dataa source name")
	flag.Parse()

	logger := core.Logger(*dev)

	db, err := openDb(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &core.App{
		Logger:      logger,
		FeedRepo:    &repository.NewsfeedSqliteRepo{DB: db},
		FeedService: &service.API{Client: &http.Client{Timeout: 5 * time.Second}},
	}

	app.Logger.Info("starting server", "addr", *addr)
	app.Logger.Debug("dev mode", "dev", *dev)
	err = http.ListenAndServe(":"+*addr, app.Routes())
	app.Logger.Error(err.Error())
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "dsn")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
