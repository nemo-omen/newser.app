package main

import (
	"database/sql"
	"flag"
	"net/http"
	"time"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	session "github.com/spazzymoto/echo-scs-session"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"newser.app/infra/dao"
	"newser.app/infra/repository"
	"newser.app/server/handler"
	custommiddleware "newser.app/server/middleware"
	"newser.app/server/service"
)

// repositories
var (
	userRepo         repository.UserRepository
	subscriptionRepo repository.SubscriptionRepository
)

// services
var (
	authService    service.AuthService
	sessionManager *scs.SessionManager
)

func main() {
	addr := flag.String("addr", ":4321", "HTTP network address")
	dev := flag.Bool("dev", false, "Whether to run the server in development mode")
	dsn := flag.String("dsn", "data/newser.sqlite", "Sqlite data source name")
	flag.Parse()

	conf := custommiddleware.NewConfig(*dev, *dsn)
	db, err := sql.Open("sqlite3", *dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sessionManager = scs.New()
	sessionManager.Store = sqlite3store.New(db)
	sessionManager.Lifetime = (7 * 24) * time.Hour

	app := echo.New()
	setLogLevel(app, *dev)
	app.Static("/static", "view/static")
	app.Use(conf.SetConfig)
	app.Use(custommiddleware.ContextValue)
	app.Use((custommiddleware.CtxFlash))
	app.Use(custommiddleware.AuthContext)
	initHandlers(app, *dsn, sessionManager)

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

func initHandlers(app *echo.Echo, dsn string, sessionManager *scs.SessionManager) {
	db := initDB(app, dsn)

	userRepo = repository.NewUserGormRepo(db)
	subscriptionRepo = repository.NewSubscriptionGormRepo(db)

	authService = service.NewAuthService(userRepo)
	api := service.NewAPI(&http.Client{})
	subscriptionService := service.NewSubscriptionService(subscriptionRepo)

	homeHandler := handler.HomeHandler{}
	authHandler := handler.NewAuthHandler(authService)
	deskHandler := handler.NewDeskHandler(api, subscriptionService, authService)

	app.GET("/", homeHandler.Home)
	authGroup := app.Group("/auth")
	authGroup.Use(session.LoadAndSave(sessionManager))
	authGroup.GET("/signup", authHandler.GetSignup)
	authGroup.POST("/signup", authHandler.PostSignup)
	authGroup.GET("/login", authHandler.GetLogin)
	authGroup.POST("/login", authHandler.PostLogin)

	deskGroup := app.Group("/desk")
	deskGroup.Use(session.LoadAndSave(sessionManager))
	deskGroup.GET("/", deskHandler.GetDeskIndex)
}

func initDB(app *echo.Echo, dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		app.Logger.Fatal(err)
	}

	db.AutoMigrate(&dao.UserGorm{})
	return db
}
