package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	egateway "github.com/gsstoykov/go-ethereum-fetcher/fetcher/ethereum"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/model"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/repository"
)

// TransactionHandler handles transaction-related requests.
type TransactionHandler struct {
	tr repository.ITransactionRepository
	ur repository.IUserRepository
	eg egateway.IEthereumGateway
}

// NewTransactionHandler creates a new TransactionHandler instance.
func NewTransactionHandler(tr repository.ITransactionRepository, ur repository.IUserRepository, eg egateway.IEthereumGateway) *TransactionHandler {
	return &TransactionHandler{
		tr: tr,
		ur: ur,
		eg: eg,
	}
}

// FetchTransactions handles fetching all transactions.
// It retrieves all transactions from the database using the transaction repository
// and returns them in the response.
func (th *TransactionHandler) FetchTransactions(ctx *gin.Context) {
	var ts []model.Transaction
	ts, err := th.tr.FindAll() // Retrieve all transactions from the repository
	if err != nil {
		log.Printf("Failed to fetch transactions: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"transactions": ts}) // Return the transactions in the response
}

// FetchTransactionsList handles fetching a list of transactions by their hashes.
// It first checks if transaction hashes were provided via the RLP middleware, otherwise,
// it falls back to query parameters. It retrieves transactions from the database or
// the Ethereum gateway if not found, and associates them with a user if provided.
func (th *TransactionHandler) FetchTransactionsList(ctx *gin.Context) {
	var txHashes []string

	// Check if transactionHashes were set via the RLP middleware
	if hashes, exists := ctx.Get("transactionHashes"); exists {
		var ok bool
		if txHashes, ok = hashes.([]string); !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid transaction hashes format in context"})
			return
		}
	} else {
		// Fallback to query parameters if no RLP data is provided
		txHashes = ctx.QueryArray("transactionHashes")
	}

	var user *model.User = nil
	username, exists := ctx.Get("username") // Retrieve the username from the context if set
	if exists {
		user, _ = th.ur.FindByUsername(fmt.Sprint(username)) // Fetch the user by username
	}

	var txs []model.Transaction
	for _, txHash := range txHashes {
		tx, err := th.tr.FindByTransactionHash(txHash) // Retrieve the transaction by hash from the repository
		if err != nil {
			log.Printf("Failed to find transaction by hash in db %s: %v\n", txHash, err)
			tx, err = th.eg.GetByTransactionHash(txHash) // Fetch the transaction from the Ethereum gateway if not found
			if err != nil {
				log.Printf("Failed to fetch transaction from Ethereum gateway: %v\n", err)
				continue
			}
			tx, err = th.tr.Create(tx) // Store the fetched transaction in the database
			if err != nil {
				log.Printf("Failed to store transaction: %v\n", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store transaction!"})
				return
			}
		}

		if user != nil {
			// Associate the transaction with the user if a user is provided
			err = th.ur.AddTransactionToUser(user, tx)
			if err != nil {
				log.Printf("Failed to associate transaction with user %s: %v\n", user.Username, err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to associate transaction with user!"})
				return
			}
		}
		txs = append(txs, *tx) // Add the transaction to the list to be returned
	}
	ctx.JSON(http.StatusOK, gin.H{"transactions": txs}) // Return the transactions in the response
}

// CreateTransaction handles the creation of a new transaction.
// It binds the JSON request data to the transaction model and stores it in the database.
func (th *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var t model.Transaction
	if err := ctx.ShouldBindJSON(&t); err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}

	ct, err := th.tr.Create(&t) // Create a new transaction in the repository
	if err != nil {
		log.Printf("Failed to store transaction: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store transaction!"})
		return
	}

	// Respond with the created transaction
	ctx.JSON(http.StatusCreated, gin.H{"transaction": ct})
}
