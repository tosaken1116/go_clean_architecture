package model

type ListUser struct {
	ListID string `gorm:"primary_key"`
	UserID string `gorm:"primary_key"`
}
