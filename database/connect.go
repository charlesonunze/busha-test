package database

import (
	"fmt"

	"github.com/charlesonunze/busha-test/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	DB, err = gorm.Open(postgres.Open("postgres://mpgzeajm:W4HnsFei0jSRGhhGfsQ9JkN6H4BofTLJ@batyr.db.elephantsql.com/mpgzeajm"))
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
