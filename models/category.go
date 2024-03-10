package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	Name         string         `json:"name"`
	PodcastLists []*PodcastList `gorm:"many2many:podcast_list_categories;" json:"podcast_lists"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    *time.Time     `json:"deleted_at"`
}

func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	category.ID = uuid.New()
	return
}
