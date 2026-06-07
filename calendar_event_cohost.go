package models

import (
	"time"
)

type CalendarEventCohostRole string

const (
	CalendarEventCohostRoleAdmin CalendarEventCohostRole = "admin"
	CalendarEventCohostRoleStaff CalendarEventCohostRole = "staff"
)

type CalendarEventCohost struct {
	ID                   uint                    `gorm:"primaryKey" json:"id"`
	CalendarEventID      uint                    `gorm:"index:idx_cohost_event_role,not null" json:"calendarEventId"`
	CalendarEvent        *CalendarEvent          `gorm:"foreignKey:CalendarEventID" json:"calendarEvent,omitempty"`
	AccountID            uint                    `gorm:"index;not null" json:"accountId"`
	Account              *AccountDTO             `gorm:"foreignKey:AccountID" json:"account"`
	Role                 CalendarEventCohostRole `gorm:"index:idx_cohost_event_role;default:admin" json:"role"`
	InvitationAcceptedAt *time.Time              `json:"invitationAcceptedAt"`
	InvitationDeclinedAt *time.Time              `json:"invitationDeclinedAt"`
	AutoAcceptInvitation bool                    `gorm:"default:false" json:"-"`
	CreatedAt            time.Time               `json:"createdAt"`
	UpdatedAt            time.Time               `json:"updatedAt"`
}
