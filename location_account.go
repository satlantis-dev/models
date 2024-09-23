package models

import (
	"time"
)

type LocationAccountType string

const (
	LocationAccountTypeOwner   LocationAccountType = "owner"
	LocationAccountTypeManager LocationAccountType = "manager"
	LocationAccountTypeMember  LocationAccountType = "member"
	LocationAccountTypeVisitor LocationAccountType = "visitor"
)

type LocationAccount struct {
	ID              uint                `gorm:"primaryKey" json:"id"`
	LocationID      uint                `gorm:"index" json:"locationId"`
	Location        Location            `json:"location"`
	AccountID       uint                `gorm:"index" json:"accountId"`
	Type            LocationAccountType `gorm:"not null" json:"type"`
	ClaimCode       string              `gorm:"type:text" json:"claimCode"`
	ClaimVerifiedAt *time.Time          `json:"-"`
	ReferredBy      *string             `gorm:"type:text" json:"referredBy"`
}
