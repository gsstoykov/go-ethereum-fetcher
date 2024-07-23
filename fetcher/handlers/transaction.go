package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	egateway "github.com/gsstoykov/go-ethereum-fetcher/fetcher/ethereum"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/model"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/repository"
)

type TransactionHandler struct {
	tr repository.ITransactionRepository
	ur repository.IUserRepository
	eg egateway.IEthereumGateway
}

func NewTransactionHandler(tr repository.ITransactionRepository, ur repository.IUserRepository, eg egateway.IEthereumGateway) *TransactionHandler {
	return &TransactionHandler{
		tr: tr,
		ur: ur,
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
	// Get the list of transactionHashes
	txHashes := ctx.QueryArray("transactionHashes")

	username, _ := ctx.Get("username")

	user, err := th.ur.FindByUsername(fmt.Sprint(username))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var txs []model.Transaction

	// Fetch for all transactionHashes
	for _, txHash := range txHashes {
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
		if user != nil {
			err = th.ur.AddTransactionToUser(user, tx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		}
		txs = append(txs, *tx)
	}
	ctx.JSON(http.StatusOK, gin.H{"transactions": txs})
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
