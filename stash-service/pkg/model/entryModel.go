package model

import (
	"time"

	"gorm.io/gorm"
	"joicejoseph.dev/clutter-cache/stash-service/internal/data"
)

// Entry represents a single entry in the app, which could be a link, text, image, or other content type.
type Entry struct {
	gorm.Model 
	User        string         `json:"user"`
	EntryType   data.EntryType `sql:"type:enum('Image', 'Link', 'Text');default:'Text'" json:"entryType"`      
	Content     string         `json:"content"`
	Description string         `json:"description,omitempty"`
	Tags        string         `json:"tags,omitempty"` 
	IsFavorite  bool           `json:"isFavorite"`
	Reminder    *time.Time     `json:"reminder,omitempty"`	
}