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

	"newser.app/internal/usecase/newsfeed"
	"newser.app/internal/usecase/subscription"

	// "newser.app/internal/usecase/collection"
	"newser.app/internal/usecase/discovery"
)

// repositories
var (
	authRepo repository.AuthRepository
	// userRepo         repository.UserRepository
	subscriptionRepo repository.SubscriptionRepository
	// newsfeedRepo     repository.NewsfeedRepository
	// articleRepo      repository.ArticleRepository
	// collectionRepo   repository.CollectionRepository
	// personRepo       repository.PersonRepository
	// imageRepo        repository.ImageRepository
)

// services
var (
	authService         auth.AuthService
	sessionService      session.SessionService
	subscriptionService subscription.SubscriptionService
	newsfeedService     newsfeed.NewsfeedService
	// collectionService   service.CollectionService
	discoveryService discovery.DiscoveryService
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
	app.Use(custommiddleware.SetCurrentPath)
	app.Use(custommiddleware.PageTitle(sessionManager))
	app.Use(custommiddleware.SidebarLinks(
		sessionManager,
		&subscriptionService,
		&authService,
	),
	)
	homeHandler := webhandler.NewWebHomeHandler(sessionService)
	authHandler := webhandler.NewAuthWebHandler(
		authService,
		sessionService,
	)
	appHandler := webhandler.NewWebAppHandler(
		sessionService,
		authService,
		subscriptionService,
	)
	searchHandler := webhandler.NewWebSearchHandler(
		sessionService,
		discoveryService,
	)
	subscriptionHandler := webhandler.NewWebSubscriptionHandler(
		sessionService,
		subscriptionService,
		authService,
		discoveryService,
	)

	homeHandler.Routes(
		app,
		custommiddleware.AuthContext(sessionManager),
	)

	authHandler.Routes(
		app,
		custommiddleware.CtxFlash(sessionManager),
		custommiddleware.AuthContext(sessionManager),
		custommiddleware.HTMX,
	)

	appHandler.Routes(
		app,
		custommiddleware.CtxFlash(sessionManager),
		custommiddleware.AuthContext(sessionManager),
		// custommiddleware.Auth(sessionManager),
		custommiddleware.HTMX,
	)

	searchHandler.Routes(
		app,
		custommiddleware.CtxFlash(sessionManager),
		custommiddleware.AuthContext(sessionManager),
		// custommiddleware.Auth(sessionManager),
		custommiddleware.HTMX,
	)
	subscriptionHandler.Routes(
		app,
		custommiddleware.CtxFlash(sessionManager),
		custommiddleware.AuthContext(sessionManager),
		// custommiddleware.Auth(sessionManager),
		custommiddleware.HTMX,
	)
}

func initServices() {
	authService = auth.NewAuthService(authRepo)
	sessionService = session.NewSessionService(sessionManager)
	subscriptionService = subscription.NewSubscriptionService(subscriptionRepo)
	// newsfeedService = service.NewNewsfeedService(articleRepo, imageRepo, personRepo, newsfeedRepo, collectionRepo)
	// collectionService = service.NewCollectionService(collectionRepo, articleRepo)
	discoveryService = discovery.NewDiscoveryService(&http.Client{})
}

func initRepos(db *sqlx.DB) {
	authRepo = sqlite.NewAuthSqliteRepo(db)
	subscriptionRepo = sqlite.NewSubscriptionSqliteRepo(db)
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
