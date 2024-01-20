package model

import "gorm.io/gorm"

type PatienDeatiles struct {
	gorm.Model
	BloodGroup string     `json:"bloodgroup"`
	Age        string     `json:"age"`
	Place      string     `json:"place"`
	Disease    string     `json:"disease"`
	PetienUser PetienUser `gorm:"ForeignKey:uid"`
	//	Uid        uint64     `JSON:"uid, omitempty"`
}
type Department struct {
	gorm.Model
	Dept string `json:"deptName"`
	Date string `json:"date"`
}
