package routes

import (
	"irbistest/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(h *handlers.Handlers) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	// Static files
	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Routes
	r.Get("/", h.HomeHandler)
	r.Get("/take", h.TakeTokens)
	r.Get("/refresh", h.RefreshToken)
	return r
}
