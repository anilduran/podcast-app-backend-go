package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                    uuid.UUID            `gorm:"type:uuid;primary_key;" json:"id"`
	Username              string               `json:"username"`
	Email                 string               `json:"email"`
	Password              string               `json:"password"`
	ProfilePhotoUrl       string               `json:"profile_photo_url"`
	PodcastLists          []PodcastList        `gorm:"foreignKey:CreatorID;" json:"podcast_lists"`
	Playlists             []Playlist           `gorm:"foreignKey:CreatorID;" json:"playlists"`
	PodcastListComments   []PodcastListComment `gorm:"foreignKey:UserID;" json:"podcast_list_comments"`
	PodcastComments       []PodcastComment     `gorm:"foreignKey:UserID;" json:"podcast_comments"`
	LikedPodcasts         []*Podcast           `gorm:"many2many:podcast_likes;" json:"liked_podcasts"`
	FollowingPodcastLists []*PodcastList       `gorm:"many2many:following_podcast_lists;" json:"following_podcast_lists"`
	CreatedAt             *time.Time           `json:"created_at"`
	UpdatedAt             *time.Time           `json:"updated_at"`
	DeletedAt             *time.Time           `json:"deleted_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
