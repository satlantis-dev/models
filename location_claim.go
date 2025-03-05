package models

import (
	"time"
)

type LocationClaim struct {
	LocationID        uint        `gorm:"index;primaryKey" json:"locationId"`
	Location          LocationDTO `gorm:"foreignKey:LocationID" json:"location"`
	OwnerAccountID    uint        `gorm:"index" json:"ownerAccountId"`
	OwnerAccount      AccountDTO  `gorm:"foreignKey:OwnerAccountID" json:"ownerAccount"`
	BusinessAccountID *uint       `gorm:"index" json:"businessAccountId,omitempty"`
	BusinessAccount   *AccountDTO `gorm:"foreignKey:BusinessAccountID" json:"businessAccount,omitempty"`
	ClaimCode         string      `gorm:"type:text" json:"claimCode"`
	ClaimVerifiedAt   *time.Time  `json:"-"`
	ReferredBy        *string     `gorm:"type:text" json:"referredBy"`
}
