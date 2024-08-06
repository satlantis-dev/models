package database

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
	LocationID      uint                `gorm:"primaryKey" json:"locationId"`
	AccountID       uint                `gorm:"primaryKey" json:"accountId"`
	Type            LocationAccountType `gorm:"not null" json:"type"`
	ClaimCode       string              `gorm:"type:text" json:"claimCode"`
	ClaimVerifiedAt *time.Time          `json:"-"`
}
