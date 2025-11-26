package controllers

import (
	"net/http"
	"strconv"
	"time"

	"go-auth-api/database"
	"go-auth-api/models"

	"github.com/gin-gonic/gin"
)

// --- EXISTING FUNCTIONS ---

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

	// Add Creator as Organizer with "Going" status
	database.DB.Create(&models.EventParticipant{
		EventID: event.ID,
		UserID:  userID.(uint),
		Role:    "organizer",
		Status:  "Going",
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

	// Check if requester is the organizer
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

	// Create participant with "Pending" status
	database.DB.Create(&models.EventParticipant{
		EventID: uint(eventID),
		UserID:  body.InviteeID,
		Role:    "attendee",
		Status:  "Pending",
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

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

// --- NEW FUNCTIONS FOR PHASE 1 REQUIREMENTS ---

// SearchEvents handles filtering by keyword, date, and user role
func SearchEvents(c *gin.Context) {
	query := c.Query("q")
	dateParam := c.Query("date") // Expect format YYYY-MM-DD
	roleParam := c.Query("role") // 'organizer' or 'attendee'
	userID, _ := c.Get("userID")

	// Start building the query
	tx := database.DB.Model(&models.Event{})

	// 1. Keyword Filter (Title or Description)
	if query != "" {
		tx = tx.Where("(title LIKE ? OR description LIKE ?)", "%"+query+"%", "%"+query+"%")
	}

	// 2. Date Filter
	if dateParam != "" {
		// Use DATE() function to compare just the day part, ignoring time
		tx = tx.Where("DATE(date) = ?", dateParam)
	}

	// 3. User Role Filter
	if roleParam == "organizer" {
		// Show events I created
		tx = tx.Where("created_by_id = ?", userID)
	} else if roleParam == "attendee" {
		// Show events I am invited to (requires Join)
		tx = tx.Joins("JOIN event_participants ON events.id = event_participants.event_id").
			Where("event_participants.user_id = ? AND event_participants.role = ?", userID, "attendee")
	}

	var events []models.Event
	if err := tx.Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed"})
		return
	}

	c.JSON(http.StatusOK, events)
}

// 2. Get Event By ID (Details Page)
func GetEventById(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")

	var event models.Event
	if err := database.DB.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Fetch all participants for this event
	var participants []models.EventParticipant
	// Preload "User" so we can see the names/emails of attendees if needed
	database.DB.Preload("User").Where("event_id = ?", id).Find(&participants)

	// Determine the role of the user requesting this data
	userRole := "viewer" // default
	if event.CreatedByID == userID.(uint) {
		userRole = "organizer"
	} else {
		for _, p := range participants {
			if p.UserID == userID.(uint) {
				userRole = "attendee"
				break
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"event":        event,
		"participants": participants,
		"userRole":     userRole,
	})
}

// 3. RSVP to Event
func RSVPEvent(c *gin.Context) {
	eventID := c.Param("id")
	userID, _ := c.Get("userID")

	var input struct {
		Status string `json:"status"` // Going, Maybe, Not Going
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
		return
	}

	// Find the participant entry
	var participant models.EventParticipant
	if err := database.DB.Where("event_id = ? AND user_id = ?", eventID, userID).First(&participant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "You are not invited to this event"})
		return
	}

	// Update status
	participant.Status = input.Status
	database.DB.Save(&participant)

	c.JSON(http.StatusOK, gin.H{"message": "RSVP updated", "status": input.Status})
}
