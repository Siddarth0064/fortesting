package repository

import (
	"errors"
	"fmt"
	model "hospetal/internal/models"

	//"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (*Repo, error) {
	if db == nil {
		return nil, errors.New("error in NewRepo Connecting to DataBase")
	}
	return &Repo{db: db}, nil
}

type Patient interface {
	CreatePatient(p model.PetienUser) (model.PetienUser, error)
	FetchUserByEmail(string) (model.PetienUser, error)
	FetchUserByPhono(id uint64) (model.PetienUser, error)
	CreatePD(pd model.PatienDeatiles) (model.PatienDeatiles, error)
}

func (r *Repo) CreatePatient(u model.PetienUser) (model.PetienUser, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return model.PetienUser{}, errors.New("Error in creating patient user details %w")

		//return model.PetienUser{}, log.Error().Msg("Error in creating patient user deatils", err)
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
func (r *Repo) FetchUserByPhono(pho uint64) (model.PetienUser, error) {
	var u model.PetienUser
	tx := r.db.Where("pho_no=?", pho)
	if tx.Error != nil {
		fmt.Println("error in fetching phone nummber===========================", pho)
		return model.PetienUser{}, nil
	}
	return u, nil
}
func (r *Repo) CreatePD(pd model.PatienDeatiles) (model.PatienDeatiles, error) {
	err := r.db.Create(&pd).Error
	if err != nil {
		return model.PatienDeatiles{}, err
	}
	return pd, nil
}
