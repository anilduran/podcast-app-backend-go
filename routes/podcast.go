package routes

import (
	"net/http"

	"example.com/podcast-app-go/db"
	"example.com/podcast-app-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreatePodcastInput struct {
	Name          string    `form:"name" binding:"required"`
	Description   string    `form:"description" binding:"required"`
	ImageUrl      string    `form:"image_url" binding:"required"`
	PodcastUrl    string    `form:"podcast_url" binding:"required"`
	PodcastListID uuid.UUID `form:"podcast_list_id" binding:"required"`
	IsVisible     bool      `form:"is_visible" binding:"required"`
}

type UpdatePodcastInput struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	ImageUrl    string `form:"image_url"`
	PodcastUrl  string `form:"podcast_url"`
	IsVisible   bool   `form:"is_visible"`
}

type CreatePodcastCommentInput struct {
	Content string `form:"content" binding:"required"`
}

type UpdatePodcastCommentInput struct {
	Content string `form:"content"`
}

func GetPodcasts(c *gin.Context) {

	var podcasts []models.Podcast

	result := db.DB.Find(&podcasts)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": podcasts,
	})

}

func GetPodcastByID(c *gin.Context) {
	var podcast models.Podcast

	id, _ := uuid.Parse(c.Param("id"))

	result := db.DB.First(&podcast, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func CreatePodcast(c *gin.Context) {

	var input CreatePodcastInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	podcast := models.Podcast{
		Name:          input.Name,
		Description:   input.Description,
		ImageUrl:      input.ImageUrl,
		PodcastUrl:    input.PodcastUrl,
		PodcastListID: input.PodcastListID,
		IsVisible:     input.IsVisible,
	}

	result := db.DB.Create(&podcast)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, podcast)

}

func UpdatePodcast(c *gin.Context) {

	var input models.Podcast

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := uuid.Parse(c.GetString("userId"))

	var podcast models.Podcast

	result := db.DB.First(&podcast, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Name != "" {
		podcast.Name = input.Name
	}

	if input.Description != "" {
		podcast.Description = input.Description
	}

	if input.ImageUrl != "" {
		podcast.ImageUrl = input.ImageUrl
	}

	if input.PodcastUrl != "" {
		podcast.PodcastUrl = input.PodcastUrl
	}

	if input.IsVisible != podcast.IsVisible {
		podcast.IsVisible = !podcast.IsVisible
	}

	result = db.DB.Save(&podcast)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, podcast)

}

func DeletePodcast(c *gin.Context) {

	id, _ := uuid.Parse(c.GetString("userId"))

	var podcast models.Podcast

	result := db.DB.First(&podcast, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&podcast)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, podcast)

}

func GetPodcastComments(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("id"))

	var comments []models.PodcastComment

	result := db.DB.Where("podcast_id = ?", id).Find(&comments)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})

}

func CreatePodcastComment(c *gin.Context) {

	var input CreatePodcastCommentInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := uuid.Parse(c.Param("id"))

	userId, _ := uuid.Parse(c.GetString("userId"))

	podcastComment := models.PodcastComment{
		Content:   input.Content,
		PodcastID: id,
		UserID:    userId,
	}

	result := db.DB.Create(&podcastComment)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, podcastComment)
}

func UpdatePodcastComment(c *gin.Context) {

	var input UpdatePodcastCommentInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := uuid.Parse(c.Param("commentId"))

	var podcastComment models.PodcastComment

	result := db.DB.First(&podcastComment, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	podcastComment.Content = input.Content

	result = db.DB.Save(&podcastComment)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, podcastComment)

}

func DeletePodcastComment(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("commentId"))

	var podcastComment models.PodcastComment

	result := db.DB.First(&podcastComment, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&podcastComment)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, podcastComment)

}

func LikePodcast(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	err := db.DB.First(&user, userId).Error

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := uuid.Parse(c.Param("id"))

	var podcast models.Podcast

	err = db.DB.First(&podcast, id).Error

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = db.DB.Model(&user).Association("Podcasts").Append(&podcast)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, podcast)
}

func UnlikePodcast(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	err := db.DB.First(&user, userId).Error

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := uuid.Parse(c.Param("id"))

	var podcast models.Podcast

	err = db.DB.First(&podcast, id).Error

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = db.DB.Model(&user).Association("Podcasts").Delete(&podcast)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, podcast)

}
