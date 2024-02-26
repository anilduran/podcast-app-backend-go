package routes

import (
	"net/http"

	"example.com/podcast-app-go/db"
	"example.com/podcast-app-go/models"
	"example.com/podcast-app-go/utils"
	"github.com/gin-gonic/gin"
)

type SignUpInput struct {
	Username string `form:"username" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type SignInInput struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func SignUp(c *gin.Context) {

	var input SignUpInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.Where("username = ? OR email = ?", input.Username, input.Email).Find(&user)

	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exits",
		})
		return
	}

	user = models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: utils.HashPassword(input.Password),
	}

	result = db.DB.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}

func SignIn(c *gin.Context) {

	var input SignInInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.Where("email = ?", input.Email).First(&user)

	if result.Error == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid credentials!",
		})
		return
	}

	comparePasswords := utils.ComparePasswords(user.Password, input.Password)

	if !comparePasswords {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials!",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
