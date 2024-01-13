package model

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Department string `json:"dept"`
	Name       string `json:"name"`
}
