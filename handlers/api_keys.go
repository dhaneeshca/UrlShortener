package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"UrlShortener/models"

	"gorm.io/gorm"
)

// Generate a random API Key
func generateAPIKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	apiKey := make([]byte, 32)
	for i := range apiKey {
		apiKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(apiKey)
}

// Generate API Key
func GenerateAPIKey(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserID uint `json:"user_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.UserID == 0 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		apiKey := models.APIKey{
			UserID: req.UserID,
			Key:    generateAPIKey(),
		}

		if err := db.Create(&apiKey).Error; err != nil {
			http.Error(w, "Failed to generate API key", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"api_key": apiKey.Key})
	}
}

// List API Keys for a User
func ListAPIKeys(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		var apiKeys []models.APIKey
		if err := db.Where("user_id = ? AND revoked = ?", userID, false).Find(&apiKeys).Error; err != nil {
			http.Error(w, "Failed to retrieve API keys", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(apiKeys)
	}
}

// Revoke API Key
func RevokeAPIKey(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserID uint   `json:"user_id"`
			APIKey string `json:"api_key"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.UserID == 0 || req.APIKey == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if err := db.Model(&models.APIKey{}).Where("user_id = ? AND key = ?", req.UserID, req.APIKey).
			Update("revoked", true).Error; err != nil {
			http.Error(w, "Failed to revoke API key", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"message": "API key revoked"})
	}
}
