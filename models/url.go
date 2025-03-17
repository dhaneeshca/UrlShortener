package models

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	ID           uint           `gorm:"primaryKey"`
	LongURL      string         `gorm:"type:text;not null"`
	ShortURL     string         `gorm:"type:varchar(50);unique;not null"`
	CustomDomain string         `gorm:"type:varchar(200)"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	ExpiryDate   time.Time      `gorm:"type:date"`
	DeletedAt    gorm.DeletedAt `gorm:"index"` // Soft delete support
}
