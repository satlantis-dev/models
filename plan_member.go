package models

import (
	"time"
)

// PlanMember tracks which non-default Plan an account is currently on.
// Accounts with no PlanMember row are on the default Starter plan - that
// case is never persisted here.
type PlanMember struct {
	AccountID             uint                `gorm:"primaryKey" json:"accountId"`
	Account               *AccountDTO         `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	PlanID                uint                `gorm:"not null;index" json:"planId"`
	Plan                  *Plan               `gorm:"foreignKey:PlanID;constraint:OnDelete:RESTRICT;" json:"plan,omitempty"`
	StartDate             *time.Time          `gorm:"type:timestamptz" json:"startDate,omitempty"`
	ExpiryDate            *time.Time          `gorm:"type:timestamptz;index" json:"expiryDate,omitempty"`
	CreatedAt             time.Time           `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt             time.Time           `gorm:"autoUpdateTime" json:"updatedAt"`
	OpenPlanSubscriptions *[]PlanSubscription `gorm:"-" json:"openPlanSubscriptions,omitempty"`
}

func (PlanMember) TableName() string {
	return "plan_members"
}
