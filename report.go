package models

import "time"

type ContentReport struct {
	ID             uint                    `gorm:"primaryKey" json:"id"`
	CreatedByID    uint                    `gorm:"index" json:"createdById"`
	CreatedBy      AccountDTO              `json:"createdBy"`
	ReportedUserID *uint                   `gorm:"index" json:"reportedUserId"`
	ReportedUser   *AccountDTO             `json:"reportedUser"`
	ReportedItemID *uint                   `gorm:"index" json:"reportedItemId"`
	Type           string                  `json:"type"`
	Reason         string                  `json:"reason"`
	Link           string                  `json:"link"`
	Payload        *map[string]interface{} `gorm:"type:jsonb;serializer:json" json:"payload,omitempty"`
	IP             string                  `json:"ip"`
	Location       *string                 `json:"location"`
	Device         *string                 `json:"device"`
	AppVersion     *string                 `json:"appVersion"`
	CreatedAt      time.Time               `json:"-"`
	DeletedAt      *time.Time              `json:"-,omitempty"`
}
