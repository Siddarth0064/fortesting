package database

import (
	"context"
	"errors"
	"fmt"
	//model "hospetal/internal/models"

	//"hospetal/internal/models"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("DataBase is Not Connecting to PostgresQl")
	}
	return db, nil
}
func Connection() (*gorm.DB, error) {
	log.Info().Msg("DataBase connection Initializing GORM.DB")
	db, err := Open()
	if err != nil {
		return nil, fmt.Errorf("Error in Initializing DataBase %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("PostgresQl failed to Connect %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error in PingContext %w", err)
	}

	//db.Migrator().DropTable(&model.PatienDeatiles{}, &model.PetienUser{})
	// err = db.Migrator().AutoMigrate(&model.UserLogin{}, &model.UserSignUp{}, &model.PetienUser{})
	// if err != nil {
	// 	return nil, fmt.Errorf("Error in Auto Migrate DataBase %w", err)
	// }
	return db, nil
}
