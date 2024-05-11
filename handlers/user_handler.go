package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/repository"
)

type UserHandler struct {
	ur repository.IUserRepository
}

func NewUserHandler(ur repository.IUserRepository) *UserHandler {
	return &UserHandler{
		ur: ur,
	}
}

func (uh UserHandler) GetUsers(c *gin.Context) {
	message := "Hello from gin no users currently!"
	c.JSON(http.StatusOK, gin.H{"message": message})
}
