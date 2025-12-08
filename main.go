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

	// --- CORS CONFIGURATION ---
	r.Use(cors.New(cors.Config{
		// *** CRITICAL FIX: ADD YOUR LIVE FRONTEND URL HERE ***
		// You must replace the placeholder with the HTTPS URL provided by Railway 
		// for your event-planner-frontend service (e.g., https://my-app-xxxx.railway.app).
		AllowOrigins:     []string{
			"http://localhost:4200", // Keep for local development
			"https://event-planner-frontend-production-c144.up.railway.app", // <--- REPLACE THIS PLACEHOLDER
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// --- END CORS CONFIGURATION ---

	database.ConnectDatabase()

	// 2. Added Migration for Events and Participants so the tables get created
	database.DB.AutoMigrate(&models.User{}, &models.Event{}, &models.EventParticipant{})

	r.POST("/register", routes.Signup)
	r.POST("/login", routes.Login)

	// 3. Register the Event Routes (This enables /events URLs)
	routes.EventRoutes(r)

	r.Run(":8080")
}
