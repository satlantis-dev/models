package models

import "time"

type CalendarSubscription struct {
	AccountID  uint      `gorm:"index;uniqueIndex:idx_calendar_subscription_unique" json:"accountId"`
	Account    *Account  `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	CalendarID uint      `gorm:"index;uniqueIndex:idx_calendar_subscription_unique" json:"calendarId"`
	Calendar   *Calendar `gorm:"foreignKey:CalendarID;constraint:OnDelete:CASCADE;" json:"calendar,omitempty"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}

func (CalendarSubscription) TableName() string {
	return "calendar_subscriptions"
}
