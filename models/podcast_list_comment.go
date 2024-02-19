package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PodcastListComment struct {
	ID            uuid.UUID  `gorm:"type:uuid;primary_key;"`
	Content       string     `json:"content"`
	UserID        uuid.UUID  `json:"user_id"`
	PodcastListID uuid.UUID  `json:"podcast_list_id"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

func (podcastListComment *PodcastListComment) BeforeCreate(tx *gorm.DB) (err error) {
	podcastListComment.ID = uuid.New()
	return
}
