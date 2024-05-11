package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/model"
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

func (uh UserHandler) FetchUsers(ctx *gin.Context) {
	var us []model.User
	us, err := uh.ur.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": us})
}

func (uh UserHandler) CreateUser(ctx *gin.Context) {
	var u model.User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cu, err := uh.ur.Create(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": cu})
}
