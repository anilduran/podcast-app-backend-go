package routes

import (
	"net/http"

	"example.com/podcast-app-go/db"
	"example.com/podcast-app-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateCategoryInput struct {
	Name string `form:"name" binding:"required"`
}

type UpdateCategoryInput struct {
	Name string `form:"name"`
}

func GetCategories(c *gin.Context) {

	var categories []models.Category

	result := db.DB.Find(&categories)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})

}

func GetCategoryByID(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("id"))

	var category models.Category

	result := db.DB.First(&category, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, category)

}

func CreateCategory(c *gin.Context) {

	var input CreateCategoryInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var category = models.Category{
		Name: input.Name,
	}

	result := db.DB.Create(&category)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, category)

}

func UpdateCategory(c *gin.Context) {

	var input UpdateCategoryInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var category models.Category

	id, _ := uuid.Parse(c.Param("id"))

	result := db.DB.First(&category, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Name != "" {
		category.Name = input.Name
	}

	result = db.DB.Save(&category)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, category)

}

func DeleteCategory(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("id"))

	var category models.Category

	result := db.DB.First(&category, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&category)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &category)

}

func GetPodcastListsByCategoryID(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("id"))

	var category models.Category

	result := db.DB.First(&category, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var podcastLists []models.PodcastList

	err := db.DB.Model(&category).Association("PodcastLists").Find(&podcastLists)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": podcastLists,
	})

}
