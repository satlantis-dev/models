package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountLocationReview struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID  uint           `gorm:"uniqueIndex:idx_account_location" json:"accountId"`
	Account    AccountDTO     `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	GoogleID   string         `gorm:"uniqueIndex:idx_account_location" json:"googleId"`
	ReviewText *string        `json:"reviewText"`
	IsPositive bool           `gorm:"not null;default:true" json:"isPositive"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	PriceLevel PriceLevel     `gorm:"default:PRICE_LEVEL_UNSPECIFIED" json:"priceLevel"`
	Tags       JSONBMapSlice  `gorm:"type:jsonb" json:"tags"`
}

func (AccountLocationReview) TableName() string {
	return "account_location_reviews"
}
