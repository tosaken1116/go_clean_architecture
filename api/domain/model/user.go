package model

import "time"

type User struct {
	ID           int64     `gorm:"primary_key"`
	ScreenID     string    `gorm:"unique;size:14"`
	PasswordHash string
	UserName     string
	Email        string    `gorm:"unique;not null"`
	IsPublic     bool      `gorm:"default:true"`
	IconURL      string
	HeaderURL    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `gorm:"index"`
}