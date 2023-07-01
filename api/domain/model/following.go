package model

import "time"

type Following struct {
	UserID    string          `gorm:"primary_key"`
	TargetID  string          `gorm:"primary_key"`
	State     FollowingState `gorm:"type:following_state"`
	CreatedAt time.Time
}