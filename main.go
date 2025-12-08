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
		AllowOrigins:[]string{
			"http://localhost:4200", 
			"https://event-planner-frontend-production-c144.up.railway.app",
		},
		// ... rest of CORS
	}))

	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.User{}, &models.Event{}, &models.EventParticipant{})

	// 1. 🚀 PUBLIC ROUTES (UNPROTECTED)
	// These routes are open to everyone
	r.POST("/register", routes.Signup)
	r.POST("/login", routes.Login)

    // 2. 🔒 PROTECTED ROUTES
    // Create a new group for authenticated access.
    protected := r.Group("/")
    protected.Use(middleware.AuthMiddleware()) // Apply AuthMiddleware only to this group
    {
        // Now, add the Event Routes using the protected group:
        // *** CRITICAL CHANGE HERE: Pass the protected group, not the main router 'r' ***
        routes.EventRoutes(protected) 
    }

	r.Run(":8080")
}
