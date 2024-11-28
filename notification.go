package models

import (
	"time"
)

type Notification struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	CreatorAccountID   uint      `gorm:"index" json:"creatorAccountId"` // Account that created the event
	CreatorAccount     Account   `gorm:"foreignKey:CreatorAccountID;references:ID" json:"creatorAccount"`
	RecipientAccountID uint      `gorm:"index" json:"-"` // Account that will be notified
	Type               string    `gorm:"type:varchar(255)" json:"type"`
	Message            string    `gorm:"type:text" json:"message"`
	IsRead             bool      `gorm:"default:false" json:"isRead"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"-"`
}
