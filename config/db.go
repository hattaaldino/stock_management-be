package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {

	var err error

	db_username := os.Getenv("db_username")
	db_password := os.Getenv("db_password")
	db_server := os.Getenv("db_server")
	db_port := os.Getenv("db_port")
	db_name := os.Getenv("db_name")

	dsn := fmt.Sprintf("host=%s	user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", db_server, db_username, db_password, db_name, db_port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	} else {
		return nil
	}
}
