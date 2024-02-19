package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PodcastList struct {
	ID          uuid.UUID   `gorm:"type:uuid;primary_key;"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Categories  []*Category `gorm:"many2many:podcast_list_categories;"`
	CreatorID   uuid.UUID   `json:"creator_id"`
	CreatedAt   *time.Time  `json:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at"`
}

func (podcastList *PodcastList) BeforeCreate(tx *gorm.DB) (err error) {
	podcastList.ID = uuid.New()
	return
}
