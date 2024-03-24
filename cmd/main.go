package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
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
	"newser.app/internal/usecase/collection"
	"newser.app/internal/usecase/session"

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
	collectionRepo repository.CollectionRepository
	// personRepo       repository.PersonRepository
	// imageRepo        repository.ImageRepository
)

// services
var (
	authService         auth.AuthService
	sessionService      session.SessionService
	subscriptionService subscription.SubscriptionService
	// newsfeedService     newsfeed.NewsfeedService
	collectionService collection.CollectionService
	discoveryService  discovery.DiscoveryService
	addr              string
	cookieDomain      string
	dsn               string
)

// sessionManager
var sessionManager *scs.SessionManager

func main() {
	// addr := flag.String("addr", ":4321", "HTTP network address")
	// dev := flag.Bool("dev", false, "Whether to run the server in development mode")
	// dsn := flag.String("dsn", "data/newser.sqlite", "Sqlite data source name")
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	mode := os.Getenv("MODE")
	isDev := mode == "dev"

	if isDev {
		addr = ":" + os.Getenv("DEV_ADDR")
		cookieDomain = os.Getenv("DEV_COOKIE_DOMAIN")
		dsn = os.Getenv("DEV_DSN")
	} else {
		addr = ":" + os.Getenv("PROD_ADDR")
		cookieDomain = os.Getenv("PROD_COOKIE_DOMAIN")
		dsn = os.Getenv("PROD_DSN")
	}
	fmt.Println("addr:", addr)
	fmt.Println("cookieDomain:", cookieDomain)
	fmt.Println("dsn:", dsn)
	// flag.Parse()

	conf := custommiddleware.NewConfig(isDev)

	app := echo.New()
	setLogLevel(app, isDev)
	sessionDb, err := openSessionDB(dsn)
	if err != nil {
		app.Logger.Fatal(err.Error())
	}

	db, err := openDB(dsn)
	if err != nil {
		app.Logger.Fatal(err.Error())
	}
	defer db.Close()

	sessionManager = initSessions(app, sessionDb)

	app.Static("/static", "view/static")

	app.Use(middleware.CSRFWithConfig(
		middleware.CSRFConfig{
			TokenLookup:    "cookie:_csrf",
			CookieName:     "_csrf",
			CookieSameSite: http.SameSiteStrictMode,
			CookieHTTPOnly: true,
			CookiePath:     "/",
		},
	))
	app.Use(echosession.LoadAndSave(sessionManager))
	app.Use(custommiddleware.ContextValue)
	app.Use(conf.SetConfig)

	initRepos(db)
	initServices()
	initApiHandlers(app)
	initWebHandlers(app)

	app.Logger.Fatal(app.Start(addr))
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
	app.Use(custommiddleware.LayoutPreference(sessionManager))
	app.Use(custommiddleware.VewPreference(sessionManager))
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
		collectionService,
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
	collectionHandler := webhandler.NewWebCollectionHandler(
		sessionService,
		collectionService,
		authService,
		subscriptionService,
	)

	noteHandler := webhandler.NewWebNoteHandler(sessionService)

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
	collectionHandler.Routes(
		app,
		custommiddleware.CtxFlash(sessionManager),
		custommiddleware.AuthContext(sessionManager),
		// custommiddleware.Auth(sessionManager),
		custommiddleware.HTMX,
	)
	noteHandler.Routes(
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
	collectionService = collection.NewCollectionService(collectionRepo)
	discoveryService = discovery.NewDiscoveryService(&http.Client{})
}

func initRepos(db *sqlx.DB) {
	authRepo = sqlite.NewAuthSqliteRepo(db)
	subscriptionRepo = sqlite.NewSubscriptionSqliteRepo(db)
	collectionRepo = sqlite.NewCollectionSqliteRepo(db)
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
	// sessionManager.Cookie.Domain = cookieDomain
	sessionManager.Cookie.SameSite = http.SameSiteStrictMode
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
