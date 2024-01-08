package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a request handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Handle the incoming request
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	// Start the HTTP server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
