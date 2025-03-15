package models

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	ID        uint           `gorm:"primaryKey"`
	LongURL   string         `gorm:"type:text;not null"`
	ShortURL  string         `gorm:"type:varchar(50);unique;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete support
}
