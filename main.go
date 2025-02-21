package main

import (
	"log"
	"net/http"
)

// define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
// check if the current request URL path exactly matches "/". if it doesn't, use
// the http.NotFound() function to send a 404 response to the client.
// importantly, we then return from the handler. if we don't return the handler
// would keep executing and also write the "Hello from Snippetbox" message.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// add a showSnippet handler function.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// use r.Method to check whether the request is using POST or not. note that
	// http.MethodPost is a constant equal to the string "POST".
	if r.Method != http.MethodPost {
		// use the Header().Set() method to add an 'Allow: POST' header to the
		// responde header map. the first parameter is the header name, and
		// the second parameter is the header value
		// if it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. we then return from the function so that the
		// subsequent code is not executed.
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		// use the http.Error() function to send a 405 status code and "Method Not
		// Allowed" string as the response body.
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	// register the two new handler functions and corresponding URL patterns with
	// the servemux, in exactly the same way that we did before.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by the http.ListenAndServe() is always non-nil.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
