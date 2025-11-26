package models

import "gorm.io/gorm"

type EventParticipant struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey" json:"id"`
	UserID  uint   `json:"user_id"`
	EventID uint   `json:"event_id"`
	Role    string `json:"role"`   // "organizer" or "attendee"
	Status  string `json:"status"` // "Going", "Maybe", "Not Going", "Pending"
	User    User   `gorm:"foreignKey:UserID" json:"user"`
	Event   Event  `gorm:"foreignKey:EventID" json:"-"`
}
