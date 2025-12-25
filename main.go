package main

import (
<<<<<<< HEAD
	"context"
	"go-auth-api/database"
	"go-auth-api/models"
	"go-auth-api/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
=======
	"go-auth-api/database"
	"go-auth-api/models"
	"go-auth-api/routes"
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. Fixed the "DELET" typo to "DELETE"
<<<<<<< HEAD
	// Configure CORS with environment-based origins
	corsOrigins := os.Getenv("CORS_ORIGINS")
	if corsOrigins == "" {
		// Default for development - restrict in production
		corsOrigins = "http://localhost:3000,http://localhost:8080"
	}

	origins := strings.Split(corsOrigins, ",")
	// Trim spaces from origins
	for i, origin := range origins {
		origins[i] = strings.TrimSpace(origin)
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
=======
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
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

<<<<<<< HEAD
	// Health check endpoints for OpenShift probes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"time":   time.Now().Unix(),
		})
	})

	r.GET("/ready", func(c *gin.Context) {
		// Check database connectivity
		if database.DB != nil {
			sqlDB, err := database.DB.DB()
			if err == nil && sqlDB.Ping() == nil {
				c.JSON(200, gin.H{
					"status": "ready",
					"time":   time.Now().Unix(),
				})
				return
			}
		}
		c.JSON(503, gin.H{
			"status": "not ready",
			"time":   time.Now().Unix(),
		})
	})

	// Get port from environment variable with default fallback
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Validate port is a number
	if _, err := strconv.Atoi(port); err != nil {
		port = "8080"
	}

	// Start server with graceful shutdown
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Channel to listen for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-quit
	log.Println("Server shutting down...")

	// Context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
=======
	r.Run(":8080")
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
}
