package main

import (
	"go-auth-api/database"
	"go-auth-api/models"
	"go-auth-api/routes"
	"go-auth-api/middleware" // Make sure this is imported

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// --- CORS CONFIGURATION ---
	// ... (Your existing CORS config) ...
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200",
								  "https://event-planner-frontend-production-c144.up.railway.app",},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.User{}, &models.Event{}, &models.EventParticipant{})

	r.POST("/register", routes.Signup)
	r.POST("/login", routes.Login)

    
    protected := r.Group("/")
    protected.Use(middleware.AuthMiddleware()) 
    {
        routes.EventRoutes(protected) 
    }

	r.Run(":8080")
}
