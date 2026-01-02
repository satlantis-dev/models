package models

import (
	"github.com/lib/pq"
	"time"
)

type OnboardingProfile struct {
	//AccountID           uint            `gorm:"primaryKey" json:"accountId"`
	//Account             *AccountDTO     `gorm:"foreignKey:AccountID;references:ID;constraint:OnDelete:CASCADE" json:"account,omitempty"`
	//CreatedAt           time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	//UpdatedAt           time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	//IsHost              *bool           `json:"isHost,omitempty"`
	//SuggestedGoogleIDs  *pq.StringArray `gorm:"type:text[]" json:"suggestedGoogleIds,omitempty"`
	//InterestStatement   *string         `gorm:"type:text" json:"interestStatement,omitempty"`
	//InferredInterests   *pq.StringArray `gorm:"type:text[]" json:"inferredInterests,omitempty"`
	//InferredInterestIDs *pq.Int32Array  `gorm:"type:integer[]" json:"inferredInterestIds,omitempty"`
	AccountID                        uint           `gorm:"primaryKey" json:"accountId"`
	Account                          *AccountDTO    `gorm:"foreignKey:AccountID;references:ID;constraint:OnDelete:CASCADE" json:"account,omitempty"`
	CreatedAt                        time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt                        time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	SuggestedLocationGoogleIDs       pq.StringArray `gorm:"type:text[]" json:"suggestedLocationGoogleIds"`
	InferredLocationTags             pq.StringArray `gorm:"type:text[]" json:"inferredLocationTags"`
	InferredInterestsFromLocations   pq.StringArray `gorm:"type:text[]" json:"inferredInterestsFromLocations"`
	InferredInterestIDsFromLocations pq.Int32Array  `gorm:"type:integer[]" json:"inferredInterestIdsFromLocations"`
	InterestStatement                string         `gorm:"type:text" json:"interestStatement"`
	InferredInterestsFromStatement   pq.StringArray `gorm:"type:text[]" json:"inferredInterestsFromStatement"`
	InferredInterestIDsFromStatement pq.Int32Array  `gorm:"type:integer[]" json:"inferredInterestIdsFromStatement"`
	Completed                        bool           `gorm:"default:false" json:"completed"`
}

func (OnboardingProfile) TableName() string {
	return "onboarding_profiles"
}

type OnboardingProfileInput struct {
	InterestStatement  *string   `json:"interestStatement,omitempty"`
	SuggestedGoogleIDs *[]string `json:"suggestedGoogleIds,omitempty"`
	IsHost             *bool     `json:"isHost"`
}
