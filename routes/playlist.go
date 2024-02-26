package routes

import (
	"net/http"

	"example.com/podcast-app-go/db"
	"example.com/podcast-app-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreatePlaylistInput struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
	ImageUrl    string `form:"image_url" binding:"required"`
}

type UpdatePlaylistInput struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	ImageUrl    string `form:"image_url"`
}

func GetPlaylists(c *gin.Context) {

	var playlists []models.Playlist

	result := db.DB.Find(&playlists)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": playlists,
	})

}

func GetPlaylistByID(c *gin.Context) {

	var playlist models.Playlist

	id, _ := uuid.Parse(c.Param("id"))

	result := db.DB.First(&playlist, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, playlist)

}

func CreatePlaylist(c *gin.Context) {

	var input CreatePlaylistInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userId, _ := uuid.Parse(c.GetString("userId"))

	playlist := models.Playlist{
		Name:        input.Name,
		Description: input.Description,
		ImageUrl:    input.ImageUrl,
		CreatorID:   userId,
	}
	result := db.DB.Create(&playlist)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, playlist)

}

func UpdatePlaylist(c *gin.Context) {

	var input UpdatePlaylistInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var playlist models.Playlist

	id, _ := uuid.Parse(c.Param("id"))

	result := db.DB.First(&playlist, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Name != "" {
		playlist.Name = input.Name
	}

	if input.Description != "" {
		playlist.Description = input.Description
	}

	if input.ImageUrl != "" {
		playlist.ImageUrl = input.ImageUrl
	}

	result = db.DB.Save(&playlist)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, playlist)

}

func DeletePlaylist(c *gin.Context) {

	id, _ := uuid.Parse(c.GetString("userId"))

	var playlist models.Playlist

	result := db.DB.First(&playlist, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&playlist)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, playlist)

}

func GetPlaylistPodcasts(c *gin.Context) {

}

func AddPodcastToPlaylist(c *gin.Context) {

}

func RemovePodcastFromPlaylist(c *gin.Context) {

}

func GetPodcastsByPlaylistID(c *gin.Context) {

}
