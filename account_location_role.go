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
	BookmarkedBy
	LikedBy
	DislikedBy
)

type AccountLocationRole struct {
	AccountID  uint                    `gorm:"index;primaryKey" json:"accountId"`
	Account    AccountDTO              `gorm:"foreignKey:AccountID" json:"account"`
	LocationID uint                    `gorm:"index;primaryKey" json:"locationId"`
	Location   *LocationDTO            `gorm:"foreignKey:LocationID" json:"location,omitempty"`
	Type       AccountLocationRoleType `gorm:"not null;primaryKey" json:"type"`
	CreatedAt  time.Time               `json:"-"`
	UpdatedAt  time.Time               `json:"-"`
	DeletedAt  gorm.DeletedAt          `gorm:"index" json:"-,omitempty"`
}

type UserReview struct {
	AccountID  uint         `gorm:"primaryKey;index" json:"accountId"`
	Account    AccountDTO   `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	LocationID uint         `gorm:"primaryKey;index" json:"locationId"`
	Location   *LocationDTO `gorm:"foreignKey:LocationID" json:"location,omitempty"`
	ReviewText string       `json:"reviewText"`
	IsPositive bool         `gorm:"not null;default:true" json:"isPositive"`
	IsPublic   bool         `gorm:"not null;default:true" json:"isPublic"`
}
