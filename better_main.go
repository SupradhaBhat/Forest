package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define routes and handlers
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/user/{name}", userHandler).Methods("GET")

	// Use a middleware for logging requests
	router.Use(loggingMiddleware)

	// Serve static files from the "static" directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the HTTP server on port 8080
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the home route
	fmt.Fprintln(w, "Welcome to the home page!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Extract variables from the URL path using Gorilla Mux
	vars := mux.Vars(r)
	name := vars["name"]

	// Handle the user route
	fmt.Fprintf(w, "Hello, %s!", name)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log information about the request
		fmt.Printf("Request received: %s %s\n", r.Method, r.URL.Path)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
