package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountLocationRoleType int

const (
	Owner AccountLocationRoleType = iota + 1
	Admin
	Staff
	DiscoveredBy
)

type AccountLocationRole struct {
	AccountID  uint `gorm:"primaryKey" json:"accountId"`
	LocationID uint `gorm:"primaryKey" json:"locationId"`

	Type      AccountLocationRoleType `gorm:"not null" json:"type"`
	CreatedAt time.Time               `json:"-"`
	UpdatedAt time.Time               `json:"-"`
	DeletedAt gorm.DeletedAt          `gorm:"index" json:"-"`

	Account  *Account  `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"account,omitempty"`
	Location *Location `gorm:"foreignKey:LocationID;constraint:OnDelete:CASCADE" json:"location,omitempty"`
}

func (AccountLocationRole) TableName() string {
	return "account_location_roles"
}
