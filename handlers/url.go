package handlers

import (
	"net/http"

	"UrlShortener/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Delete (soft delete) a short URL
func DeleteShortURL(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		shortCode, exists := vars["short_url"]
		if !exists {
			http.Error(w, "Short URL is required", http.StatusBadRequest)
			return
		}

		// Soft delete the short URL
		if err := db.Where("short_url = ?", shortCode).Delete(&models.URL{}).Error; err != nil {
			http.Error(w, "Failed to delete URL", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
