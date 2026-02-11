package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountCalendarRoleType int

const (
	CalendarOwner AccountCalendarRoleType = iota + 1
	CalendarContributor
	CalendarViewer
	CalendarInvited
)

type AccountCalendarRole struct {
	AccountID  uint                    `gorm:"not null;index;uniqueIndex:idx_account_calendar_role" json:"accountId"`
	Account    *Account                `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	CalendarID uint                    `gorm:"not null;index;uniqueIndex:idx_account_calendar_role" json:"calendarId"`
	Calendar   *Calendar               `gorm:"foreignKey:CalendarID;constraint:OnDelete:CASCADE;" json:"calendar,omitempty"`
	Type       AccountCalendarRoleType `gorm:"not null" json:"type"`
	CreatedAt  time.Time               `json:"-"`
	UpdatedAt  time.Time               `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt          `gorm:"index" json:"-"`
}

func (AccountCalendarRole) TableName() string {
	return "account_calendar_roles"
}
