package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PodcastList struct {
	ID              uuid.UUID   `gorm:"type:uuid;primary_key;" json:"id"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	ImageUrl        string      `json:"image_url"`
	Categories      []*Category `gorm:"many2many:podcast_list_categories;" json:"categories"`
	CreatorID       uuid.UUID   `json:"creator_id"`
	IsVisible       bool        `json:"is_visible" gorm:"default:true"`
	FollowedByUsers []*User     `gorm:"many2many:following_podcast_lists;" json:"followed_by_users"`
	CreatedAt       *time.Time  `json:"created_at"`
	UpdatedAt       *time.Time  `json:"updated_at"`
	DeletedAt       *time.Time  `json:"deleted_at"`
}

func (podcastList *PodcastList) BeforeCreate(tx *gorm.DB) (err error) {
	podcastList.ID = uuid.New()
	return
}
