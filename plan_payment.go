package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// PlanPayment tracks each charge attempt (Stripe or Lightning) for a
// PlanSubscription.
type PlanPayment struct {
	ID                 uint              `gorm:"primaryKey;autoIncrement" json:"id"`
	SubscriptionID     uint              `gorm:"not null;index" json:"subscriptionId"`
	Subscription       *PlanSubscription `gorm:"foreignKey:SubscriptionID;constraint:OnDelete:CASCADE" json:"subscription,omitempty"`
	PaymentMethod      PaymentMethod     `gorm:"type:varchar(32);not null" json:"paymentMethod"`
	Status             PaymentStatus     `gorm:"type:varchar(32);not null;default:'pending';index" json:"status"`
	Amount             int64             `gorm:"not null" json:"amount"`
	Currency           OrderCurrency     `gorm:"type:varchar(8);not null" json:"currency"`
	BillingPeriodStart *time.Time        `gorm:"type:timestamptz" json:"billingPeriodStart,omitempty"`
	BillingPeriodEnd   *time.Time        `gorm:"type:timestamptz;index" json:"billingPeriodEnd,omitempty"`
	Metadata           *datatypes.JSON   `gorm:"type:jsonb" json:"metadata,omitempty"`
	PaidAt             *time.Time        `gorm:"type:timestamptz" json:"paidAt,omitempty"`
	ExpiresAt          *time.Time        `gorm:"type:timestamptz" json:"expiresAt,omitempty"`
	FailedAt           *time.Time        `gorm:"type:timestamptz" json:"failedAt,omitempty"`

	// Stripe fields
	StripePaymentIntentID *string `gorm:"type:varchar(128);uniqueIndex" json:"stripePaymentIntentId,omitempty"`
	StripeChargeID        *string `gorm:"type:varchar(128);index" json:"stripeChargeId,omitempty"`

	// Lightning fields
	LightningPaymentHash    *string `gorm:"uniqueIndex;size:64" json:"lightningPaymentHash,omitempty"`
	LightningPaymentRequest *string `gorm:"type:text" json:"lightningPaymentRequest,omitempty"`
	LightningPreimage       *string `gorm:"size:64" json:"lightningPreimage,omitempty"`
	LightningProvider       *string `gorm:"size:32" json:"lightningProvider,omitempty"`
	LightningProviderTxID   *string `gorm:"index" json:"lightningProviderTxId,omitempty"`
	LightningAddress        *string `gorm:"type:text" json:"lightningAddress,omitempty"`

	CreatedAt time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`
}

func (PlanPayment) TableName() string {
	return "plan_payments"
}
