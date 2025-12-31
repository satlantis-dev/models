package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// StripeAccountStatus represents the status of a connected Stripe account
type StripeAccountStatus string

const (
	StripeAccountPending    StripeAccountStatus = "pending"
	StripeAccountActive     StripeAccountStatus = "active"
	StripeAccountRestricted StripeAccountStatus = "restricted"
	StripeAccountDisabled   StripeAccountStatus = "disabled"
)

// AccountStripeConnect stores Stripe Connect account information linked to a user account
type AccountStripeConnect struct {
	ID               uint                `gorm:"primaryKey" json:"id"`
	AccountID        uint                `gorm:"not null;index" json:"accountId"`
	Account          *Account            `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"-"`
	StripeAccountID  string              `gorm:"uniqueIndex;not null;size:64" json:"stripeAccountId"` // acct_xxx
	IsDefault        bool                `gorm:"default:false" json:"isDefault"`
	Status           StripeAccountStatus `gorm:"type:varchar(32);default:'pending'" json:"status"`
	ChargesEnabled   bool                `gorm:"default:false" json:"chargesEnabled"`
	PayoutsEnabled   bool                `gorm:"default:false" json:"payoutsEnabled"`
	DetailsSubmitted bool                `gorm:"default:false" json:"detailsSubmitted"`
	DefaultCurrency  *string             `gorm:"size:8" json:"defaultCurrency,omitempty"`
	Country          *string             `gorm:"size:8" json:"country,omitempty"`
	BusinessType     *string             `gorm:"size:32" json:"businessType,omitempty"`
	Email            *string             `json:"email,omitempty"`
	Metadata         *datatypes.JSON     `gorm:"type:jsonb" json:"metadata,omitempty"`
	ConnectedAt      time.Time           `json:"connectedAt"`
	CreatedAt        time.Time           `json:"createdAt"`
	UpdatedAt        time.Time           `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt      `gorm:"index" json:"-"`
}

// AccountStripeConnectDTO represents the public-facing DTO for Stripe Connect status
type AccountStripeConnectDTO struct {
	ID               uint                `json:"id"`
	StripeAccountID  string              `json:"stripeAccountId"`
	Status           StripeAccountStatus `json:"status"`
	ChargesEnabled   bool                `json:"chargesEnabled"`
	PayoutsEnabled   bool                `json:"payoutsEnabled"`
	DetailsSubmitted bool                `json:"detailsSubmitted"`
	DefaultCurrency  *string             `json:"defaultCurrency,omitempty"`
	Country          *string             `json:"country,omitempty"`
	Email            *string             `json:"email,omitempty"`
	IsDefault        bool                `json:"isDefault"`
	ConnectedAt      time.Time           `json:"connectedAt"`
}

// ToDTO converts AccountStripeConnect to its DTO representation
func (a *AccountStripeConnect) ToDTO() AccountStripeConnectDTO {
	return AccountStripeConnectDTO{
		ID:               a.ID,
		StripeAccountID:  a.StripeAccountID,
		Status:           a.Status,
		ChargesEnabled:   a.ChargesEnabled,
		PayoutsEnabled:   a.PayoutsEnabled,
		DetailsSubmitted: a.DetailsSubmitted,
		DefaultCurrency:  a.DefaultCurrency,
		Country:          a.Country,
		Email:            a.Email,
		IsDefault:        a.IsDefault,
		ConnectedAt:      a.ConnectedAt,
	}
}
