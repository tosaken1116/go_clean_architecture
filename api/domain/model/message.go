package model

import "time"

type Message struct {
	ID          string    `gorm:"primary_key"`
	Description string
	SenderID    string    `gorm:"type:uuid"`
	RecipientID string    `gorm:"type:uuid"`
	CreatedAt   time.Time
}