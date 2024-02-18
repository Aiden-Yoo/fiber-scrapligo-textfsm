package database

import (
	"log"
	"os"

	"github.com/aiden-yoo/getthemall/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	env := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.AutoMigrate(&model.Device{})
	if err != nil {
		log.Fatal(err)
	}

}
