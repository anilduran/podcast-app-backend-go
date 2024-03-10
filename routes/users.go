package routes

import (
	"net/http"

	"example.com/podcast-app-go/db"
	"example.com/podcast-app-go/models"
	"example.com/podcast-app-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserInput struct {
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ProfilePhotoUrl string `json:"profile_photo_url"`
}

type UpdateUserInput struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ProfilePhotoUrl string `json:"profile_photo_url"`
}

func GetUsers(c *gin.Context) {

	var users []models.User

	result := db.DB.Find(&users)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})

}

func GetUserByID(c *gin.Context) {

	var user models.User

	id, _ := uuid.Parse(c.Param("id"))

	result := db.DB.First(&user, id)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {

	var input CreateUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.Where("username = ? OR email = ?", input.Username, input.Email).First(&user)

	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exists!",
		})
		return
	}

	user = models.User{
		Username:        input.Username,
		Email:           input.Email,
		Password:        utils.HashPassword(input.Password),
		ProfilePhotoUrl: input.ProfilePhotoUrl,
	}

	result = db.DB.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, user)

}

func UpdateUser(c *gin.Context) {

	var input UpdateUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.Status(http.StatusBadGateway)
		return
	}

	id, _ := uuid.Parse(c.Param("id"))

	var user models.User

	result := db.DB.Where("username = ? OR email = ? AND id = ?", input.Username, input.Email, id).First(&user)

	if result.Error == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result = db.DB.First(&user, id)

	if result.Error != nil {
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

	if input.ProfilePhotoUrl != "" {
		user.ProfilePhotoUrl = input.ProfilePhotoUrl
	}

	result = db.DB.Save(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func DeleteUser(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("id"))

	var user models.User

	result := db.DB.First(&user, id)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result = db.DB.Delete(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetPodcastListsByUserID(c *gin.Context) {

	userId, _ := uuid.Parse(c.Param("id"))

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
