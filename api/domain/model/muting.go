package model
type Muting struct {
	UserID   string `gorm:"primary_key"`
	TargetID string `gorm:"primary_key"`
}
