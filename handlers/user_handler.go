package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/repository"
)

type UserHandler struct {
	ur *repository.UserRepository
}

func NewUserHandler(ur *repository.UserRepository) *UserHandler {
	return &UserHandler{
		ur: ur,
	}
}

func GetUsers(c *gin.Context) {

	message := "Hello from gin no users currently!"
	c.JSON(http.StatusOK, gin.H{"message": message})
}
