package middleware

import (
	"net/http"
	"strings"

	"UrlShortener/models"

	"gorm.io/gorm"
)

// ValidateAPIKey checks if a valid API key is provided in the request header
func ValidateAPIKey(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Try getting API key from "X-API-Key" header
			apiKey := r.Header.Get("X-API-Key")

			// If empty, check for "Authorization: Bearer <API_KEY>"
			if apiKey == "" {
				authHeader := r.Header.Get("Authorization")
				if strings.HasPrefix(authHeader, "Bearer ") {
					apiKey = strings.TrimPrefix(authHeader, "Bearer ")
				}
			}

			// If API key is still missing, reject the request
			if apiKey == "" {
				http.Error(w, "API key is required", http.StatusUnauthorized)
				return
			}

			// Validate API key in the database
			var key models.APIKey
			if err := db.Where("key = ? AND revoked = false", apiKey).First(&key).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					http.Error(w, "Invalid API key", http.StatusForbidden)
				} else {
					http.Error(w, "Database error", http.StatusInternalServerError)
				}
				return
			}

			// If everything is fine, pass control to the next handler
			next.ServeHTTP(w, r)
		})
	}
}
