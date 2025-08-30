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
	AccountID  uint                    `gorm:"index;primaryKey" json:"accountId"`
	Account    *AccountDTO             `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"account,omitempty"`
	LocationID uint                    `gorm:"index;primaryKey" json:"locationId"`
	Location   *LocationDTO            `gorm:"foreignKey:LocationID;constraint:OnDelete:CASCADE" json:"location,omitempty"`
	Type       AccountLocationRoleType `gorm:"not null" json:"type"`
	CreatedAt  time.Time               `json:"-"`
	UpdatedAt  time.Time               `json:"-"`
	DeletedAt  gorm.DeletedAt          `gorm:"index" json:"-"`
}

func (AccountLocationRole) TableName() string {
	return "account_location_roles"
}
