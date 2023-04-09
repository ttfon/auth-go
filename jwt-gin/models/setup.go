package models

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Can't load variable from .env file %v", err)
	}
	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASS")
	DbName := os.Getenv("DB_NAME")
	fmt.Println(DbUser, DbPassword, DbHost, DbPort, DbName)
	Dburl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(Dbdriver, Dburl)

	if err != nil {
		fmt.Println("Cann't connect to database", Dbdriver)
		log.Fatalf("Connection failed: %v", err)
	} else {
		fmt.Println("Connected to database", Dbdriver)
	}

	DB.AutoMigrate(&User{})

}
