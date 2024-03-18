package main

import (
	"database/sql"
	"flag"
	"net/http"
	"time"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	echosession "github.com/spazzymoto/echo-scs-session"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"

	// "newser.app/infra/repository"
	// "newser.app/server/handler"
	"newser.app/internal/infra/repository"
	"newser.app/internal/infra/repository/sqlite"
	apihandler "newser.app/internal/server/api/handler"
	custommiddleware "newser.app/internal/server/middleware"
	webhandler "newser.app/internal/server/web/handler"
	"newser.app/internal/usecase/auth"
	"newser.app/internal/usecase/session"
)

// repositories
var (
	authRepo repository.AuthRepository
	// userRepo         repository.UserRepository
	// subscriptionRepo repository.SubscriptionRepository
	// newsfeedRepo     repository.NewsfeedRepository
	// articleRepo      repository.ArticleRepository
	// collectionRepo   repository.CollectionRepository
	// personRepo       repository.PersonRepository
	// imageRepo        repository.ImageRepository
)

// services
var (
	authService    auth.AuthService
	sessionService session.SessionService
	// api                 service.API
	// newsfeedService     service.NewsfeedService
	// subscriptionService service.SubscriptionService
	// collectionService   service.CollectionService
)

// sessionManager
var sessionManager *scs.SessionManager

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

	sessionManager = initSessions(app, sessionDb)

	app.Static("/static", "view/static")
	app.Use(middleware.CSRFWithConfig(
		middleware.CSRFConfig{
			TokenLookup:    "cookie:_csrf",
			CookiePath:     "/",
			CookieDomain:   "localhost",
			CookieSecure:   true,
			CookieHTTPOnly: true,
			CookieSameSite: http.SameSiteStrictMode,
		},
	))
	app.Use(echosession.LoadAndSave(sessionManager))
	app.Use(custommiddleware.ContextValue)
	app.Use(conf.SetConfig)

	initRepos(db)
	initServices()
	initApiHandlers(app)
	initWebHandlers(app)

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

func initApiHandlers(app *echo.Echo) {
	apiAuthHandler := apihandler.NewAuthApiHandler(authService)
	apiAuthHandler.Routes(app)
}

func initWebHandlers(app *echo.Echo) {
	webHomeHandler := webhandler.NewWebHomeHandler(sessionService)
	webAuthHandler := webhandler.NewAuthWebHandler(authService, sessionService)
	webAppHandler := webhandler.NewWebAppHandler(sessionService, authService)

	webHomeHandler.Routes(
		app,
		custommiddleware.AuthContext(sessionManager),
	)
	webAuthHandler.Routes(
		app,
		custommiddleware.CtxFlash(sessionManager),
		custommiddleware.AuthContext(sessionManager),
		custommiddleware.HTMX,
	)
	webAppHandler.Routes(
		app,
		custommiddleware.CtxFlash(sessionManager),
		custommiddleware.AuthContext(sessionManager),
		custommiddleware.HTMX,
	)
}

func initServices() {
	authService = auth.NewAuthService(authRepo)
	sessionService = session.NewSessionService(sessionManager)
}

func initRepos(db *sqlx.DB) {
	authRepo = sqlite.NewAuthSqliteRepo(db)
}

// func initHandlers(app *echo.Echo, db *sqlx.DB, sessionManager *scs.SessionManager, isDev bool) {
// 	userRepo = repository.NewUserSqliteRepo(db)
// 	newsfeedRepo = repository.NewNewsfeedSqliteRepo(db)
// 	articleRepo = repository.NewArticleSqliteRepo(db)
// 	subscriptionRepo = repository.NewSubscriptionSqliteRepo(db)
// 	collectionRepo = repository.NewCollectionSqliteRepo(db)
// 	personRepo = repository.NewPersonSqliteRepository(db)
// 	imageRepo = repository.NewImageSqliteRepo(db)

// 	userRepo.Migrate()
// 	newsfeedRepo.Migrate()
// 	articleRepo.Migrate()
// 	subscriptionRepo.Migrate()
// 	collectionRepo.Migrate()
// 	personRepo.Migrate()
// 	imageRepo.Migrate()

// 	authService = service.NewAuthService(userRepo, collectionRepo)
// 	api = service.NewAPI(&http.Client{})
// 	subscriptionService = service.NewSubscriptionService(subscriptionRepo, newsfeedRepo, articleRepo, collectionRepo)
// 	newsfeedService = service.NewNewsfeedService(articleRepo, imageRepo, personRepo, newsfeedRepo, collectionRepo)
// 	collectionService = service.NewCollectionService(collectionRepo, articleRepo)

// 	homeHandler := handler.NewHomeHandler(sessionManager)
// 	authHandler := handler.NewAuthHandler(authService, sessionManager)
// 	deskHandler := handler.NewDeskHandler(api, subscriptionService, authService, newsfeedService, collectionService, sessionManager)

// 	app.GET("/", homeHandler.Home)
// 	authGroup := app.Group("/auth")
// 	authGroup.GET("/signup", authHandler.GetSignup)
// 	authGroup.POST("/signup", authHandler.PostSignup)
// 	authGroup.GET("/login", authHandler.GetLogin)
// 	authGroup.POST("/login", authHandler.PostLogin)
// 	authGroup.POST("/logout", authHandler.PostLogout)

// 	deskGroup := app.Group("/desk")
// 	deskGroup.Use(custommiddleware.Auth(sessionManager))
// 	deskGroup.Use(custommiddleware.SidebarLinks(
// 		sessionManager,
// 		&subscriptionService,
// 		&authService,
// 	))
// 	deskGroup.Use(custommiddleware.CtxCardState(sessionManager))
// 	deskGroup.Use(custommiddleware.PageTitle(sessionManager))
// 	deskGroup.Use(custommiddleware.ViewPreference(sessionManager))
// 	deskGroup.Use(custommiddleware.ReadPreference(sessionManager))
// 	deskGroup.GET("/", deskHandler.GetDeskIndex)
// 	deskGroup.GET("/search", deskHandler.GetDeskSearch)
// 	deskGroup.POST("/search", deskHandler.PostDeskSearch)
// 	deskGroup.POST("/subscribe", deskHandler.PostDeskSubscribe)
// 	deskGroup.GET("/articles/:articleid", deskHandler.GetDeskArticle)
// 	deskGroup.GET("/articles/update/", deskHandler.GetDeskUpdateArticles)
// 	deskGroup.GET("/feeds/:feedid", deskHandler.GetDeskNewsfeed)
// 	deskGroup.GET("/collections/:collectionname", deskHandler.GetDeskCollection)
// 	deskGroup.GET("/notes", deskHandler.GetDeskNotes)
// 	deskGroup.GET("/control/unreadcount", deskHandler.GetDeskUnreadCount)
// 	deskGroup.GET("/control/pagetitle", deskHandler.GetDeskPageTitle)
// 	deskGroup.POST("/control/setview", deskHandler.PostDeskSetView)
// 	deskGroup.POST("/control/setreadview", deskHandler.PostDeskSetReadView)
// 	deskGroup.POST("/collections/read", deskHandler.PostDeskAddToRead)
// 	deskGroup.POST("/collections/unread", deskHandler.PostDeskAddToUnread)
// 	deskGroup.POST("/control/setcollapse", deskHandler.PostDeskCardCollapsed)
// }

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
