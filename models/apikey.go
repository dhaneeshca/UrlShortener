package models

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type APIKey struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Key       string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Revoked   bool      `gorm:"default:false"`
}

// Generate a random API key
func GenerateRandomKey(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err) // Handle errors properly in production
	}
	return hex.EncodeToString(bytes)
}
