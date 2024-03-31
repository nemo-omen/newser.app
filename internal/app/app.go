package app

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/vingarcia/ksql"
	ksqlite "github.com/vingarcia/ksql/adapters/modernc-ksqlite"
)

type App struct {
	echo *echo.Echo
	db   *ksql.DB
}

func NewApp(ctx context.Context) *App {
	db, err := ksqlite.New(ctx, "file::memory:?cache=shared", ksql.Config{})
	if err != nil {
		log.Fatal("failed to create ksqlite DB: %w", err)
	}

	app := &App{
		echo: echo.New(),
		db:   &db,
	}

	app.echo.Logger.SetLevel(log.DEBUG)
	app.echo.Use(middleware.Logger())
	app.echo.Use(middleware.Recover())

	return app
}

func (a *App) Start() error {
	if err := a.startService(); err != nil {
		a.echo.Logger.Fatal(err)
		return err
	}

	a.echo.Logger.SetLevel(log.INFO)
	return a.echo.Start(":8080")
}
