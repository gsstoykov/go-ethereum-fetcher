package repository

import (
	"errors"

	"github.com/gsstoykov/go-ethereum-fetcher/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(userId uint) (*model.User, error)
	FindById(userId uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindAll() ([]model.User, error)
	FindUserTransactions(userID uint) ([]model.Transaction, error)
	AddTransactionToUser(user model.User, tx model.Transaction) error
}

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := r.Db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(user *model.User) (*model.User, error) {
	existingUser, err := r.FindById(user.ID)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("user not found")
	}

	// Update user attributes
	existingUser.Username = user.Username

	// Perform the update operation
	if err := r.Db.Save(existingUser).Error; err != nil {
		return nil, err
	}
	return existingUser, nil
}

func (r *UserRepository) Delete(userId uint) (*model.User, error) {
	var user model.User
	if err := r.Db.Where("id = ?", userId).Delete(user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindById(userId uint) (*model.User, error) {
	var user model.User
	if err := r.Db.Where("id = ?", userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.Db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindUserTransactions(userID uint) ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.Db.Table("transactions").
		Select("transactions.*").
		Joins("JOIN user_transactions ON user_transactions.transaction_id = transactions.id").
		Where("user_transactions.user_id = ?", userID).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *UserRepository) AddTransactionToUser(user model.User, tx model.Transaction) error {
	// Append transaction to user's transactions
	if err := r.Db.Model(&user).Association("Transactions").Append(&tx); err != nil {
		return err
	}

	return nil
}
