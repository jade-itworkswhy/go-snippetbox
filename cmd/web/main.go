package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	// Define a new command-line flag with the name 'addr', a default value of ":4000" // and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are // encountered during parsing the application will be terminated.
	flag.Parse()

	// note: alternatively, we can use os.Getenv("ENV_NAME")
	// BUTYou can’t specify a
	// default setting (the return value from os.Getenv() is the empty string if the environment
	// variable doesn’t exist), you don’t get the -help functionality that you do with commandline
	// flags, and the return value from os.Getenv() is always a string — you don’t get
	// automatic type conversions like you do with flag.Int(), flag.Bool() and the other
	// command line flag functions.

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux() // this makes a function statisfies handler interface
	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// Register the other application routes as normal.
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)

}
