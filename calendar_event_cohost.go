package models

import (
	"time"
)

type CalendarEventCohostRoleType string

const (
	CalendarEventCohostRoleAdmin CalendarEventCohostRoleType = "admin"
	CalendarEventCohostRoleStaff CalendarEventCohostRoleType = "staff"
)

type CalendarEventCohost struct {
	ID                   uint                        `gorm:"primaryKey" json:"id"`
	CalendarEventID      uint                        `gorm:"not null;index;index:idx_cohost_event_role" json:"calendarEventId"`
	CalendarEvent        *CalendarEvent              `gorm:"foreignKey:CalendarEventID" json:"calendarEvent,omitempty"`
	AccountID            uint                        `gorm:"not null;index;index:idx_cohost_event_role" json:"accountId"`
	Account              *AccountDTO                 `gorm:"foreignKey:AccountID" json:"account"`
	Type                 CalendarEventCohostRoleType `gorm:"type:varchar(32);not null;default:'admin';index:idx_cohost_event_role" json:"type"`
	InvitationReceivedAt *time.Time                  `gorm:"type:timestamptz" json:"invitationReceivedAt,omitempty"`
	InvitationAcceptedAt *time.Time                  `gorm:"type:timestamptz" json:"invitationAcceptedAt,omitempty"`
	InvitationDeclinedAt *time.Time                  `gorm:"type:timestamptz" json:"invitationDeclinedAt,omitempty"`
	AutoAcceptInvitation bool                        `gorm:"default:false" json:"-"`
	CreatedAt            time.Time                   `json:"createdAt"`
	UpdatedAt            time.Time                   `json:"updatedAt"`
}
