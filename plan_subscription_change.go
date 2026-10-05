package models

import (
	"time"

	"gorm.io/gorm"
)

type PlanSubscriptionChangeType string

const (
	PlanSubscriptionChangeTypePlanChange   PlanSubscriptionChangeType = "plan_change"
	PlanSubscriptionChangeTypePeriodChange PlanSubscriptionChangeType = "period_change"
)

type PlanSubscriptionChangeStatus string

const (
	PlanSubscriptionChangeStatusPending   PlanSubscriptionChangeStatus = "pending"
	PlanSubscriptionChangeStatusApplied   PlanSubscriptionChangeStatus = "applied"
	PlanSubscriptionChangeStatusCancelled PlanSubscriptionChangeStatus = "cancelled"
)

// PlanSubscriptionChange is a downgrade or period change scheduled on a plan
// subscription, applied when its current period ends. Mirrors
// CommunityMembershipSubscriptionChange.
type PlanSubscriptionChange struct {
	ID             uint                         `gorm:"primaryKey;autoIncrement" json:"id"`
	SubscriptionID uint                         `gorm:"not null;index;uniqueIndex:idx_plan_sub_change_one_pending,where:status = 'pending' AND deleted_at IS NULL" json:"subscriptionId"`
	Subscription   *PlanSubscription            `gorm:"foreignKey:SubscriptionID;constraint:OnDelete:CASCADE;" json:"subscription,omitempty"`
	AccountID      uint                         `gorm:"not null;index" json:"accountId"`
	ChangeType     PlanSubscriptionChangeType   `gorm:"type:varchar(32);not null;index" json:"changeType"`
	Status         PlanSubscriptionChangeStatus `gorm:"type:varchar(32);not null;default:'pending';index" json:"status"`
	OldPlanID      *uint                        `gorm:"index" json:"oldPlanId,omitempty"`
	OldPlan        *Plan                        `gorm:"foreignKey:OldPlanID;constraint:OnDelete:SET NULL;" json:"oldPlan,omitempty"`
	NewPlanID      *uint                        `gorm:"index" json:"newPlanId,omitempty"`
	NewPlan        *Plan                        `gorm:"foreignKey:NewPlanID;constraint:OnDelete:SET NULL;" json:"newPlan,omitempty"`
	OldPeriod      *PlanPeriod                  `gorm:"type:varchar(16)" json:"oldPeriod,omitempty"`
	NewPeriod      *PlanPeriod                  `gorm:"type:varchar(16)" json:"newPeriod,omitempty"`
	CreatedAt      time.Time                    `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time                    `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt      *gorm.DeletedAt              `gorm:"index" json:"-"`
}

func (PlanSubscriptionChange) TableName() string {
	return "plan_subscription_changes"
}
