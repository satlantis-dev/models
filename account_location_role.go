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
	DisoveredBy
	RecommendedBy
	BookmarkedBy
)

type AccountLocationRole struct {
	AccountID  uint                    `gorm:"index;primaryKey" json:"accountId"`
	Account    AccountDTO              `gorm:"foreignKey:AccountID" json:"account"`
	LocationID uint                    `gorm:"index;primaryKey" json:"locationId"`
	Location   *Location               `gorm:"foreignKey:LocationID" json:"location,omitempty"`
	Type       AccountLocationRoleType `gorm:"not null;primaryKey" json:"type"`
	CreatedAt  time.Time               `json:"-"`
	UpdatedAt  time.Time               `json:"-"`
	DeletedAt  gorm.DeletedAt          `gorm:"index" json:"-,omitempty"`
}
