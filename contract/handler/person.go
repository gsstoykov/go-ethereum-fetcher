package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/contract/model"
	"github.com/gsstoykov/go-ethereum-fetcher/contract/repository"
)

type PersonHandler struct {
	pr repository.IPersonRepository
}

func NewPersonHandler(pr repository.IPersonRepository) *PersonHandler {
	return &PersonHandler{
		pr: pr,
	}
}

func (ph PersonHandler) SavePerson(ctx *gin.Context) {
	var p model.Person
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cp, err := ph.pr.Create(&p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"person": cp})
}

func (ph PersonHandler) ListPeople(ctx *gin.Context) {
	var ps []model.Person
	ps, err := ph.pr.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"people": ps})
}
