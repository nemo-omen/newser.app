package main

import (
	"flag"

	// "github.com/alexedwards/scs/v2"
	// "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"newser.app/cmd/web/handler"
	custommiddleware "newser.app/cmd/web/middleware"
)

func main() {
	addr := flag.String("addr", ":4321", "HTTP network address")
	dev := flag.Bool("dev", false, "Whether to run the server in development mode")
	dsn := flag.String("dsn", "internal/data/newser.sqlite", "Sqlite dataa source name")
	flag.Parse()

	conf := custommiddleware.NewConfig(*dev, *dsn)

	app := echo.New()
	setLogLevel(app, *dev)
	app.Static("/static", "ui/static")
	app.Use(conf.SetConfig)
	app.Use(custommiddleware.ContextValue)
	setHandlers(app, *dsn)

	app.Logger.Fatal(app.Start(*addr))
}

func setLogLevel(app *echo.Echo, dev bool) {
	if l, ok := app.Logger.(*log.Logger); ok {
		if dev {
			l.SetLevel(log.DEBUG)
			app.Logger.Debugf("Is Dev?: %v", dev)
		} else {
			l.SetLevel(log.INFO)
		}
	}
}

func setHandlers(app *echo.Echo, dsn string) {
	homeHandler := handler.HomeHandler{}
	authHandler := handler.NewAuthHandler(dsn)

	app.GET("/", homeHandler.Home)
	authGroup := app.Group("/auth")
	authGroup.GET("/signup", authHandler.GetSignup)
	authGroup.POST("/signup", authHandler.PostSignup)
	authGroup.GET("/login", authHandler.GetLogin)
}
