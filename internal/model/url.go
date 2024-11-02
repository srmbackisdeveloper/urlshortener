package model

import "time"

type URL struct {
	ID          int    `gorm:"primaryKey"`
	OriginalURL string `gorm:"not null"`
	ShortCode   string `gorm:"uniqueIndex;not null"`
	Clicks      int    `gorm:"default:0"`
	CreatedAt   time.Time
}