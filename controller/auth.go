package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const tokenString = "test_token"

func Auth(c *gin.Context) {
	authToken := c.Request.Header.Get("x-auth-token")
	if authToken == "" {
		c.JSON(http.StatusForbidden, gin.H{})
		c.Abort()
		return
	}

	if authToken != tokenString {
		c.JSON(http.StatusForbidden, gin.H{})
		c.Abort()
		return
	}
	c.Next()
}