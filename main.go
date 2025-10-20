//go build -o out && ./out
package main

import (
	"net/http"
)

func main() {
	//Create a new http.ServeMux
	mux := http.NewServeMux()

	//Create a new http.Server struct.
	server := http.Server{}

	//Use the new "ServeMux" as the server's handler
	server.Handler = mux

	//Set the .Addr field to ":8080"
	server.Addr = ":8080"

	//Use the server's ListenAndServe method to start the server
	server.ListenAndServe()

	//Build and run your server (e.g. go build -o out && ./out)

	//Open http://localhost:8080 in your browser. You should see a 404 error because we haven't connected any handler logic yet. Don't worry, that's what is expected for the tests to pass for now.
	//for {}
}

/*





 */
