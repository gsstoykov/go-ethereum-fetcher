package repository

import (
	"errors"
	"fmt"

	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(userId uint) error
	FindById(userId uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindAll() ([]model.User, error)
	FindUserTransactions(userID uint) ([]model.Transaction, error)
	AddTransactionToUser(user *model.User, tx *model.Transaction) error
}

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := r.Db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("could not create user: %w", err)
	}
	return user, nil
}

func (r *UserRepository) Update(user *model.User) (*model.User, error) {
	existingUser, err := r.FindById(user.ID)
	if err != nil {
		return nil, fmt.Errorf("could not find user by ID: %w", err)
	}
	if existingUser == nil {
		return nil, errors.New("user not found")
	}

	if err := r.Db.Model(existingUser).Updates(user).Error; err != nil {
		return nil, fmt.Errorf("could not update user: %w", err)
	}
	return existingUser, nil
}

func (r *UserRepository) Delete(userId uint) error {
	if err := r.Db.Delete(&model.User{}, userId).Error; err != nil {
		return fmt.Errorf("could not delete user: %w", err)
	}
	return nil
}

func (r *UserRepository) FindById(userId uint) (*model.User, error) {
	var user model.User
	if err := r.Db.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find user by ID: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.Db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, fmt.Errorf("could not find user by username: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.Db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("could not find all users: %w", err)
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
		return nil, fmt.Errorf("could not find transactions for user: %w", err)
	}
	return transactions, nil
}

func (r *UserRepository) AddTransactionToUser(user *model.User, tx *model.Transaction) error {
	if user == nil {
		return errors.New("user not found")
	}
	if err := r.Db.Model(user).Association("Transactions").Append(tx); err != nil {
		return fmt.Errorf("could not add transaction to user: %w", err)
	}
	return nil
}
