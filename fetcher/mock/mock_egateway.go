package mocks

import (
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/model"
	"github.com/stretchr/testify/mock"
)

type MockEthereumGateway struct {
	mock.Mock
}

func (m *MockEthereumGateway) GetByTransactionHash(hash string) (*model.Transaction, error) {
	args := m.Called(hash)
	return args.Get(0).(*model.Transaction), args.Error(1)
}
