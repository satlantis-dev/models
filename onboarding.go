package models

import (
	"time"

	"github.com/lib/pq"
)

type OnboardingProfile struct {
	AccountID                      uint           `gorm:"primaryKey" json:"accountId"`
	Account                        *AccountDTO    `gorm:"foreignKey:AccountID;references:ID;constraint:OnDelete:CASCADE" json:"account,omitempty"`
	CreatedAt                      time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt                      time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	SuggestedLocationGoogleIDs     pq.StringArray `gorm:"type:text[]" json:"suggestedLocationGoogleIds"`
	InferredLocationTags           pq.StringArray `gorm:"type:text[]" json:"inferredLocationTags"`
	InferredInterestsFromLocations pq.StringArray `gorm:"type:text[]" json:"inferredInterestsFromLocations"`
	InterestStatement              string         `gorm:"type:text" json:"interestStatement"`
	InferredInterestsFromStatement pq.StringArray `gorm:"type:text[]" json:"inferredInterestsFromStatement"`
	Completed                      bool           `gorm:"default:false" json:"completed"`
}

func (OnboardingProfile) TableName() string {
	return "onboarding_profiles"
}
