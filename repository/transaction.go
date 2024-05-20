package repository

import (
	"errors"

	"github.com/gsstoykov/go-ethereum-fetcher/model"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	Create(transaction *model.Transaction) (*model.Transaction, error)
	Update(transaction *model.Transaction) (*model.Transaction, error)
	Delete(transactionId int) (*model.Transaction, error)
	FindById(transactionId int) (*model.Transaction, error)
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
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) Update(transaction *model.Transaction) (*model.Transaction, error) {
	existingTransaction, err := r.FindById(transaction.Id)
	if err != nil {
		return nil, err
	}
	if existingTransaction == nil {
		return nil, errors.New("transaction not found")
	}

	// Update transaction attributes

	// Perform the update operation
	if err := r.Db.Save(existingTransaction).Error; err != nil {
		return nil, err
	}
	return existingTransaction, nil
}

func (r *TransactionRepository) Delete(transactionId int) (*model.Transaction, error) {
	var transaction *model.Transaction
	if err := r.Db.Where("id = ?", transactionId).Delete(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) FindById(transactionId int) (*model.Transaction, error) {
	var transaction *model.Transaction
	if err := r.Db.Where("id = ?", transactionId).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) FindByTransactionHash(transactionHash string) (*model.Transaction, error) {
	var transaction model.Transaction
	if err := r.Db.Where("transaction_hash = ?", transactionHash).First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No transaction found, return nil
		}
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) FindAll() ([]model.Transaction, error) {
	var transactions []model.Transaction
	if err := r.Db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
