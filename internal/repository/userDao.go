package repository

import (
	"errors"
	model "hospetal/internal/models"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (*Repo, error) {
	if db == nil {
		return nil, errors.New("Error in NewRepo Connecting to DataBase")
	}
	return &Repo{db: db}, nil
}

type Patient interface {
	CreatePatient(p model.PetienUser) (model.PetienUser, error)
	FetchUserByEmail(string) (model.PetienUser, error)
}

func (r *Repo) CreatePatient(u model.PetienUser) (model.PetienUser, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return model.PetienUser{}, err
	}
	return u, nil
}
func (r *Repo) FetchUserByEmail(s string) (model.PetienUser, error) {
	var u model.PetienUser
	tx := r.db.Where("email=?", s).First(&u)
	if tx.Error != nil {
		return model.PetienUser{}, nil
	}
	return u, nil
}
