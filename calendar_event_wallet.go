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
	LightningAddress *string        `gorm:"uniqueIndex" json:"lightningAddress,omitempty"`
	Balance          *int64         `json:"balance"` // Balance in sats
	BalanceMsat      *int64         `gorm:"type:bigint" json:"balanceMsat,omitempty"`
	Currency         *string        `gorm:"type:varchar(8)" json:"currency"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
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

// CalendarEventWalletTransaction stores all event wallet transactions.
type CalendarEventWalletTransaction struct {
	ID                        uint                           `gorm:"primaryKey" json:"id"`
	CalendarEventWalletID     uint                           `gorm:"not null;index" json:"calendarEventWalletId"`
	CalendarEventWallet       *CalendarEventWallet           `gorm:"foreignKey:CalendarEventWalletID;constraint:OnDelete:CASCADE" json:"-"`
	CalendarEventID           uint                           `gorm:"not null;index" json:"calendarEventId"`
	TransactionType           AccountWalletTransactionType   `gorm:"type:varchar(32);not null" json:"transactionType"`
	Status                    AccountWalletTransactionStatus `gorm:"type:varchar(32);default:'pending'" json:"status"`
	CounterpartyTransactionID *uint                          `gorm:"index" json:"counterpartyTransactionId,omitempty"`

	// FROM Side
	FromAccountID  *uint    `gorm:"index" json:"fromAccountId,omitempty"`
	FromAccount    *Account `gorm:"foreignKey:FromAccountID;constraint:OnDelete:SET NULL" json:"-"`
	FromIdentifier *string  `json:"fromIdentifier,omitempty"`
	FromAmount     int64    `gorm:"type:bigint;not null" json:"fromAmount"` // Amount in smallest unit (sats, cents)
	FromAmountMsat *int64   `gorm:"type:bigint" json:"fromAmountMsat,omitempty"`
	FromCurrency   string   `gorm:"type:varchar(16);not null;default:'BTC';index" json:"fromCurrency"` // BTC, USD, etc.
	FromAmountUsd  *float64 `json:"fromAmountUsd,omitempty"`                                           // USD value at time of transaction

	// TO Side
	ToAccountID  *uint    `gorm:"index" json:"toAccountId,omitempty"`
	ToAccount    *Account `gorm:"foreignKey:ToAccountID;constraint:OnDelete:SET NULL" json:"-"`
	ToIdentifier *string  `json:"toIdentifier,omitempty"`
	ToAmount     int64    `gorm:"type:bigint;not null" json:"toAmount"` // Amount in smallest unit
	ToAmountMsat *int64   `gorm:"type:bigint" json:"toAmountMsat,omitempty"`
	ToCurrency   string   `gorm:"type:varchar(16);not null;default:'BTC';index" json:"toCurrency"`
	ToAmountUsd  *float64 `json:"toAmountUsd,omitempty"` // USD value at time of transaction

	// Fees
	Fee         int64    `gorm:"type:bigint;default:0" json:"fee"`
	FeeMsat     *int64   `gorm:"type:bigint" json:"feeMsat,omitempty"`
	FeeCurrency string   `gorm:"type:varchar(16);not null;default:'BTC'" json:"feeCurrency"`
	FeeUsd      *float64 `json:"feeUsd,omitempty"`

	// Exchange Rates
	ExchangeRate    *float64 `json:"exchangeRate,omitempty"`    // FromCurrency/ToCurrency rate
	ExchangeRateUsd *float64 `json:"exchangeRateUsd,omitempty"` // FromCurrency/USD rate

	// Lightning Network Details
	LightningInvoice     *string `gorm:"type:text" json:"lightningInvoice,omitempty"`
	LightningPaymentHash *string `gorm:"uniqueIndex:idx_event_wallet_lightning_hash,priority:1;size:64" json:"lightningPaymentHash,omitempty"`
	LightningPreimage    *string `gorm:"size:64" json:"lightningPreimage,omitempty"`

	// Provider References
	ProviderTransactionID *string `gorm:"index" json:"providerTransactionId,omitempty"`
	ProviderName          *string `gorm:"type:varchar(32)" json:"providerName,omitempty"` // IBEX, Stripe, etc.

	// Metadata
	Memo     *string         `gorm:"type:text" json:"memo,omitempty"`
	Metadata *datatypes.JSON `gorm:"type:jsonb" json:"metadata,omitempty"`

	// Timestamps
	CompletedAt *time.Time     `json:"completedAt,omitempty"`
	FailedAt    *time.Time     `json:"failedAt,omitempty"`
	CanceledAt  *time.Time     `json:"canceledAt,omitempty"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
