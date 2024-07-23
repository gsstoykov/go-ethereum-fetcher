package repository

import (
	"errors"
	"fmt"

	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/model"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	Create(transaction *model.Transaction) (*model.Transaction, error)
	Update(transaction *model.Transaction) (*model.Transaction, error)
	Delete(transactionId uint) (*model.Transaction, error)
	FindById(transactionId uint) (*model.Transaction, error)
	FindByTransactionHash(transactionHash string) (*model.Transaction, error)
	FindAll() ([]model.Transaction, error)
}

type TransactionRepository struct {
	Db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &TransactionRepository{Db: db}
}

func (r *TransactionRepository) Create(transaction *model.Transaction) (*model.Transaction, error) {
	if err := r.Db.Create(transaction).Error; err != nil {
		return nil, fmt.Errorf("could not create transaction: %w", err)
	}
	return transaction, nil
}

func (r *TransactionRepository) Update(transaction *model.Transaction) (*model.Transaction, error) {
	existingTransaction, err := r.FindById(transaction.ID)
	if err != nil {
		return nil, fmt.Errorf("could not find transaction: %w", err)
	}
	if existingTransaction == nil {
		return nil, errors.New("transaction not found")
	}

	// Update the fields of the existing transaction

	if err := r.Db.Save(existingTransaction).Error; err != nil {
		return nil, fmt.Errorf("could not update transaction: %w", err)
	}
	return existingTransaction, nil
}

func (r *TransactionRepository) Delete(transactionId uint) (*model.Transaction, error) {
	transaction, err := r.FindById(transactionId)
	if err != nil {
		return nil, fmt.Errorf("could not find transaction: %w", err)
	}
	if transaction == nil {
		return nil, errors.New("transaction not found")
	}

	if err := r.Db.Delete(transaction).Error; err != nil {
		return nil, fmt.Errorf("could not delete transaction: %w", err)
	}
	return transaction, nil
}

func (r *TransactionRepository) FindById(transactionId uint) (*model.Transaction, error) {
	var transaction model.Transaction
	if err := r.Db.First(&transaction, transactionId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No transaction found, return nil
		}
		return nil, fmt.Errorf("could not find transaction by id: %w", err)
	}
	return &transaction, nil
}

func (r *TransactionRepository) FindByTransactionHash(transactionHash string) (*model.Transaction, error) {
	var transaction model.Transaction
	if err := r.Db.Where("transaction_hash = ?", transactionHash).First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No transaction found, return nil
		}
		return nil, fmt.Errorf("could not find transaction by hash: %w", err)
	}
	return &transaction, nil
}

func (r *TransactionRepository) FindAll() ([]model.Transaction, error) {
	var transactions []model.Transaction
	if err := r.Db.Find(&transactions).Error; err != nil {
		return nil, fmt.Errorf("could not find all transactions: %w", err)
	}
	return transactions, nil
}
