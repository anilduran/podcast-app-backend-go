package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Podcast struct {
	ID          uuid.UUID   `gorm:"type:uuid;primary_key;"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	ImageUrl    string      `json:"image_url"`
	Playlists   []*Playlist `gorm:"many2many:playlist_podcast;"`
	CreatedAt   *time.Time  `json:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at"`
}

func (podcast *Podcast) BeforeCreate(tx *gorm.DB) (err error) {
	podcast.ID = uuid.New()
	return
}
