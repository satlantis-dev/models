package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type NewsletterStatus string

const (
	NewsletterStatusDraft     NewsletterStatus = "draft"
	NewsletterStatusScheduled NewsletterStatus = "scheduled"
	NewsletterStatusSending   NewsletterStatus = "sending"
	NewsletterStatusSent      NewsletterStatus = "sent"
	NewsletterStatusFailed    NewsletterStatus = "failed"
)

type CommunityNewsletter struct {
	ID           uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	CommunityID  uint             `gorm:"index" json:"communityId"`
	Community    *Community       `gorm:"constraint:OnDelete:CASCADE;" json:"community,omitempty"`
	AccountID    uint             `gorm:"index" json:"accountId"`
	Account      *Account         `gorm:"constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	ContentJson  json.RawMessage  `gorm:"type:jsonb" json:"contentJson,omitempty"`
	ContentHtml  string           `gorm:"type:text" json:"contentHtml,omitempty"`
	Subject      string           `gorm:"type:varchar(255)" json:"subject"`
	Status       NewsletterStatus `gorm:"type:varchar(50);default:'draft'" json:"status"`
	ScheduledFor *time.Time       `json:"scheduledFor,omitempty"`
	SentAt       *time.Time       `json:"sentAt,omitempty"`
	CreatedAt    time.Time        `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time        `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt    *gorm.DeletedAt  `gorm:"index" json:"-"`
}

func (CommunityNewsletter) TableName() string {
	return "community_newsletters"
}
