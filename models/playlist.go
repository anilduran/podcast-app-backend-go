package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Playlist struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ImageUrl    string     `json:"image_url"`
	Podcasts    []*Podcast `gorm:"many2many:playlist_podcasts;"`
	Creator     User
	CreatorID   uuid.UUID
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (playlist *Playlist) BeforeCreate(tx *gorm.DB) (err error) {
	playlist.ID = uuid.New()
	return
}
