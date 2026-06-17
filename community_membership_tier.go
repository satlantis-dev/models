package models

import (
	"time"

	"gorm.io/gorm"
)

type CommunityMembershipTier struct {
	ID                    uint                          `gorm:"primaryKey;autoIncrement" json:"id"`
	CommunityID           uint                          `gorm:"not null;uniqueIndex:idx_community_tier_name" json:"communityId"`
	Community             *Community                    `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"community,omitempty"`
	Name                  string                        `gorm:"type:text;not null;uniqueIndex:idx_community_tier_name" json:"name"`
	Blurb                 *string                       `gorm:"type:varchar(100)" json:"blurb,omitempty"`
	Description           *string                       `gorm:"type:text" json:"description,omitempty"`
	ButtonText            *string                       `gorm:"type:varchar(20)" json:"buttonText,omitempty"`
	RegistrationQuestions *RegistrationQuestionsPayload `gorm:"type:jsonb;serializer:json" json:"registrationQuestions,omitempty"`
	IsGated               bool                          `gorm:"not null;default:false" json:"isGated"`
	IsPaid                bool                          `gorm:"not null;default:false" json:"isPaid"`
	IsRecommended         bool                          `gorm:"not null;default:false" json:"isRecommended"`
	Currency              *OrderCurrency                `gorm:"-" json:"currency,omitempty"`
	MonthlyAmount         *int64                        `gorm:"type:bigint" json:"monthlyAmount,omitempty"`
	YearlyAmount          *int64                        `gorm:"type:bigint" json:"yearlyAmount,omitempty"`
	Rank                  int                           `gorm:"not null;default:0" json:"rank"`
	TrialDays             int                           `gorm:"not null;default:0" json:"trialDays"`
	CreatedAt             time.Time                     `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt             time.Time                     `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt             *gorm.DeletedAt               `gorm:"index" json:"-"`
}

func (CommunityMembershipTier) TableName() string {
	return "community_membership_tiers"
}
