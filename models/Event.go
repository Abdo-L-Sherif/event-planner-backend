package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID           uint               `gorm:"primaryKey" json:"id"`
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	Date         time.Time          `json:"date"`
	Time         string             `json:"time"`
	Location     string             `json:"location"`
	CreatedByID  uint               `json:"created_by"`
	CreatedBy    User               `gorm:"foreignKey:CreatedByID" json:"-"`
	Participants []EventParticipant `json:"participants"`
}
