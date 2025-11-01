// go build -o out && ./out
package main

import (
	"net/http"
	"//io"
)

func main() {
	//Create a new http.ServeMux
	mux := http.NewServeMux()
	//Use the http.NewServeMux's .Handle() method to add a handler for the root path (/).
	//Use a standard http.FileServer as the handler
	//Use http.Dir to convert a filepath (in our case a dot: . which indicates the current directory) to a directory for the http.FileServer.
	mux.Handle("/app/", http.StripPrefix("/app/",http.FileServer(http.Dir("."))));
	mux.HandleFunc("/healthz", readiness_endpoint)
	//Create a new http.Server struct.
	server := http.Server{}
	//Use the new "ServeMux" as the server's handler
	server.Handler = mux

	//Set the .Addr field to ":8080"
	server.Addr = ":8080"

	//Use the server's ListenAndServe method to start the server
	server.ListenAndServe()
}

func readiness_endpoint(w http.ResponseWriter, r *http.Request){
	//Write the Content-Type header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//Write the status code using w.WriteHeader
	w.WriteHeader(http.StatusOK)
	//Write the body text using w.Write
	//io.WriteString(w, "OK");
	w.Write([]byte(http.StatusText(http.StatusOK)))
}