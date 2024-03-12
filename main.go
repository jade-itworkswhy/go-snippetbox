package main

import (
	"log"
	"net/http"
)

// init home handler
// "http.ResponseWriter" -> provides methods for assembling a HTTP response and sending it to the user
// "*http.Request" -> pointer to a struct which holds info about the current request(context?)
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	if r.Method != "POST" {
		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. We then return from the function so that the
		// subsequent code is not executed.

		// It’s only possible to call w.WriteHeader() once per response, and after the status code
		// has been written it can’t be changed
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// initalise new servemux
	mux := http.NewServeMux()

	// register the home func as the handler for the root route
	// !info Go's servemux treats the URL pattern "/" like a catch-all
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :4000")

	// start the server and listen
	/**
	In other Go projects or documentation you might sometimes see network addresses written
	using named ports like ":http" or ":http-alt" instead of a number. If you use a named
	port then the http.ListenAndServe() function will attempt to look up the relevant port
	number from your /etc/services file when starting the server, returning an error if a match
	can’t be found.
	*/

	/**
	If you pass nil as the second argument to http.ListenAndServe(), the server will use
	http.DefaultServeMux for routing.

	for the sake of clarity, maintainability and security, it’s generally a good idea to avoid
	http.DefaultServeMux and the corresponding helper functions. Use your own locallyscoped
	servemux instead
	*/

	err := http.ListenAndServe(":4000", mux)

	// use log.Fatal to log the error and exit.
	// any error from http.listenAndServe() is always non-nil.
	log.Fatal(err)

}
