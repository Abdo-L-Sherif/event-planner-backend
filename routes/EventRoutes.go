package routes

import (
	"go-auth-api/controllers"
	"go-auth-api/middleware"

	"github.com/gin-gonic/gin"
)

func EventRoutes(router *gin.Engine) {
	event := router.Group("/events")
	event.Use(middleware.AuthMiddleware())

	event.POST("/", controllers.CreateEvent)
	event.GET("/organized", controllers.GetOrganizedEvents)
	event.GET("/invited", controllers.GetInvitedEvents)
	event.POST("/:id/invite", controllers.InviteUserToEvent)
	event.DELETE("/:id", controllers.DeleteEvent)
}
