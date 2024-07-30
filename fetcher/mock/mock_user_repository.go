package mocks

import (
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/model"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of IUserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *model.User) (*model.User, error) {
	args := m.Called(user)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) Update(user *model.User) (*model.User, error) {
	args := m.Called(user)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) Delete(userId uint) error {
	args := m.Called(userId)
	return args.Error(0)
}

func (m *MockUserRepository) FindById(userId uint) (*model.User, error) {
	args := m.Called(userId)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) FindByUsername(username string) (*model.User, error) {
	args := m.Called(username)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) FindAll() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserRepository) FindUserTransactions(userID uint) ([]model.Transaction, error) {
	args := m.Called(userID)
	return args.Get(0).([]model.Transaction), args.Error(1)
}

func (m *MockUserRepository) AddTransactionToUser(user *model.User, tx *model.Transaction) error {
	args := m.Called(user, tx)
	return args.Error(0)
}
