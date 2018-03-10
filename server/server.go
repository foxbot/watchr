package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Run the server
func Run() error {
	r := mux.NewRouter()
	r.HandleFunc("/api", serveWs)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./wwwroot")))

	server := http.Server{
		Addr:    os.Getenv("SERVER_ADDR"),
		Handler: handlers.LoggingHandler(os.Stdout, r),
	}
	return server.ListenAndServe()
}
