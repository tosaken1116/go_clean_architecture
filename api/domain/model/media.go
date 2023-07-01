package model
type Media struct {
	ID    string        `gorm:"primary_key"`
	Tweet string        `gorm:"type:uuid"`
	Type  ContentType `gorm:"type:content_type"`
}