package model

import "time"

type Like struct {
	TweetID   string     `gorm:"primary_key"`
	UserID    string     `gorm:"primary_key"`
	CreatedAt time.Time
}
