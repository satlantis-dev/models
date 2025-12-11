package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountRSVPAnswers struct {
	ID              uint                    `gorm:"primaryKey" json:"id"`
	AccountID       uint                    `json:"accountId"`
	Account         Account                 `gorm:"constraint:OnDelete:CASCADE;" json:"account"`
	CalendarEventID uint                    `json:"calendarEventId"`
	CalendarEvent   CalendarEvent           `gorm:"constraint:OnDelete:CASCADE;" json:"calendarEvent"`
	RsvpStatus      RsvpStatus              `json:"rsvpStatus"`
	Answers         *map[string]interface{} `gorm:"type:jsonb;serializer:json" json:"answers,omitempty"`
	CreatedAt       time.Time               `json:"createdAt"`
	DeletedAt       gorm.DeletedAt          `gorm:"index" json:"-"`
}
