package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type EventWalletWithdrawalStatus string

const (
	EventWalletWithdrawalPending    EventWalletWithdrawalStatus = "pending"
	EventWalletWithdrawalProcessing EventWalletWithdrawalStatus = "processing"
	EventWalletWithdrawalCompleted  EventWalletWithdrawalStatus = "completed"
	EventWalletWithdrawalFailed     EventWalletWithdrawalStatus = "failed"
	EventWalletWithdrawalCancelled  EventWalletWithdrawalStatus = "cancelled"
)

type EventWalletWithdrawalMethod string

const (
	EventWalletWithdrawalMethodLightning EventWalletWithdrawalMethod = "lightning"
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

type CalendarEventWalletWithdrawal struct {
	ID                    uint                        `gorm:"primaryKey" json:"id"`
	CalendarEventID       uint                        `gorm:"not null;index" json:"calendarEventId"`
	CalendarEvent         *CalendarEvent              `gorm:"foreignKey:CalendarEventID" json:"calendarEvent,omitempty"`
	AccountID             uint                        `gorm:"not null;index" json:"accountId"`
	Account               *Account                    `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	Amount                int64                       `gorm:"type:bigint;not null" json:"amount"` // Cents for fiat, sats for BTC/Lightning
	Currency              OrderCurrency               `gorm:"type:varchar(8);not null" json:"currency"`
	Fee                   int64                       `gorm:"type:bigint;default:0" json:"fee"`      // Processing fee deducted
	NetAmount             int64                       `gorm:"type:bigint;not null" json:"netAmount"` // Amount after fees
	Status                EventWalletWithdrawalStatus `gorm:"type:varchar(32);default:'pending'" json:"status"`
	WithdrawalMethod      EventWalletWithdrawalMethod `gorm:"type:varchar(32);not null" json:"withdrawalMethod"`
	DestinationAddress    string                      `gorm:"type:text;not null" json:"destinationAddress"` // Lightning address, bank account, Stripe ID, etc.
	LightningPaymentHash  *string                     `gorm:"uniqueIndex;size:64" json:"lightningPaymentHash,omitempty"`
	LightningPreimage     *string                     `gorm:"size:64" json:"lightningPreimage,omitempty"`
	LightningProvider     *string                     `gorm:"size:32" json:"lightningProvider,omitempty"`
	LightningProviderTxID *string                     `gorm:"index" json:"lightningProviderTxId,omitempty"`
	FailureReason         *string                     `gorm:"type:text" json:"failureReason,omitempty"`
	Metadata              *datatypes.JSON             `gorm:"type:jsonb" json:"metadata,omitempty"`
	RequestedAt           time.Time                   `json:"requestedAt"`
	ProcessingStartedAt   *time.Time                  `json:"processingStartedAt,omitempty"`
	CompletedAt           *time.Time                  `json:"completedAt,omitempty"`
	FailedAt              *time.Time                  `json:"failedAt,omitempty"`
	CancelledAt           *time.Time                  `json:"cancelledAt,omitempty"`
	CreatedAt             time.Time                   `json:"createdAt"`
	UpdatedAt             time.Time                   `json:"updatedAt"`
	DeletedAt             gorm.DeletedAt              `gorm:"index" json:"-"`
}
