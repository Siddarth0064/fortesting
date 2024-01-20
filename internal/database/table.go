package database

import (
	model "hospetal/internal/models"
	"log"
)

func CreateTable() {
	// Drop the table student if it exists
	db, err := Open()
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Migrator().DropTable(&model.PatienDeatiles{})
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Migrator().AutoMigrate(&model.PatienDeatiles{})
	if err != nil {
		log.Fatalln(err)
	}
}
