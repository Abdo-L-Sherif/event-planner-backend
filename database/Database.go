package database

import (
<<<<<<< HEAD
	"gorm.io/driver/sqlite"
=======
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
>>>>>>> 6bafcd2 (Replace SQLite setup with mysql)
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
<<<<<<< HEAD
	database, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Database couldn't connect")
	}

=======

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database couldn't connect: %v", err)
	}

	fmt.Println("Connected to database sucessfully")
>>>>>>> 6bafcd2 (Replace SQLite setup with mysql)
	DB = database
}
