package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// CommunityMembershipRefund represents a refund against a membership payment.
type CommunityMembershipRefund struct {
	ID        uint                        `gorm:"primarykey" json:"id"`
	PaymentID uint                        `gorm:"not null;index" json:"paymentId"`
	Payment   *CommunityMembershipPayment `gorm:"foreignKey:PaymentID;constraint:OnDelete:CASCADE" json:"payment,omitempty"`
	Amount    int64                       `gorm:"not null" json:"amount"`
	Fee       int64                       `gorm:"type:bigint;default:0" json:"fee"`
	Currency  OrderCurrency               `gorm:"not null" json:"currency"`
	Status    RefundStatus                `gorm:"not null;default:'pending'" json:"status"`

	// Supported values include "lightning" and "stripe".
	RefundMethod string `gorm:"not null" json:"refundMethod"`

	LightningAddress     *string `json:"lightningAddress,omitempty"`
	LightningPaymentHash *string `json:"lightningPaymentHash,omitempty"`
	LightningPreimage    *string `gorm:"size:64" json:"lightningPreimage,omitempty"`

	StripeRefundID        *string `json:"stripeRefundId,omitempty"`
	StripePaymentIntentID *string `json:"stripePaymentIntentId,omitempty"`

	Reason        *string         `gorm:"type:text" json:"reason,omitempty"`
	CreatedAt     time.Time       `json:"createdAt"`
	ProcessedAt   *time.Time      `json:"processedAt,omitempty"`
	FailedAt      *time.Time      `json:"failedAt,omitempty"`
	FailureReason *string         `gorm:"type:text" json:"failureReason,omitempty"`
	Metadata      *datatypes.JSON `gorm:"type:jsonb" json:"metadata,omitempty"`
	DeletedAt     gorm.DeletedAt  `gorm:"index" json:"-"`
}

func (CommunityMembershipRefund) TableName() string {
	return "community_membership_refunds"
}
