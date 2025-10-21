// go build -o out && ./out
package main

import (
	"net/http"
)

func main() {
	//Create a new http.ServeMux
	mux := http.NewServeMux()
	//Use the http.NewServeMux's .Handle() method to add a handler for the root path (/).
	//Use a standard http.FileServer as the handler
	//Use http.Dir to convert a filepath (in our case a dot: . which indicates the current directory) to a directory for the http.FileServer.
	mux.Handle("/", http.FileServer(http.Dir(".")))

	//Create a new http.Server struct.
	server := http.Server{}
	//Use the new "ServeMux" as the server's handler
	server.Handler = mux

	//Set the .Addr field to ":8080"
	server.Addr = ":8080"

	//Use the server's ListenAndServe method to start the server
	server.ListenAndServe()
}
