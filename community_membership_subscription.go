package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CommunityMembershipSubscriptionStatus string

const (
	CommunityMembershipSubscriptionStatusActive          CommunityMembershipSubscriptionStatus = "active"
	CommunityMembershipSubscriptionStatusPastDue         CommunityMembershipSubscriptionStatus = "past_due"
	CommunityMembershipSubscriptionStatusCancelled       CommunityMembershipSubscriptionStatus = "cancelled"
	CommunityMembershipSubscriptionStatusPendingApproval CommunityMembershipSubscriptionStatus = "pending_approval"
	CommunityMembershipSubscriptionStatusPendingPayment  CommunityMembershipSubscriptionStatus = "pending_payment"
)

type CommunityMembershipSubscriptionCancellationReason string

const (
	CommunityMembershipSubscriptionCancellationReasonRequestCancelled    CommunityMembershipSubscriptionCancellationReason = "request_cancelled"
	CommunityMembershipSubscriptionCancellationReasonRequestRejected     CommunityMembershipSubscriptionCancellationReason = "request_rejected"
	CommunityMembershipSubscriptionCancellationReasonSubscriptionExpired CommunityMembershipSubscriptionCancellationReason = "subscription_expired"
	CommunityMembershipSubscriptionCancellationReasonMemberRemoved       CommunityMembershipSubscriptionCancellationReason = "member_removed"
)

type CommunityMembershipSubscription struct {
	ID                     uint                                               `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID              uint                                               `gorm:"not null;index;uniqueIndex:idx_cms_one_active_per_account_community,where:(status = 'active' OR status = 'past_due') AND deleted_at IS NULL;uniqueIndex:idx_community_membership_subscription_one_open_per_account_tier,where:status != 'cancelled' AND deleted_at IS NULL" json:"accountId"`
	Account                *AccountDTO                                        `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	CommunityID            uint                                               `gorm:"not null;index;uniqueIndex:idx_cms_one_active_per_account_community,where:(status = 'active' OR status = 'past_due') AND deleted_at IS NULL" json:"communityId"`
	Community              *Community                                         `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"community,omitempty"`
	MemberID               *uint                                              `gorm:"index" json:"memberId,omitempty"`
	Member                 *CommunityMember                                   `gorm:"foreignKey:MemberID;constraint:OnDelete:SET NULL;" json:"member,omitempty"`
	TierID                 uint                                               `gorm:"not null;index;uniqueIndex:idx_community_membership_subscription_one_open_per_account_tier,where:status != 'cancelled' AND deleted_at IS NULL" json:"tierId"`
	Tier                   *CommunityMembershipTier                           `gorm:"foreignKey:TierID;constraint:OnDelete:RESTRICT;" json:"tier,omitempty"`
	Period                 CommunityMembershipPeriod                          `gorm:"type:varchar(16);not null" json:"period"`
	Amount                 *int64                                             `gorm:"type:bigint;check:chk_community_membership_subscription_amount_currency,((amount IS NULL AND currency IS NULL) OR (amount IS NOT NULL AND currency IS NOT NULL))" json:"amount,omitempty"`
	Currency               *OrderCurrency                                     `gorm:"type:varchar(8)" json:"currency,omitempty"`
	Status                 CommunityMembershipSubscriptionStatus              `gorm:"type:varchar(32);not null;default:'active';index" json:"status"`
	RequestID              *uint                                              `gorm:"index" json:"requestId,omitempty"`
	Request                *CommunityMembershipRequest                        `gorm:"foreignKey:RequestID;constraint:OnDelete:SET NULL;" json:"request,omitempty"`
	PaymentMethod          *PaymentMethod                                     `gorm:"type:varchar(32)" json:"paymentMethod,omitempty"`
	StripeCustomerID       *string                                            `gorm:"type:varchar(128);index" json:"stripeCustomerId,omitempty"`
	StripePaymentMethodID  *string                                            `gorm:"type:varchar(128);index" json:"stripePaymentMethodId,omitempty"`
	StripeSetupIntentID    *string                                            `gorm:"type:varchar(128);index" json:"stripeSetupIntentId,omitempty"`
	ProviderSubscriptionID *string                                            `gorm:"type:varchar(128);uniqueIndex" json:"providerSubscriptionId,omitempty"`
	CurrentPeriodStart     *time.Time                                         `gorm:"type:timestamptz" json:"currentPeriodStart,omitempty"`
	CurrentPeriodEnd       *time.Time                                         `gorm:"type:timestamptz;index" json:"currentPeriodEnd,omitempty"`
	CancelAtPeriodEnd      bool                                               `gorm:"not null;default:false" json:"cancelAtPeriodEnd"`
	CancelledAt            *time.Time                                         `gorm:"type:timestamptz" json:"cancelledAt,omitempty"`
	CancellationReason     *CommunityMembershipSubscriptionCancellationReason `gorm:"type:varchar(256)" json:"cancellationReason,omitempty"`
	EndedAt                *time.Time                                         `gorm:"type:timestamptz" json:"endedAt,omitempty"`
	Metadata               *datatypes.JSON                                    `gorm:"type:jsonb" json:"metadata,omitempty"`
	Payments               []CommunityMembershipPayment                       `gorm:"foreignKey:SubscriptionID" json:"payments,omitempty"`
	ScheduledChanges       []CommunityMembershipSubscriptionChange            `gorm:"foreignKey:SubscriptionID" json:"scheduledChanges,omitempty"`
	CreatedAt              time.Time                                          `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt              time.Time                                          `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt              *gorm.DeletedAt                                    `gorm:"index" json:"-"`
}

func (CommunityMembershipSubscription) TableName() string {
	return "community_membership_subscriptions"
}
