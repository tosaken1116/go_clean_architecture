package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID           string     `gorm:"primary_key"`
	ScreenID     string    `gorm:"unique;size:14"`
	PasswordHash string
	UserName     string
	Email        string    `gorm:"unique;not null"`
	IsPublic     bool      `gorm:"default:true"`
	IconURL      sql.NullString
	HeaderURL    sql.NullString
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
}