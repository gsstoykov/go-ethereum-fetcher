package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/model"
	"github.com/gsstoykov/go-ethereum-fetcher/repository"
)

type TransactionHandler struct {
	tr repository.ITransactionRepository
}

func NewTransactionHandler(tr repository.ITransactionRepository) *TransactionHandler {
	return &TransactionHandler{
		tr: tr,
	}
}

func (th TransactionHandler) FetchTransactions(ctx *gin.Context) {
	var ts []model.Transaction
	ts, err := th.tr.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"transactions": ts})
}

func (th TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var t model.Transaction
	if err := ctx.ShouldBindJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ct, err := th.tr.Create(&t)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"transaction": ct})
}
