package database

import "time"

type SocialMediaType int

const (
	Facebook SocialMediaType = iota + 1
	Instagram
	LinkedIn
	Twitter
)

type SocialMedia struct {
	ID              uint            `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time       `json:"-"`
	UpdatedAt       time.Time       `json:"-"`
	DeletedAt       *time.Time      `gorm:"index" json:"-"`
	Name            string          `gorm:"type:text" json:"name"`
	Link            string          `gorm:"type:text" json:"link"`
	AccountID       uint            `gorm:"index" json:"accountId"`
	SocialMediaType SocialMediaType `json:"socialMediaType"`
}
