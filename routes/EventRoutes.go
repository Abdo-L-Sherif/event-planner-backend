package routes

import (
	"go-auth-api/controllers"

	"github.com/gin-gonic/gin"
)

func EventRoutes(r *gin.RouterGroup) {
	events := r.Group("/events")
	// events.Use(middleware.AuthMiddleware()) // Protect all event routes
	{
		events.POST("/", controllers.CreateEvent)
		events.GET("/organized", controllers.GetOrganizedEvents)
		events.GET("/invited", controllers.GetInvitedEvents)
		events.DELETE("/:id", controllers.DeleteEvent)
		events.POST("/:id/invite", controllers.InviteUserToEvent)

		// NEW ROUTES
		events.GET("/search", controllers.SearchEvents)
		events.GET("/:id", controllers.GetEventById)
		events.POST("/:id/rsvp", controllers.RSVPEvent)
	}
}
