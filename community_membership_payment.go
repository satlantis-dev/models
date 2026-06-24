package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CommunityMembershipPayment struct {
	ID                       uint                             `gorm:"primaryKey" json:"id"`
	SubscriptionID           uint                             `gorm:"not null;index" json:"subscriptionId"`
	Subscription             *CommunityMembershipSubscription `gorm:"foreignKey:SubscriptionID;constraint:OnDelete:CASCADE" json:"subscription,omitempty"`
	Refunds                  []CommunityMembershipRefund      `gorm:"foreignKey:PaymentID" json:"refunds,omitempty"`
	PaymentMethod            PaymentMethod                    `gorm:"type:varchar(32);not null" json:"paymentMethod"`
	Status                   PaymentStatus                    `gorm:"type:varchar(32);default:'pending';index" json:"status"`
	Amount                   int64                            `gorm:"not null" json:"amount"`
	Currency                 OrderCurrency                    `gorm:"type:varchar(8);not null" json:"currency"`
	BillingPeriodStart       *time.Time                       `gorm:"type:timestamptz" json:"billingPeriodStart,omitempty"`
	BillingPeriodEnd         *time.Time                       `gorm:"type:timestamptz;index" json:"billingPeriodEnd,omitempty"`
	ExchangeRate             *float64                         `json:"exchangeRate"`
	ExchangeRateSource       *string                          `gorm:"type:varchar(50)" json:"exchangeRateSource"`
	LightningPaymentHash     *string                          `gorm:"uniqueIndex;size:64" json:"lightningPaymentHash,omitempty"`
	LightningPaymentRequest  *string                          `gorm:"type:text" json:"lightningPaymentRequest,omitempty"`
	LightningPreimage        *string                          `gorm:"size:64" json:"lightningPreimage,omitempty"`
	LightningProvider        *string                          `gorm:"size:32" json:"lightningProvider,omitempty"`
	LightningProviderTxID    *string                          `gorm:"index" json:"lightningProviderTxId,omitempty"`
	StripePaymentIntentID    *string                          `gorm:"type:varchar(128);uniqueIndex" json:"stripePaymentIntentId,omitempty"`
	StripeInvoiceID          *string                          `gorm:"type:varchar(128);index" json:"stripeInvoiceId,omitempty"`
	StripeChargeID           *string                          `gorm:"type:varchar(128);index" json:"stripeChargeId,omitempty"`
	PaymentProviderReference *string                          `gorm:"index" json:"paymentProviderReference"`
	CardLast4                *string                          `gorm:"type:varchar(4)" json:"cardLast4,omitempty"`
	CardBrand                *string                          `gorm:"type:varchar(32)" json:"cardBrand,omitempty"`
	Metadata                 *datatypes.JSON                  `gorm:"type:jsonb" json:"metadata,omitempty"`
	PaidAt                   *time.Time                       `json:"paidAt,omitempty"`
	ExpiredAt                *time.Time                       `json:"expiredAt,omitempty"`
	ExpiresAt                *time.Time                       `json:"expiresAt,omitempty"`
	FailedAt                 *time.Time                       `json:"failedAt,omitempty"`
	RefundedAt               *time.Time                       `json:"refundedAt,omitempty"`
	CancelledAt              *time.Time                       `json:"cancelledAt,omitempty"`
	CreatedAt                time.Time                        `json:"createdAt"`
	UpdatedAt                time.Time                        `json:"updatedAt"`
	DeletedAt                gorm.DeletedAt                   `gorm:"index" json:"-"`
}

func (CommunityMembershipPayment) TableName() string {
	return "community_membership_payments"
}
