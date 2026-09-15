package models

import (
	"time"
)

// AccountPlan tracks which non-default Plan an account is currently on.
// Accounts with no AccountPlan row are on the default Starter plan - that
// case is never persisted here.
type AccountPlan struct {
	AccountID             uint                `gorm:"primaryKey" json:"accountId"`
	PlanID                uint                `gorm:"not null;index" json:"planId"`
	Plan                  *Plan               `gorm:"foreignKey:PlanID;constraint:OnDelete:RESTRICT;" json:"plan,omitempty"`
	StartDate             *time.Time          `gorm:"type:timestamptz" json:"startDate,omitempty"`
	ExpiryDate            *time.Time          `gorm:"type:timestamptz;index" json:"expiryDate,omitempty"`
	CreatedAt             time.Time           `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt             time.Time           `gorm:"autoUpdateTime" json:"updatedAt"`
	OpenPlanSubscriptions *[]PlanSubscription `gorm:"-" json:"openPlanSubscriptions,omitempty"`
}

func (AccountPlan) TableName() string {
	return "account_plans"
}
