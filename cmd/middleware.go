package main

import (
	"net/http"
	"os"

	"github.com/rs/cors"
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

