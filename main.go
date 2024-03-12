package main

import (
	"log"
	"net/http"
)

// init home handler
// "http.ResponseWriter" -> provides methods for assembling a HTTP response and sending it to the user
// "*http.Request" -> pointer to a struct which holds info about the current request(context?)
func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// initalise new servemux
	mux := http.NewServeMux()

	// register the home func as the handler for the root route
	mux.HandleFunc("/", home)

	log.Print("starting server on :4000")

	// start the server and listen
	err := http.ListenAndServe(":4000", mux)
	
	// use log.Fatal to log the error and exit.
	// any error from http.listenAndServe() is always non-nil.
	log.Fatal(err)

}
