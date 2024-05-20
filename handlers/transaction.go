package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	egateway "github.com/gsstoykov/go-ethereum-fetcher/ethereum"
	"github.com/gsstoykov/go-ethereum-fetcher/model"
	"github.com/gsstoykov/go-ethereum-fetcher/repository"
)

type TransactionHandler struct {
	tr repository.ITransactionRepository
	eg egateway.IEthereumGateway
}

func NewTransactionHandler(tr repository.ITransactionRepository, eg egateway.IEthereumGateway) *TransactionHandler {
	return &TransactionHandler{
		tr: tr,
		eg: eg,
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

func (th TransactionHandler) FetchTransactionsList(ctx *gin.Context) {
	txHash := ctx.Param("hash")
	tx, err := th.tr.FindByTransactionHash(txHash)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx == nil {
		tx, err = th.eg.GetByTransactionHash(txHash)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tx, err = th.tr.Create(tx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}
	ctx.JSON(http.StatusOK, gin.H{"transaction": tx})
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
