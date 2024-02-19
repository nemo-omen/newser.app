package main

import (
	"current/custommiddleware"
	"current/handler"

	"github.com/labstack/echo/v4"
)

const (
	port = ":4321"
)

func main() {
	app := echo.New()
	app.Use(custommiddleware.NewMiddlewareContextValue)
	app.Use(custommiddleware.CurrentPath)

	homeHandler := handler.HomeHandler{}
	searchHandler := handler.SearchHandler{}
	wsHandler := handler.WsHandler{}

	app.GET("/", homeHandler.HandleGetIndex)
	app.GET("/search", searchHandler.HandleGetIndex)
	app.POST("/search", searchHandler.HandlePostSearch)
	app.GET("/livereload", wsHandler.HandleWsConnect)

	app.Logger.Fatal(app.Start(port))
}
