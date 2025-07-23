package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountLocationReview struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID     uint           `gorm:"uniqueIndex:idx_account_location" json:"accountId"`
	Account       AccountDTO     `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	LocationID    uint           `gorm:"uniqueIndex:idx_account_location" json:"locationId"`
	Location      *LocationDTO   `gorm:"foreignKey:LocationID" json:"location,omitempty"`
	ReviewText    string         `json:"reviewText"`
	IsPositive    bool           `gorm:"not null;default:true" json:"isPositive"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	PriceLevel    PriceLevel     `json:"priceLevel"`
	OSMStdTags    JSONBMapSlice  `gorm:"type:jsonb" json:"osmStdTags"`
	OSMExtraTags  JSONBMapSlice  `gorm:"type:jsonb" json:"osmExtraTags"`
	SatlantisTags JSONBMapSlice  `gorm:"type:jsonb" json:"satlantisTags"`
}

func (AccountLocationReview) TableName() string {
	return "account_location_reviews"
}
