package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PodcastComment struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	Content   string     `json:"content"`
	UserID    uuid.UUID  `json:"user_id"`
	PodcastID uuid.UUID  `json:"podcast_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (podcastComment *PodcastComment) BeforeCreate(tx *gorm.DB) (err error) {
	podcastComment.ID = uuid.New()
	return
}
