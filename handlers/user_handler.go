package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	message := "Hello from gin no users currently!"
	c.JSON(http.StatusOK, gin.H{"message": message})
}
