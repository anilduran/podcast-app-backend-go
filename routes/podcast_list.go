package routes

import (
	"net/http"

	"example.com/podcast-app-go/db"
	"example.com/podcast-app-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreatePodcastListInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageUrl    string `json:"image_url" binding:"required"`
	IsVisible   bool   `json:"is_visible" binding:"required"`
}

type UpdatePodcastListInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	IsVisible   bool   `json:"is_visible"`
}

type CreatePodcastListCommentInput struct {
	Content string `json:"content" binding:"required"`
}

type UpdatePodcastListCommentInput struct {
	Content string `json:"content"`
}

func GetPodcastLists(c *gin.Context) {

	var podcastLists []models.PodcastList

	result := db.DB.Find(&podcastLists)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": podcastLists,
	})

}

func GetPodcastListByID(c *gin.Context) {

	var podcastList models.PodcastList

	id, _ := uuid.Parse(c.Param("id"))

	result := db.DB.First(&podcastList, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, podcastList)

}

func CreatePodcastList(c *gin.Context) {

	var input CreatePodcastListInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userId, _ := uuid.Parse(c.GetString("userId"))

	podcastList := models.PodcastList{
		Name:        input.Name,
		Description: input.Description,
		ImageUrl:    input.ImageUrl,
		IsVisible:   input.IsVisible,
		CreatorID:   userId,
	}

	result := db.DB.Create(&podcastList)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, podcastList)

}

func UpdatePodcastList(c *gin.Context) {

	var input UpdatePodcastListInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := uuid.Parse(c.Param("id"))

	var podcastList models.PodcastList

	result := db.DB.First(&podcastList, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Name != "" {
		podcastList.Name = input.Name
	}

	if input.Description != "" {
		podcastList.Description = input.Description
	}

	if input.ImageUrl != "" {
		podcastList.ImageUrl = input.ImageUrl
	}

	if input.IsVisible != podcastList.IsVisible {
		podcastList.IsVisible = !podcastList.IsVisible
	}

	result = db.DB.Save(&podcastList)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, podcastList)

}

func DeletePodcastList(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("id"))

	var podcastList models.PodcastList

	result := db.DB.First(&podcastList, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&podcastList)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, podcastList)

}

func GetPodcastListComments(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))

	var podcastListComments []models.PodcastListComment

	result := db.DB.Where("podcast_list_id = ?", id).Find(&podcastListComments)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": podcastListComments,
	})
}

func CreatePodcastListComment(c *gin.Context) {
	var input CreatePodcastListCommentInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	podcastListComment := models.PodcastListComment{
		Content: input.Content,
	}

	result := db.DB.Create(&podcastListComment)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, podcastListComment)

}

func UpdatePodcastListComment(c *gin.Context) {

	var input UpdatePodcastListCommentInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := uuid.Parse(c.Param("commentId"))

	var podcastListComment models.PodcastListComment

	result := db.DB.First(&podcastListComment, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	podcastListComment.Content = input.Content

	result = db.DB.Save(&podcastListComment)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, podcastListComment)

}

func DeletePodcastListComment(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("commentId"))

	var podcastListComment models.PodcastListComment

	result := db.DB.First(&podcastListComment, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&podcastListComment)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, podcastListComment)
}

func FollowPodcastList(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	err := db.DB.First(&user, userId).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	id, _ := uuid.Parse(c.Param("id"))

	var podcastList models.PodcastList

	err = db.DB.First(&podcastList, id).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err = db.DB.Model(&user).Association("FollowingPodcastLists").Append(&podcastList)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, podcastList)

}

func UnfollowPodcastList(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	err := db.DB.First(&user, userId).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	id, _ := uuid.Parse(c.Param("id"))

	var podcastList models.PodcastList

	err = db.DB.First(&podcastList, id).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err = db.DB.Model(&user).Association("FollowingPodcastLists").Delete(&podcastList)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, podcastList)

}

func GetPodcastsByPodcastListID(c *gin.Context) {

	var podcasts []models.Podcast

	id, _ := uuid.Parse(c.Param("id"))

	err := db.DB.Where("podcast_list_id = ?", id).Find(&podcasts).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": podcasts,
	})

}

func GetCreatorByPodcastListID(c *gin.Context) {

	var podcastList models.PodcastList

	id, _ := uuid.Parse(c.Param("id"))

	err := db.DB.First(&podcastList, id).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var user models.User

	err = db.DB.Select("username", "profile_photo_url").First(&user, podcastList.CreatorID).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func SearchPodcastLists(c *gin.Context) {

	query := c.Query("query")

	var podcastLists []models.PodcastList

	err := db.DB.Where("name LIKE ?", "%"+query+"%").Find(podcastLists).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": podcastLists,
	})

}
