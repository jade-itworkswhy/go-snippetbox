package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"jade-factory/go-snippetbox/internal/models"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"

	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql" // Import for side effect(usage of "_")
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include the structured logger, but we'll
// add more to this as the build progresses.
type application struct {
	logger         *slog.Logger
	snippets       models.SnippetModelInterface // Use our new interface type.
	users          models.UserModelInterface    // Use our new interface type.
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	// Define a new command-line flag with the name 'addr', a default value of ":4000" // and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are // encountered during parsing the application will be terminated.
	// Define a new command-line flag for the MySQL DSN string.

	// we have total control over which database is used at runtime, just by using the -dsn command-line flag.
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// note: alternatively, we can use os.Getenv("ENV_NAME")
	// BUTYou can’t specify a
	// default setting (the return value from os.Getenv() is the empty string if the environment
	// variable doesn’t exist), you don’t get the -help functionality that you do with commandline
	// flags, and the return value from os.Getenv() is always a string — you don’t get
	// automatic type conversions like you do with flag.Int(), flag.Bool() and the other
	// command line flag functions.

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// To keep the main() function tidy I've put the code for creating a connection
	// pool into the separate openDB() function below. We pass openDB() the DSN
	// from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	// close the connection pool later for graceful shutdown
	defer db.Close()

	// init template in-memory cache
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// init form decoder instance
	formDecoder := form.NewDecoder()

	// init session manager
	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour
	// enable https
	sessionManager.Cookie.Secure = true

	app := &application{
		logger:         logger,
		snippets:       &models.SnippetModel{DB: db},
		users:          &models.UserModel{DB: db},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	// init tls config for customisation
	tlsConfig := &tls.Config{
		// assembly implementations
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		// can be filtered
		// MinVersion: tls.VersionTLS12,
		// MaxVersion: tls.VersionTLS12,
	}

	// server with customizable struct
	srv := &http.Server{
		Addr:           *addr,
		MaxHeaderBytes: 524288, // Go **always** addes additional 4096 bytes of headroom.
		Handler:        app.routes(),
		ErrorLog:       slog.NewLogLogger(logger.Handler(), slog.LevelWarn),
		TLSConfig:      tlsConfig,
	}

	logger.Info("starting server", "addr", srv.Addr)

	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	// initialise the connection pool
	// actual connections to the database are established lazily
	// i.e once the first db connection is requested
	// that's why we check the connection with db.Ping() below
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// check the db connection
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	// successfully conneted
	return db, nil
}
