package models

import (
	"time"

	"github.com/lib/pq"
)

type OnboardingProfile struct {
	AccountID                  uint           `gorm:"primaryKey" json:"accountId"`
	Account                    *AccountDTO    `gorm:"foreignKey:AccountID;references:ID;constraint:OnDelete:CASCADE" json:"account,omitempty"`
	CreatedAt                  time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt                  time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	Host                       bool           `gorm:"default:false" json:"host"`
	SuggestedLocationGoogleIDs pq.StringArray `gorm:"type:text[]" json:"suggestedLocationGoogleIds"`
	InterestStatement          string         `gorm:"type:text" json:"interestStatement"`
	InferredInterests          pq.StringArray `gorm:"type:text[]" json:"inferredInterests"`
	InferredInterestIDs        pq.Int32Array  `gorm:"type:integer[]" json:"inferredInterestIds"`
}

func (OnboardingProfile) TableName() string {
	return "onboarding_profiles"
}
