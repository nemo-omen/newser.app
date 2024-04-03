package app

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/vingarcia/ksql"
	"newser.app/config"
	"newser.app/datasource"
)

type App struct {
	echo *echo.Echo
	db   *ksql.DB
	cfg  config.Config
}

func NewApp(ctx context.Context, cfg config.Config) *App {
	db, err := datasource.NewDatabase(cfg.Database)
	if err != nil {
		panic(err)
	}

	app := &App{
		echo: echo.New(),
		db:   db,
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
