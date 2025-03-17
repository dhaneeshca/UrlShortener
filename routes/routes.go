package routes

import (
	"net/http"

	"UrlShortener/handlers"
	"UrlShortener/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("URL Shortener API is running 🚀"))
	}).Methods("GET")

	r.HandleFunc("/shorten", handlers.ShortenURL(db)).Methods("POST")
	r.HandleFunc("/{short_url}", handlers.RedirectDB(db)).Methods("GET")
	r.HandleFunc("/api/auth/register", handlers.RegisterUser(db)).Methods("POST")

	// API Key Management
	r.HandleFunc("/api/apikey/generate", handlers.GenerateAPIKey(db)).Methods("POST")
	r.HandleFunc("/api/apikey/list", handlers.ListAPIKeys(db)).Methods("GET")
	r.HandleFunc("/api/apikey/revoke", handlers.RevokeAPIKey(db)).Methods("POST")

	// URL Shortening (With API Key Authentication)
	apiRoutes := r.PathPrefix("/api").Subrouter()
	apiRoutes.Use(middleware.ValidateAPIKey(db))

	apiRoutes.HandleFunc("/shorten", handlers.ShortenURL(db)).Methods("POST")
	apiRoutes.HandleFunc("/urls/{short_url}", handlers.RedirectDB(db)).Methods("GET")
	apiRoutes.HandleFunc("/urls/{short_url}", handlers.DeleteShortURL(db)).Methods("DELETE")

}
