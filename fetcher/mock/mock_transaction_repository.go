package mocks

import (
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/model"
	"github.com/stretchr/testify/mock"
)

// MockTransactionRepository is a mock implementation of ITransactionRepository
type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(transaction *model.Transaction) (*model.Transaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) Update(transaction *model.Transaction) (*model.Transaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) Delete(transactionId uint) (*model.Transaction, error) {
	args := m.Called(transactionId)
	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) FindById(transactionId uint) (*model.Transaction, error) {
	args := m.Called(transactionId)
	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) FindByTransactionHash(transactionHash string) (*model.Transaction, error) {
	args := m.Called(transactionHash)
	// Ensure the return values are of the correct type
	var transaction *model.Transaction
	if value := args.Get(0); value != nil {
		transaction = value.(*model.Transaction)
	}
	return transaction, args.Error(1)
}

func (m *MockTransactionRepository) FindAll() ([]model.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]model.Transaction), args.Error(1)
}
