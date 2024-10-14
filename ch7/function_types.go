// Function types are a bridge to interfaces

// User-defined types allow functions to implement interfaces

package main

import (
	"fmt"
	"net/http"
)

type handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// f is the underlying function, and calling it with (w, r) executes the handler
	f(w, r)
}

// Example of a function that can be converted into a HandlerFunc
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Create a HandlerFunc from the HelloHandler function
	var h handler = HandlerFunc(HelloHandler)

	// Use it with http.Handle
	http.Handle("/", h)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
