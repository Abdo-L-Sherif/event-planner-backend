package main

import (
	"go-auth-api/database"
	"go-auth-api/models"
	"go-auth-api/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. Fixed the "DELET" typo to "DELETE"
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	database.ConnectDatabase()

	// 2. Added Migration for Events and Participants so the tables get created
	database.DB.AutoMigrate(&models.User{}, &models.Event{}, &models.EventParticipant{})

	r.POST("/register", routes.Signup)
	r.POST("/login", routes.Login)

	// 3. Register the Event Routes (This enables /events URLs)
	// Note: I am assuming the function is named 'EventRoutes' based on your file list.
	// If this line causes an error, check the function name inside routes/EventRoutes.go
	routes.EventRoutes(r)

	r.Run(":8080")
}
