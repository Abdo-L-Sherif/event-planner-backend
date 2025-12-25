package database

import (
	"fmt"
	"log"
	"os"
<<<<<<< HEAD
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
=======

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
<<<<<<< HEAD
=======

>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
	err := godotenv.Load()
	if err != nil {
		log.Println("Note: .env file not found, using system environment variables")
	}

<<<<<<< HEAD
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "mysql" // Default to MySQL for OpenShift compatibility
	}

	var database *gorm.DB

	// Retry logic for database connection
	maxRetries := 5
	var connectionErr error

	for i := 0; i < maxRetries; i++ {
		if dbType == "sqlite" {
			// SQLite configuration
			dbPath := os.Getenv("DB_PATH")
			if dbPath == "" {
				dbPath = "users.db"
			}
			database, connectionErr = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		} else {
			// MySQL configuration (default for OpenShift)
			dbUser := getEnvWithDefault("DB_USER", "root")
			dbPass := getEnvWithDefault("DB_PASSWORD", "")
			dbName := getEnvWithDefault("DB_NAME", "eventplanner")
			dbHost := getEnvWithDefault("DB_HOST", "localhost")
			dbPort := getEnvWithDefault("DB_PORT", "3306")

			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				dbUser, dbPass, dbHost, dbPort, dbName)

			database, connectionErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		}

		if connectionErr == nil {
			break
		}

		log.Printf("Database connection attempt %d/%d failed: %v", i+1, maxRetries, connectionErr)
		if i < maxRetries-1 {
			time.Sleep(time.Duration(i+1) * time.Second) // Exponential backoff
		}
	}

	if connectionErr != nil {
		log.Fatalf("Database connection failed after %d attempts: %v", maxRetries, connectionErr)
	}

	if err != nil {
		log.Fatalf("Database connection failed after %d attempts: %v", maxRetries, err)
	}

	// Test the connection
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	// Configure connection pool for production
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Printf("Successfully connected to %s database", dbType)
	DB = database
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
=======
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
	DB = database
}
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
