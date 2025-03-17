package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("URL Shortener API is running 🚀"))
	}).Methods("GET")

	r.HandleFunc("/shorten", ShortenURL(db)).Methods("POST")
	r.HandleFunc("/{short_url}", RedirectDB(db)).Methods("GET")
}
