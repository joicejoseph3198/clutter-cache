package model

import (
	"time"

	"gorm.io/gorm"
)

// Entry represents a single entry in the app, which could be a link, text, image, or other content type.
type Entry struct {
	gorm.Model 
	User        string        `json:"user"`
	EntryType   string        `json:"type"`          // Avoids using `Type` directly
	Title       string        `json:"title"`
	Content     string        `json:"content"`
	Description string        `json:"description,omitempty"`
	Tags        string        `json:"tags,omitempty"` // JSON column for storing tags as an array
	IsFavorite  bool          `json:"isFavorite"`
	Reminder    *time.Time    `json:"reminder,omitempty"`	
}