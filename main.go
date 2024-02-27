package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/alexedwards/scs/gormstore"
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
	authService service.AuthService
)

func main() {
	addr := flag.String("addr", ":4321", "HTTP network address")
	dev := flag.Bool("dev", false, "Whether to run the server in development mode")
	dsn := flag.String("dsn", "data/newser.sqlite", "Sqlite data source name")
	flag.Parse()

	conf := custommiddleware.NewConfig(*dev, *dsn)

	app := echo.New()
	setLogLevel(app, *dev)

	db := initDB(app, *dsn)
	sessionManager := initSessions(app, db)
	initHandlers(app, db, sessionManager)
	app.Static("/static", "view/static")
	app.Use(session.LoadAndSave(sessionManager))
	app.Use(custommiddleware.ContextValue)
	app.Use(conf.SetConfig)
	app.Use((custommiddleware.CtxFlash(sessionManager)))
	app.Use(custommiddleware.AuthContext(sessionManager))

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

func initHandlers(app *echo.Echo, db *gorm.DB, sessionManager *scs.SessionManager) {

	userRepo = repository.NewUserGormRepo(db)
	subscriptionRepo = repository.NewSubscriptionGormRepo(db)
	newsfeedRepo := repository.NewNewsfeedGormRepo(db)

	authService = service.NewAuthService(userRepo)
	api := service.NewAPI(&http.Client{})
	subscriptionService := service.NewSubscriptionService(subscriptionRepo, newsfeedRepo)

	homeHandler := handler.NewHomeHandler(sessionManager)
	authHandler := handler.NewAuthHandler(authService, sessionManager)
	deskHandler := handler.NewDeskHandler(api, subscriptionService, authService, sessionManager)

	app.GET("/", homeHandler.Home)
	authGroup := app.Group("/auth")
	authGroup.GET("/signup", authHandler.GetSignup)
	authGroup.POST("/signup", authHandler.PostSignup)
	authGroup.GET("/login", authHandler.GetLogin)
	authGroup.POST("/login", authHandler.PostLogin)
	authGroup.POST("/logout", authHandler.PostLogout)

	deskGroup := app.Group("/desk")
	deskGroup.GET("/", deskHandler.GetDeskIndex)
}

func initDB(app *echo.Echo, dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		app.Logger.Fatal(err)
	}

	db.AutoMigrate(&dao.UserGorm{})
	db.AutoMigrate(&dao.SubscriptionGorm{})
	db.AutoMigrate(&dao.NewsfeedGorm{})
	db.AutoMigrate(&dao.ArticleGorm{})
	return db
}

func initSessions(app *echo.Echo, db *gorm.DB) *scs.SessionManager {
	sessionManager := scs.New()
	sessionManager.Lifetime = (7 * 24) * time.Hour

	store, err := gormstore.New(db)
	if err != nil {
		app.Logger.Fatal(err)
	}
	sessionManager.Store = store
	return sessionManager
}
