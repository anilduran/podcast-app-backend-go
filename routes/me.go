package routes

import (
	"net/http"

	"example.com/podcast-app-go/db"
	"example.com/podcast-app-go/models"
	"example.com/podcast-app-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateMyCredentialsInput struct {
	Username        string `form:"username"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	ProfilePhotoUrl string `form:"profile_photo_url"`
}

func GetMyCredentials(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	err := db.DB.Omit("password").First(&user, userId).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func UpdateMyCredentials(c *gin.Context) {

	var input UpdateMyCredentialsInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	err = db.DB.Where("username = ? OR email = ?", input.Username, input.Email).First(&user).Error

	if err == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userId, _ := uuid.Parse(c.GetString("userId"))

	err = db.DB.First(&user, userId).Error

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if input.Username != "" {
		user.Username = input.Username
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Password != "" {
		user.Password = utils.HashPassword(input.Password)
	}

	err = db.DB.Save(&user).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetMyPodcastLists(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var podcastLists []models.PodcastList

	err := db.DB.Where("creator_id = ?", userId).Find(&podcastLists).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": podcastLists,
	})

}

func GetMyPlaylists(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var playlists []models.Playlist

	err := db.DB.Where("creator_id = ?", userId).Find(&playlists).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": playlists,
	})

}

func GetMyLikedPodcasts(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	err := db.DB.First(&user, userId).Error

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var podcasts []models.Podcast

	err = db.DB.Model(&user).Association("LikedPodcasts").Find(&podcasts)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": podcasts,
	})

}

func GetMyFollowingPodcastLists(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	err := db.DB.First(&user, userId).First(&user).Error

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var podcastLists []models.PodcastList

	err = db.DB.Model(&user).Association("FollowingPodcastLists").Find(&podcastLists)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": podcastLists,
	})

}

func GetMyPodcastListComments(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var podcastListComments []models.PodcastListComment

	err := db.DB.Where("user_id = ?", userId).Find(&podcastListComments).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": podcastListComments,
	})

}

func GetMyPodcastComments(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var podcastComments []models.PodcastComment

	err := db.DB.Where("user_id = ?", userId).Find(&podcastComments).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": podcastComments,
	})

}

func GetPresignedURLForImage(c *gin.Context) {

	id, _ := uuid.Parse(c.GetString("userId"))

	url, key, err := utils.PresignerInstace.PutObject(id.String(), "images")

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": url,
		"key": key,
	})

}

func GetPresignedURLForPodcast(c *gin.Context) {

	id, _ := uuid.Parse(c.GetString("userId"))

	url, key, err := utils.PresignerInstace.PutObject(id.String(), "podcasts")

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": url,
		"key": key,
	})

}
