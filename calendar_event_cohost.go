package models

import (
	"time"
)

type CalendarEventCohost struct {
	ID                   uint           `gorm:"primaryKey" json:"id"`
	CalendarEventID      uint           `gorm:"index;not null" json:"calendarEventId"`
	CalendarEvent        *CalendarEvent `gorm:"foreignKey:CalendarEventID" json:"calendarEvent,omitempty"`
	AccountID            uint           `gorm:"index;not null" json:"accountId"`
	Account              *AccountDTO    `gorm:"foreignKey:AccountID" json:"account"`
	InvitationAcceptedAt *time.Time     `json:"invitationAcceptedAt"`
	InvitationDeclinedAt *time.Time     `json:"invitationDeclinedAt"`
	AutoAcceptInvitation bool           `gorm:"default:false" json:"-"`
	IsEmailAdded         bool           `gorm:"default:false" json:"-"`
	CreatedAt            time.Time      `json:"createdAt"`
	UpdatedAt            time.Time      `json:"updatedAt"`
}
