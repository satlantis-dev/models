package models

import (
	"time"

	"gorm.io/gorm"
)

// WhopCompany stores Whop companies that have installed the Satlantis app.
type WhopCompany struct {
	CompanyID   string         `gorm:"primaryKey;size:64" json:"companyId"`
	CompanyName string         `gorm:"not null" json:"companyName"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
