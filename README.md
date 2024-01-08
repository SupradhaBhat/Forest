Simple HTTP server implemented in Golang

To use Gorilla Mux, you need to install it using the following command:
go get -u github.com/gorilla/mux

Execution instructions: 
go run filename.go

Visit http://localhost:8080/ and http://localhost:8080/user/John in your web browser to see the responses. The server will log information about each incoming request in the console.

http.HandleFunc is used to define a simple request handler function. The function takes two parameters: w (http.ResponseWriter) and r (http.Request). It writes a simple response to the client using fmt.Fprintf.

http.ListenAndServe starts the HTTP server on port 8080. If an error occurs, it is printed to the console.

Gorilla Mux is used for routing. Different routes are defined for the home page (/) and a user page (/user/{name}), where {name} is a variable extracted from the URL
