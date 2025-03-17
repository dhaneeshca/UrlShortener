package routes

import (
	"UrlShortener/models"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ShortenRequest struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url,omitempty"` // Optional custom short URL
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortCode := make([]byte, 6)
	for i := range shortCode {
		shortCode[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortCode)
}

// Handler for shortening URLs
func ShortenURL(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ShortenRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.LongURL == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Use provided short URL or generate one
		shortCode := req.ShortURL
		if shortCode == "" {
			shortCode = generateShortCode()
		}

		// Check if the short URL already exists
		existingURL := models.URL{}
		if err := db.Where("short_url = ?", shortCode).First(&existingURL).Error; err == nil {
			http.Error(w, "Short URL already taken", http.StatusConflict)
			return
		} else if err != gorm.ErrRecordNotFound {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		// Save to database
		url := models.URL{LongURL: req.LongURL, ShortURL: shortCode}
		if err := db.Create(&url).Error; err != nil {
			http.Error(w, "Failed to save URL", http.StatusInternalServerError)
			return
		}

		// Response
		resp := ShortenResponse{ShortURL: "http://localhost:8080/" + shortCode}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func RedirectDB(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		arguments := mux.Vars(r)
		short_code, exists := arguments["short_url"]
		if !exists {
			http.Error(w, "short URL is mandatory", http.StatusBadRequest)
			return
		}

		var url models.URL
		if err := db.Where("Short_url = ?", short_code).First(&url).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, "Invalid URL", http.StatusNotFound)
			} else {
				http.Error(w, "Database error", http.StatusInternalServerError)
			}
			return
		}

		http.Redirect(w, r, url.LongURL, http.StatusFound)
	}
}
