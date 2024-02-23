package main

import (
	"current/custommiddleware"
	"current/handler"
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	port = ":4321"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Could not load .env file! Exiting.")
	}

	mode := custommiddleware.NewMode(os.Getenv("MODE"))

	app := echo.New()
	app.Static("/public", "public")
	app.Use(mode.SetMode)
	app.Use(custommiddleware.NewMiddlewareContextValue)
	app.Use(custommiddleware.CurrentPath)
	app.Use(custommiddleware.HTMX)
	// app.Use(middleware.Logger())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4321"},
	}))

	app.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))))

	homeHandler := handler.HomeHandler{}
	authHandler := handler.AuthHandler{}
	appHandler := handler.AppHandler{}
	searchHandler := handler.SearchHandler{}
	subHandler := handler.SubscriptionHandler{}
	wsHandler := handler.WsHandler{}

	app.GET("/", homeHandler.HandleGetIndex)
	app.GET("/livereload", wsHandler.HandleWsConnect)

	appGroup := app.Group("/app")
	appGroup.GET("/", appHandler.HandleGetIndex)
	appGroup.GET("/search", searchHandler.HandleGetIndex)
	appGroup.POST("/search", searchHandler.HandlePostSearch)

	authGroup := app.Group("/auth")
	authGroup.GET("/login", authHandler.HandleGetLogin)
	authGroup.POST("/login", authHandler.HandlePostLogin)
	authGroup.GET("/signup", authHandler.HandleGetSignup)
	authGroup.POST("/signup", authHandler.HandlePostSignup)

	subGroup := appGroup.Group("/subscriptions")
	subGroup.GET("/", subHandler.HandleGetIndex)
	subGroup.POST("/", subHandler.HandlePostSubscribe)

	app.Logger.Fatal(app.Start(port))
}
