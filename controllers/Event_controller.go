package controllers

import (
	"net/http"
	"strconv"
	"time"

	"go-auth-api/database"
	"go-auth-api/models"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	var input struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Date        time.Time `json:"date"`
		Time        string    `json:"time"`
		Location    string    `json:"location"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")

	event := models.Event{
		Title:       input.Title,
		Description: input.Description,
		Date:        input.Date,
		Time:        input.Time,
		Location:    input.Location,
		CreatedByID: userID.(uint),
	}

	if err := database.DB.Create(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	database.DB.Create(&models.EventParticipant{
		EventID: event.ID,
		UserID:  userID.(uint),
		Role:    "organizer",
	})

	c.JSON(http.StatusOK, gin.H{"message": "Event Created", "event": event})
}

func GetOrganizedEvents(c *gin.Context) {
	userID, _ := c.Get("userID")

	var events []models.Event
	database.DB.Where("created_by_id = ?", userID).Find(&events)

	c.JSON(http.StatusOK, gin.H{"organized_events": events})
}

func GetInvitedEvents(c *gin.Context) {
	userID, _ := c.Get("userID")

	var events []models.Event
	database.DB.Joins("JOIN event_participants ON events.id = event_participants.event_id").
		Where("event_participants.user_id = ? AND event_participants.role = ?", userID, "attendee").
		Find(&events)

	c.JSON(http.StatusOK, gin.H{"invited_events": events})
}

func InviteUserToEvent(c *gin.Context) {
	eventID, _ := strconv.Atoi(c.Param("id"))
	userID, _ := c.Get("userID")

	var ep models.EventParticipant
	if err := database.DB.Where("event_id = ? AND user_id = ? AND role = ?", eventID, userID, "organizer").First(&ep).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only organizer can invite others"})
		return
	}

	var body struct {
		InviteeID uint `json:"invitee_id"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&models.EventParticipant{
		EventID: uint(eventID),
		UserID:  body.InviteeID,
		Role:    "attendee",
	})

	c.JSON(http.StatusOK, gin.H{"message": "User invited successfully"})
}

func DeleteEvent(c *gin.Context) {
	eventID, _ := strconv.Atoi(c.Param("id"))
	userID, _ := c.Get("userID")

	var event models.Event
	if err := database.DB.First(&event, eventID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if event.CreatedByID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only event creator can delete it"})
		return
	}

	database.DB.Where("event_id = ?", eventID).Delete(&models.EventParticipant{})
	database.DB.Delete(&event)

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted sucessfully"})
}
