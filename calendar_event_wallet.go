package models

import (
	"gorm.io/gorm"
	"time"
)

type WalletProvider string

var (
	WalletProviderIbex WalletProvider = "IBEX"
)

type CalendarEventWallet struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	CalendarEventID  uint           `gorm:"not null;index" json:"calendarEventId"`
	Name             string         `json:"name"`
	Provider         WalletProvider `gorm:"type:varchar(16)" json:"provider"`
	ProviderWalletId *string        `json:"providerWalletId"`
	ProviderUserId   *string        `json:"providerUserId"`
	Balance          *int64         `json:"balance"`
	Currency         *string        `gorm:"type:varchar(8)" json:"currency"`
	CreatedAt        time.Time      `json:"createdAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
