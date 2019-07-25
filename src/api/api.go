package api

// All REST API resources.

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	rest "github.com/arush15june/revshellz/src/api/rest"
)

var r *chi.Mux

// InitRouter initializes the router with required middleware and routes.
func InitRouter() *chi.Mux {
	r = chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/chans", rest.GetChansResource)

	return r
}

// InitRestApi initializes the router and starts listening on `port`.
func InitRestApi(port string) {
	InitRouter()
	port = ":" + port
	go http.ListenAndServe(port, r)
}
