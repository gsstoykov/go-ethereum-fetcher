package repository

import (
	"errors"
	"fmt"

	"github.com/gsstoykov/go-ethereum-fetcher/contract/model"
	"gorm.io/gorm"
)

// IPersonRepository defines the interface for person repository operations.
type IPersonRepository interface {
	Create(person *model.Person) (*model.Person, error)
	Delete(personId uint) (*model.Person, error)
	FindAll() ([]model.Person, error)
	GetByID(personId uint) (*model.Person, error)
}

// PersonRepository implements the IPersonRepository interface.
type PersonRepository struct {
	Db *gorm.DB
}

// NewPersonRepository creates a new instance of PersonRepository.
func NewPersonRepository(db *gorm.DB) IPersonRepository {
	return &PersonRepository{Db: db}
}

// Create inserts a new person record into the database.
func (r *PersonRepository) Create(person *model.Person) (*model.Person, error) {
	if err := r.Db.Create(person).Error; err != nil {
		return nil, fmt.Errorf("PersonRepository.Create: failed to create person %v: %w", person, err)
	}
	return person, nil
}

// Delete removes a person record from the database by its ID.
func (r *PersonRepository) Delete(personId uint) (*model.Person, error) {
	var person model.Person
	if err := r.Db.First(&person, personId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("PersonRepository.Delete: person with ID %d not found: %w", personId, err)
		}
		return nil, fmt.Errorf("PersonRepository.Delete: failed to find person with ID %d: %w", personId, err)
	}

	if err := r.Db.Delete(&person).Error; err != nil {
		return nil, fmt.Errorf("PersonRepository.Delete: failed to delete person with ID %d: %w", personId, err)
	}
	return &person, nil
}

// FindAll retrieves all person records from the database.
func (r *PersonRepository) FindAll() ([]model.Person, error) {
	var people []model.Person
	if err := r.Db.Find(&people).Error; err != nil {
		return nil, fmt.Errorf("PersonRepository.FindAll: failed to find all people: %w", err)
	}
	return people, nil
}

// GetByID retrieves a person record by its ID.
func (r *PersonRepository) GetByID(personId uint) (*model.Person, error) {
	var person model.Person
	if err := r.Db.First(&person, personId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("PersonRepository.GetByID: failed to find person with ID %d: %w", personId, err)
	}
	return &person, nil
}
