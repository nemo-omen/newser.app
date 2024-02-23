package main

import (
	"current/custommiddleware"
	"current/handler"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
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
	app := pocketbase.New()
	// app.Static("/public", "public")
	// app := echo.New()

	homeHandler := handler.HomeHandler{}
	appHandler := handler.AppHandler{}
	authHandler := handler.AuthHandler{}
	searchHandler := handler.SearchHandler{}
	subHandler := handler.SubscriptionHandler{}
	wsHandler := handler.WsHandler{}

	// Load middleware
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		r := e.Router
		r.GET("/public/**/*", apis.StaticDirectoryHandler(os.DirFS("./public"), false))

		r.Use(mode.SetMode)
		r.Use(custommiddleware.NewMiddlewareContextValue)
		r.Use(custommiddleware.CurrentPath)
		r.Use(custommiddleware.HTMX)
		// app.Use(middleware.Logger())
		r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:4321"},
		}))

		// r.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))))
		return nil
	})

	// routes
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		r := e.Router
		r.GET("/", homeHandler.HandleGetIndex)
		r.GET("/livereload", wsHandler.HandleWsConnect)

		appGroup := r.Group("/app")
		appGroup.GET("/", appHandler.HandleGetIndex)
		appGroup.GET("/search", searchHandler.HandleGetIndex)
		appGroup.POST("/search", searchHandler.HandlePostSearch)

		authGroup := r.Group("/auth")
		authGroup.GET("/login", authHandler.HandleGetLogin)
		authGroup.POST("/login", authHandler.HandlePostLogin)
		authGroup.GET("/signup", authHandler.HandleGetSignup)
		authGroup.POST("/signup", authHandler.HandlePostSignup)

		subGroup := appGroup.Group("/subscriptions")
		subGroup.GET("/", subHandler.HandleGetIndex)
		subGroup.POST("/", subHandler.HandlePostSubscribe)

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
	// app.Logger.Fatal(app.Start(port))
}
