package model

import "time"

type Tweet struct {
	ID          string     `gorm:"primary_key"`
	Description string     `gorm:"size:200"`
	TweetUser   string     `gorm:"type:uuid"`
	Activity    int        `gorm:"default:0"`
	CreatedAt   time.Time
	DeletedAt   *time.Time `gorm:"index"`
	ReplyTo     *Tweet     `gorm:"foreignkey:ReplyToID"`
	ReplyToID   *string
}
