package srv

import (
	"log"
	"net/http"
)

func addRoutes(mux *http.ServeMux, logger *log.Logger) {
	mux.Handle("POST /register", handleRegisterPost(logger))
	mux.Handle("POST /login", handleRegisterPost(logger))
	mux.Handle("GET /register", handleRegisterGet(logger))
	mux.Handle("GET /login", handleLoginGet(logger))
	mux.Handle("GET /user/{id}", handleUserGet(logger))
}

func NewServer(logger *log.Logger) http.Handler {
	var mux *http.ServeMux = http.NewServeMux()
	var handler http.Handler

	addRoutes(mux, logger)

	handler = mux
	return handler
}
