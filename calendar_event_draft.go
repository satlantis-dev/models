package models

import (
	"time"

	"gorm.io/datatypes"
)

type CalendarEventDraft struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	AccountID uint           `gorm:"index;not null" json:"accountId"`
	Account   *Account       `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"account,omitempty"`
	Title     string         `json:"title"`
	Draft     datatypes.JSON `gorm:"type:jsonb" json:"draft"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
}

