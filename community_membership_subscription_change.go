package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CommunityMembershipSubscriptionChangeType string

const (
	CommunityMembershipSubscriptionChangeTypeTierChange   CommunityMembershipSubscriptionChangeType = "tier_change"
	CommunityMembershipSubscriptionChangeTypePeriodChange CommunityMembershipSubscriptionChangeType = "period_change"
	CommunityMembershipSubscriptionChangeTypeCancel       CommunityMembershipSubscriptionChangeType = "cancel"
)

type CommunityMembershipSubscriptionChangeStatus string

const (
	CommunityMembershipSubscriptionChangeStatusPending   CommunityMembershipSubscriptionChangeStatus = "pending"
	CommunityMembershipSubscriptionChangeStatusApplied   CommunityMembershipSubscriptionChangeStatus = "applied"
	CommunityMembershipSubscriptionChangeStatusCancelled CommunityMembershipSubscriptionChangeStatus = "cancelled"
	CommunityMembershipSubscriptionChangeStatusFailed    CommunityMembershipSubscriptionChangeStatus = "failed"
)

type CommunityMembershipSubscriptionChange struct {
	ID                 uint                                        `gorm:"primaryKey;autoIncrement" json:"id"`
	SubscriptionID     uint                                        `gorm:"not null;index;uniqueIndex:idx_community_membership_sub_change_one_pending,where:status = 'pending' AND deleted_at IS NULL" json:"subscriptionId"`
	Subscription       *CommunityMembershipSubscription            `gorm:"foreignKey:SubscriptionID;constraint:OnDelete:CASCADE;" json:"subscription,omitempty"`
	AccountID          uint                                        `gorm:"not null;index" json:"accountId"`
	CommunityID        uint                                        `gorm:"not null;index" json:"communityId"`
	ChangeType         CommunityMembershipSubscriptionChangeType   `gorm:"type:varchar(32);not null;index" json:"changeType"`
	Status             CommunityMembershipSubscriptionChangeStatus `gorm:"type:varchar(32);not null;default:'pending';index" json:"status"`
	OldTierID          *uint                                       `gorm:"index" json:"oldTierId,omitempty"`
	OldTier            *CommunityMembershipTier                    `gorm:"foreignKey:OldTierID;constraint:OnDelete:SET NULL;" json:"oldTier,omitempty"`
	NewTierID          *uint                                       `gorm:"index" json:"newTierId,omitempty"`
	NewTier            *CommunityMembershipTier                    `gorm:"foreignKey:NewTierID;constraint:OnDelete:SET NULL;" json:"newTier,omitempty"`
	OldPeriod          *CommunityMembershipPeriod                  `gorm:"type:varchar(16)" json:"oldPeriod,omitempty"`
	NewPeriod          *CommunityMembershipPeriod                  `gorm:"type:varchar(16)" json:"newPeriod,omitempty"`
	OldAmount          *int64                                      `gorm:"type:bigint" json:"oldAmount,omitempty"`
	OldCurrency        *OrderCurrency                              `gorm:"type:varchar(8)" json:"oldCurrency,omitempty"`
	NewAmount          *int64                                      `gorm:"type:bigint" json:"newAmount,omitempty"`
	NewCurrency        *OrderCurrency                              `gorm:"type:varchar(8)" json:"newCurrency,omitempty"`
	EffectiveAt        time.Time                                   `gorm:"type:timestamptz;not null;index" json:"effectiveAt"`
	AppliedAt          *time.Time                                  `gorm:"type:timestamptz" json:"appliedAt,omitempty"`
	CreatedByAccountID *uint                                       `gorm:"index" json:"createdByAccountId,omitempty"`
	Metadata           *datatypes.JSON                             `gorm:"type:jsonb" json:"metadata,omitempty"`
	CreatedAt          time.Time                                   `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt          time.Time                                   `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt          *gorm.DeletedAt                             `gorm:"index" json:"-"`
}

func (CommunityMembershipSubscriptionChange) TableName() string {
	return "community_membership_subscription_changes"
}
