package models

import (
	"gorm.io/gorm"
	"time"
)

type AccountRSVPAnswers struct {
	ID              uint                    `gorm:"primaryKey" json:"id"`
	AccountID       uint                    `json:"accountId"`
	Account         Account                 `json:"account"`
	CalendarEventID uint                    `json:"calendarEventId"`
	CalendarEvent   CalendarEvent           `json:"calendarEvent"`
	RsvpStatus      RsvpStatus              `json:"rsvpStatus"`
	Answers         *map[string]interface{} `gorm:"type:jsonb;serializer:json" json:"answers,omitempty"`
	CreatedAt       time.Time               `json:"createdAt"`
	DeletedAt       gorm.DeletedAt          `gorm:"index" json:"-"`
}
