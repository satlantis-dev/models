package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// AccountWalletTransactionType represents the type of wallet transaction
type AccountWalletTransactionType string

const (
	WalletTransactionSend    AccountWalletTransactionType = "send"    // Send to external address
	WalletTransactionReceive AccountWalletTransactionType = "receive" // Receive from external
)

// AccountWalletTransactionStatus represents the status of a wallet transaction
type AccountWalletTransactionStatus string

const (
	WalletTransactionPending   AccountWalletTransactionStatus = "pending"
	WalletTransactionCompleted AccountWalletTransactionStatus = "completed"
	WalletTransactionFailed    AccountWalletTransactionStatus = "failed"
	WalletTransactionCancelled AccountWalletTransactionStatus = "cancelled"
)

// AccountWallet stores user wallet information
type AccountWallet struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	AccountID        uint           `gorm:"not null;uniqueIndex" json:"accountId"`
	Account          *Account       `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"-"`
	Provider         WalletProvider `gorm:"type:varchar(32);not null;default:'IBEX'" json:"provider"` // Wallet provider
	ProviderWalletId *string        `json:"providerWalletId"`                                         // Provider's wallet/account ID
	LightningAddress *string        `gorm:"uniqueIndex" json:"lightningAddress,omitempty"`            // username@satlantis.io
	Balance          int64          `gorm:"type:bigint;default:0" json:"balance"`                     // Balance in sats
	BalanceMsat      *int64         `gorm:"type:bigint" json:"balanceMsat,omitempty"`
	Currency         OrderCurrency  `gorm:"type:varchar(8);default:'BTC'" json:"currency"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

// AccountWalletTransaction stores all wallet transactions
type AccountWalletTransaction struct {
	ID                        uint                           `gorm:"primaryKey" json:"id"`
	AccountWalletID           uint                           `gorm:"not null;index" json:"accountWalletId"`
	AccountWallet             *AccountWallet                 `gorm:"foreignKey:AccountWalletID;constraint:OnDelete:CASCADE" json:"-"`
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
	LightningPaymentHash *string `gorm:"uniqueIndex:idx_lightning_hash_type,priority:1;size:64" json:"lightningPaymentHash,omitempty"`
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
