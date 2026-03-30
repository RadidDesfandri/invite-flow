package route

import (
	"net/http"

	"backend/handler"
)

type Router struct {
	mux      *http.ServeMux
	notFound http.HandlerFunc
}

func NewRouter(systemHandler *handler.SystemHandler, healthHandler *handler.HealthHandler, userHandler *handler.UserHandler) *Router {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", systemHandler.Root)
	mux.HandleFunc("GET /health", healthHandler.Check)
	mux.HandleFunc("GET /api/v1/ping", systemHandler.Ping)

	// Example user module routes
	mux.HandleFunc("POST /api/v1/users", userHandler.Create)
	mux.HandleFunc("GET /api/v1/users", userHandler.List)
	mux.HandleFunc("GET /api/v1/users/{id}", userHandler.GetByID)

	return &Router{
		mux:      mux,
		notFound: systemHandler.NotFound,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, pattern := r.mux.Handler(req)
	if pattern == "" || (pattern == "GET /" && req.URL.Path != "/") {
		r.notFound(w, req)
		return
	}
	handler.ServeHTTP(w, req)
}
