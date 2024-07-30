package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	mocks "github.com/gsstoykov/go-ethereum-fetcher/fetcher/mock"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchTransactions(t *testing.T) {
	mockTransactionRepo := new(mocks.MockTransactionRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockEthereumGateway := new(mocks.MockEthereumGateway)

	// Create a sample transaction
	sampleTransactions := []model.Transaction{
		{TransactionHash: "0x123"},
	}

	// Setup expectations
	mockTransactionRepo.On("FindAll").Return(sampleTransactions, nil)

	// Create the transaction handler
	th := NewTransactionHandler(mockTransactionRepo, mockUserRepo, mockEthereumGateway)

	// Setup the gin router
	router := gin.Default()
	router.GET("/transactions", th.FetchTransactions)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/transactions", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the results
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "0x123")

	mockTransactionRepo.AssertExpectations(t)
}

func TestFetchTransactionsList(t *testing.T) {
	mockTransactionRepo := new(mocks.MockTransactionRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockEthereumGateway := new(mocks.MockEthereumGateway)

	// Create a sample transaction
	sampleTransaction := model.Transaction{TransactionHash: "0x123"}

	// Define an error to return
	mockError := errors.New("transaction not found")

	// Setup expectations
	mockTransactionRepo.On("FindByTransactionHash", "0x123").Return(nil, mockError)
	mockEthereumGateway.On("GetByTransactionHash", "0x123").Return(&sampleTransaction, nil)
	mockTransactionRepo.On("Create", &sampleTransaction).Return(&sampleTransaction, nil)

	// Create the transaction handler
	th := NewTransactionHandler(mockTransactionRepo, mockUserRepo, mockEthereumGateway)

	// Setup the gin router
	router := gin.Default()
	router.GET("/transactions/list", th.FetchTransactionsList)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/transactions/list?transactionHashes=0x123", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the results
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "0x123")

	mockTransactionRepo.AssertExpectations(t)
}

func TestCreateTransaction(t *testing.T) {
	mockTransactionRepo := new(mocks.MockTransactionRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockEthereumGateway := new(mocks.MockEthereumGateway)

	// Create a sample transaction
	sampleTransaction := model.Transaction{TransactionHash: "0x123"}

	// Setup expectations
	mockTransactionRepo.On("Create", mock.AnythingOfType("*model.Transaction")).Return(&sampleTransaction, nil)

	// Create the transaction handler
	th := NewTransactionHandler(mockTransactionRepo, mockUserRepo, mockEthereumGateway)

	// Setup the gin router
	router := gin.Default()
	router.POST("/transactions", th.CreateTransaction)

	// Create a request to send to the above route
	transactionJSON, _ := json.Marshal(sampleTransaction)
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(transactionJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the results
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "0x123")

	mockTransactionRepo.AssertExpectations(t)
}
