package models

import (
	"time"
)

type LocationAccount struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	LocationID      uint       `gorm:"index" json:"locationId"`
	Location        Location   `json:"location"`
	AccountID       uint       `gorm:"index" json:"accountId"`
	Account         AccountDTO `json:"account"`
	ClaimCode       string     `gorm:"type:text" json:"claimCode"`
	ClaimVerifiedAt *time.Time `json:"-"`
	ReferredBy      *string    `gorm:"type:text" json:"referredBy"`
}
