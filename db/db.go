package db

import (
	"os"

	"example.com/podcast-app-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() (err error) {

	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Podcast{},
		&models.Playlist{},
		&models.PodcastComment{},
		&models.PodcastListComment{},
		&models.PodcastComment{},
	)

	return
}
