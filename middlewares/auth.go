package middlewares

import (
	"net/http"

	"example.com/podcast-app-go/utils"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	token := c.Request.Header.Get("x-access-token")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized!",
		})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token is not valid!",
		})
		return
	}

	c.Set("userId", userId.String())

	c.Next()
}
