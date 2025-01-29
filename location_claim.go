package models

import (
	"time"
)

type LocationClaim struct {
	LocationID      uint        `gorm:"index;primaryKey" json:"locationId"`
	Location        LocationDTO `json:"location"`
	AccountID       uint        `gorm:"index" json:"accountId"`
	Account         AccountDTO  `json:"account"`
	ClaimCode       string      `gorm:"type:text" json:"claimCode"`
	ClaimVerifiedAt *time.Time  `json:"-"`
	ReferredBy      *string     `gorm:"type:text" json:"referredBy"`
}
