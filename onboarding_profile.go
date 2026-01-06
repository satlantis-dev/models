package models

import (
	"time"

	"github.com/lib/pq"
)

type OnboardingProfile struct {
	AccountID           uint            `gorm:"primaryKey" json:"accountId"`
	Account             *AccountDTO     `gorm:"foreignKey:AccountID;references:ID;constraint:OnDelete:CASCADE" json:"account,omitempty"`
	CreatedAt           time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt           time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	IsHost              *bool           `json:"isHost,omitempty"`
	SuggestedGoogleIDs  *pq.StringArray `gorm:"type:text[]" json:"suggestedGoogleIds,omitempty"`
	InterestStatement   *string         `gorm:"type:text" json:"interestStatement,omitempty"`
	InferredInterests   *pq.StringArray `gorm:"type:text[]" json:"inferredInterests,omitempty"`
	InferredInterestIDs *pq.Int32Array  `gorm:"type:integer[]" json:"inferredInterestIds,omitempty"`
}

func (OnboardingProfile) TableName() string {
	return "onboarding_profiles"
}

type OnboardingProfileInput struct {
	InterestStatement  *string   `json:"interestStatement,omitempty"`
	SuggestedGoogleIDs *[]string `json:"suggestedGoogleIds,omitempty"`
	IsHost             *bool     `json:"isHost"`
}
