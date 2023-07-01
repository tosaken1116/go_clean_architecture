package model

import "time"
type List struct {
	ID        string     `gorm:"primary_key"`
	Name      string
	CreatedBy string     `gorm:"type:uuid"`
	CreatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	IsPublic  bool       `gorm:"default:true"`
}