package database

import (
	"fmt"
	"os"

	"github.com/charlesonunze/busha-test/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("POSTGRES_URI")))
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")

	DB.AutoMigrate(&model.Movie{}, &model.Comment{}, &model.Character{})
	fmt.Println("Database Migrated")
}

func CloseDB() error {
	var db *gorm.DB

	pdb, err := db.DB()
	if err != nil {
		return err
	}

	pdb.Close()
	fmt.Println("Closing DB Connection")
	return nil
}
