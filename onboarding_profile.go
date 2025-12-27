package models

import (
	"time"

	"github.com/lib/pq"
)

type OnboardingProfile struct {
	AccountID                  uint            `gorm:"primaryKey" json:"accountId"`
	Account                    *AccountDTO     `gorm:"foreignKey:AccountID;references:ID;constraint:OnDelete:CASCADE" json:"account,omitempty"`
	CreatedAt                  time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt                  time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	IsHost                     *bool           `json:"isHost,omitempty"`
	SuggestedLocationGoogleIDs *pq.StringArray `gorm:"type:text[]" json:"suggestedLocationGoogleIds,omitempty"`
	InterestStatement          *string         `gorm:"type:text" json:"interestStatement,omitempty"`
	InferredInterests          *pq.StringArray `gorm:"type:text[]" json:"inferredInterests,omitempty"`
	InferredInterestIDs        *pq.Int32Array  `gorm:"type:integer[]" json:"inferredInterestIds,omitempty"`
}

func (OnboardingProfile) TableName() string {
	return "onboarding_profiles"
}
