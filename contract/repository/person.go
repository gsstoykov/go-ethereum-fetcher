package repository

import (
	"github.com/gsstoykov/go-ethereum-fetcher/contract/model"
	"gorm.io/gorm"
)

type IPersonRepository interface {
	Create(person *model.Person) (*model.Person, error)
	Delete(personId uint) (*model.Person, error)
	FindAll() ([]model.Person, error)
}

type PersonRepository struct {
	Db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) IPersonRepository {
	return &PersonRepository{Db: db}
}

func (r *PersonRepository) Create(person *model.Person) (*model.Person, error) {
	if err := r.Db.Create(person).Error; err != nil {
		return nil, err
	}
	return person, nil
}

func (r *PersonRepository) Delete(personId uint) (*model.Person, error) {
	var person *model.Person
	if err := r.Db.Where("id = ?", personId).Delete(person).Error; err != nil {
		return nil, err
	}
	return person, nil
}

func (r *PersonRepository) FindAll() ([]model.Person, error) {
	var people []model.Person
	if err := r.Db.Find(&people).Error; err != nil {
		return nil, err
	}
	return people, nil
}
