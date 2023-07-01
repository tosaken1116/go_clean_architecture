package model

type Retweet struct {
	TweetID string `gorm:"primary_key"`
	UserID  string `gorm:"primary_key"`
}