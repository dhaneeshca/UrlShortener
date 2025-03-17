package handlers

import (
	"UrlShortener/models"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ShortenRequest struct {
	LongURL      string `json:"long_url"`
	ShortURL     string `json:"short_url,omitempty"`
	CustomDomain string `json:"custom_domain,omitempty"`
	ExpiryDays   int32  `json:"expiry_days,omitempty"`
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

const expiration_days_const = 1000

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

		//custom domain is for storage only as the redirectiopn will be done by the domain owner,
		// we will handle the redirection part alone
		customDomain := req.CustomDomain

		// add expiry date
		expiration_days := req.ExpiryDays
		if expiration_days == 0 {
			expiration_days = expiration_days_const
		}
		expiry_date := time.Now().AddDate(0, 0, int(expiration_days))

		// Check if the short URL already exists
		existingURL := models.URL{}
		if err := db.Where("short_url = ? ", shortCode).First(&existingURL).Error; err == nil {
			http.Error(w, "Short URL already taken", http.StatusConflict)
			return
		} else if err != gorm.ErrRecordNotFound {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		// Save to database
		url := models.URL{LongURL: req.LongURL, ShortURL: shortCode, ExpiryDate: expiry_date, CustomDomain: customDomain}
		if err := db.Create(&url).Error; err != nil {
			http.Error(w, "Failed to save URL", http.StatusInternalServerError)
			return
		}

		domain := "http://localhost:8080/"
		if customDomain != "" {
			domain = customDomain + "/"
		}

		// Response
		resp := ShortenResponse{ShortURL: domain + shortCode}
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
		if url.ExpiryDate.Before(time.Now()) {
			http.Error(w, "Url Has expired", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, url.LongURL, http.StatusFound)
	}
}
