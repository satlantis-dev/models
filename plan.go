package models

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	ID                    uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                  string          `gorm:"type:text;not null;uniqueIndex" json:"name"`
	Description           *string         `gorm:"type:text" json:"description,omitempty"`
	Currency              OrderCurrency   `gorm:"type:varchar(8);not null" json:"currency"`
	MonthlyAmount         *int64          `gorm:"type:bigint" json:"monthlyAmount,omitempty"`
	YearlyAmount          *int64          `gorm:"type:bigint" json:"yearlyAmount,omitempty"`
	BTCDiscountPercentage *uint           `json:"btcDiscountPercentage,omitempty"`
	Rank                  int             `gorm:"not null;default:0" json:"rank"`
	IsPaid                bool            `gorm:"not null;default:false" json:"isPaid"`
	IsHidden              bool            `gorm:"not null;default:false" json:"isHidden"`
	IsActive              bool            `gorm:"not null;default:true" json:"isActive"`
	TrialDays             int             `gorm:"not null;default:0" json:"trialDays"`
	CreatedAt             time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt             time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt             *gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Plan) TableName() string {
	return "plans"
}
