package entity

import (
	"fmt"

	"gorm.io/driver/mysql"

	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database ", DB)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", DB)
	}

	DB.AutoMigrate(&User{}, &Shop{}, &Product{}, &Type{}, &Report{})

}
