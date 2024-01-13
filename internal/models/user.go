package model

import "gorm.io/gorm"

type PetienUser struct {
	gorm.Model
	UserName string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	PhoNo    string `json:"phono" validate:"required"`
	PassHas  string `json:"-" validate:"required"`
}
type UserSignUp struct {
	UserName string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	PassHas  string `json:"passhas" validate:"required"`
	PhoNo    string `json:"phono" validate:"required"`
}
type UserLogin struct {
	Name    string `json:"name"`
	Email   string `json:"email" validate:"required"`
	PassHas string `json:"passhas" validate:"required"`
	PhoNo   string `json:"phono" validate:"required"`
}
