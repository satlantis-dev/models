package models

import (
	"time"
)

// Platform defines the type of platform where push notifications are sent iOS or Android
type Platform string

const (
	PlatformIOS     Platform = "IOS"
	PlatformAndroid Platform = "ANDROID"
)

// Device represents a user device for general purposes and receiving push notifications
type Device struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	AccountID  uint       `gorm:"index" json:"accountId"`
	Account    AccountDTO `gorm:"foreignKey:AccountID" json:"account"`
	DeviceID   string     `json:"deviceId"`
	Token      *string    `json:"token"`
	Platform   Platform   `json:"platform"`
	AppVersion *string    `json:"appVersion,omitempty"`
	CreatedAt  time.Time  `json:"createdAt"`
	DeletedAt  *time.Time `gorm:"index" json:"-,omitempty"`
}
