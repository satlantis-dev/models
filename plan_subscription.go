package models

import (
	"time"

	"gorm.io/gorm"
)

type PlanPeriod string

const (
	PlanPeriodMonthly PlanPeriod = "monthly"
	PlanPeriodAnnual  PlanPeriod = "annual"
)

type PlanSubscriptionStatus string

const (
	PlanSubscriptionStatusActive         PlanSubscriptionStatus = "active"
	PlanSubscriptionStatusPastDue        PlanSubscriptionStatus = "past_due"
	PlanSubscriptionStatusCancelled      PlanSubscriptionStatus = "cancelled"
	PlanSubscriptionStatusPendingPayment PlanSubscriptionStatus = "pending_payment"
)

type PlanSubscriptionCancellationReason string

const (
	PlanSubscriptionCancellationReasonUserCancelled       PlanSubscriptionCancellationReason = "user_cancelled"
	PlanSubscriptionCancellationReasonSubscriptionExpired PlanSubscriptionCancellationReason = "subscription_expired"
	PlanSubscriptionCancellationReasonPlanChange          PlanSubscriptionCancellationReason = "plan_change"
)

type PlanSubscription struct {
	ID        uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID uint        `gorm:"not null;index;uniqueIndex:idx_plan_subscription_one_open_per_account_plan,where:status != 'cancelled' AND deleted_at IS NULL" json:"accountId"`
	Account   *AccountDTO `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	PlanID    uint        `gorm:"not null;index;uniqueIndex:idx_plan_subscription_one_open_per_account_plan,where:status != 'cancelled' AND deleted_at IS NULL" json:"planId"`

	Period        PlanPeriod             `gorm:"type:varchar(16);not null" json:"period"`
	Amount        *int64                 `gorm:"type:bigint;check:chk_plan_subscription_amount_currency,((amount IS NULL AND currency IS NULL) OR (amount IS NOT NULL AND currency IS NOT NULL))" json:"amount,omitempty"`
	Currency      *OrderCurrency         `gorm:"type:varchar(8)" json:"currency,omitempty"`
	Status        PlanSubscriptionStatus `gorm:"type:varchar(32);not null;default:'active';index" json:"status"`
	PaymentMethod *PaymentMethod         `gorm:"type:varchar(32)" json:"paymentMethod,omitempty"`

	CurrentPeriodStart *time.Time `gorm:"type:timestamptz" json:"currentPeriodStart,omitempty"`
	CurrentPeriodEnd   *time.Time `gorm:"type:timestamptz;index" json:"currentPeriodEnd,omitempty"`
	PastDueSince       *time.Time `gorm:"type:timestamptz;index" json:"pastDueSince,omitempty"`
	EndedAt            *time.Time `gorm:"type:timestamptz" json:"endedAt,omitempty"`

	CancelAtPeriodEnd  bool                                `gorm:"not null;default:false" json:"cancelAtPeriodEnd"`
	CancelledAt        *time.Time                          `gorm:"type:timestamptz" json:"cancelledAt,omitempty"`
	CancellationReason *PlanSubscriptionCancellationReason `gorm:"type:varchar(256)" json:"cancellationReason,omitempty"`

	CreatedAt time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`
}

func (PlanSubscription) TableName() string {
	return "plan_subscriptions"
}
