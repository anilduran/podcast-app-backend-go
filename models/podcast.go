package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Podcast struct {
	ID            uuid.UUID   `gorm:"type:uuid;primary_key;" json:"id"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	ImageUrl      string      `json:"image_url"`
	PodcastUrl    string      `json:"podcast_url"`
	Playlists     []*Playlist `gorm:"many2many:playlist_podcasts;" json:"playlists"`
	PodcastListID uuid.UUID   `json:"podcast_list_id"`
	PodcastList   PodcastList `gorm:"foreignKey:PodcastListID;" json:"podcast_lists"`
	LikedByUsers  []*User     `gorm:"many2many:podcast_likes;"`
	IsVisible     bool        `json:"is_visible" gorm:"default:true"`
	CreatedAt     *time.Time  `json:"created_at"`
	UpdatedAt     *time.Time  `json:"updated_at"`
	DeletedAt     *time.Time  `json:"deleted_at"`
}

func (podcast *Podcast) BeforeCreate(tx *gorm.DB) (err error) {
	podcast.ID = uuid.New()
	return
}
