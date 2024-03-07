package main

import (
	"database/sql"
	"flag"
	"net/http"
	"time"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	session "github.com/spazzymoto/echo-scs-session"
	"newser.app/infra/repository"
	"newser.app/server/handler"
	custommiddleware "newser.app/server/middleware"
	"newser.app/server/service"
)

// repositories
var (
	userRepo         repository.UserRepository
	subscriptionRepo repository.SubscriptionRepository
	newsfeedRepo     repository.NewsfeedRepository
	articleRepo      repository.ArticleRepository
	collectionRepo   repository.CollectionRepository
	personRepo       repository.PersonRepository
	imageRepo        repository.ImageRepository
)

// services
var (
	authService         service.AuthService
	api                 service.API
	newsfeedService     service.NewsfeedService
	subscriptionService service.SubscriptionService
)

func main() {
	addr := flag.String("addr", ":4321", "HTTP network address")
	dev := flag.Bool("dev", false, "Whether to run the server in development mode")
	dsn := flag.String("dsn", "data/newser.sqlite", "Sqlite data source name")
	flag.Parse()

	conf := custommiddleware.NewConfig(*dev, *dsn)

	app := echo.New()
	setLogLevel(app, *dev)
	sessionDb, err := openSessionDB(*dsn)
	if err != nil {
		app.Logger.Fatal(err.Error())
	}

	db, err := openDB(*dsn)
	if err != nil {
		app.Logger.Fatal(err.Error())
	}
	defer db.Close()

	sessionManager := initSessions(app, sessionDb)
	initHandlers(app, db, sessionManager, *dev)
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

func initHandlers(app *echo.Echo, db *sqlx.DB, sessionManager *scs.SessionManager, isDev bool) {

	userRepo = repository.NewUserSqliteRepo(db)
	newsfeedRepo = repository.NewNewsfeedSqliteRepo(db)
	articleRepo = repository.NewArticleSqliteRepo(db)
	subscriptionRepo = repository.NewSubscriptionSqliteRepo(db)
	collectionRepo = repository.NewCollectionSqliteRepo(db)
	personRepo = repository.NewPersonSqliteRepository(db)
	imageRepo = repository.NewImageSqliteRepo(db)

	userRepo.Migrate()
	newsfeedRepo.Migrate()
	articleRepo.Migrate()
	subscriptionRepo.Migrate()
	collectionRepo.Migrate()
	personRepo.Migrate()
	imageRepo.Migrate()

	authService = service.NewAuthService(userRepo, collectionRepo)
	api = service.NewAPI(&http.Client{})
	subscriptionService = service.NewSubscriptionService(subscriptionRepo, newsfeedRepo, articleRepo, collectionRepo)
	newsfeedService = service.NewNewsfeedService(articleRepo, imageRepo, personRepo, newsfeedRepo)

	homeHandler := handler.NewHomeHandler(sessionManager)
	authHandler := handler.NewAuthHandler(authService, sessionManager)
	deskHandler := handler.NewDeskHandler(api, subscriptionService, authService, newsfeedService, sessionManager)

	app.GET("/", homeHandler.Home)
	authGroup := app.Group("/auth")
	authGroup.GET("/signup", authHandler.GetSignup)
	authGroup.POST("/signup", authHandler.PostSignup)
	authGroup.GET("/login", authHandler.GetLogin)
	authGroup.POST("/login", authHandler.PostLogin)
	authGroup.POST("/logout", authHandler.PostLogout)

	deskGroup := app.Group("/desk")
	deskGroup.Use(custommiddleware.Auth(sessionManager))
	deskGroup.Use(custommiddleware.SidebarLinks(
		sessionManager,
		&subscriptionService,
		&authService,
	))
	deskGroup.GET("/", deskHandler.GetDeskIndex)
	deskGroup.GET("/search", deskHandler.GetDeskSearch)
	deskGroup.POST("/search", deskHandler.PostDeskSearch)
	deskGroup.POST("/subscribe", deskHandler.PostDeskSubscribe)
	deskGroup.GET("/articles/:articleid", deskHandler.GetDeskArticle)
	deskGroup.GET("/feeds/:feedid", deskHandler.GetDeskNewsfeed)
	deskGroup.GET("/collections/:collectionname", deskHandler.GetDeskCollection)
}

func openDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", dsn)
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

func openSessionDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
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

func initSessions(app *echo.Echo, db *sql.DB) *scs.SessionManager {
	app.Logger.Debug("Migrating sessions table...")
	sessionQ := `
	CREATE TABLE IF NOT EXISTS sessions(
		token TEXT PRIMARY KEY,
        data BLOB NOT NULL,
        expiry REAL NOT NULL
	);
	`
	_, err := db.Exec(sessionQ)
	if err != nil {
		app.Logger.Fatal("error migrating sessions table", err)
	} else {
		app.Logger.Debug("completed migrating sessions table")
	}
	sessionManager := scs.New()
	sessionManager.Lifetime = (7 * 24) * time.Hour

	sessionManager.Store = sqlite3store.New(db)
	return sessionManager
}
