package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/wenjun99/be/internal/handlers"
)

func Cors(next http.Handler) http.Handler {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("DOMAIN")},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodDelete,
			http.MethodPatch,
			http.MethodPatch,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	return cors.Handler(next)
}

func Routes() *chi.Mux {
	r := chi.NewRouter()

	// Setting up the middleware
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(Cors)

	// Setting up the routes
	r.Route("/event", func(r chi.Router) {
		r.Get("/", handlers.App.Test)
	})

	return r
}
