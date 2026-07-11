package models

import (
	"time"

	"gorm.io/gorm"
)

// WhopCompany stores Whop companies that have installed the Satlantis app.
type WhopCompany struct {
	CompanyID   string         `gorm:"primaryKey;size:64" json:"companyId"`
	CompanyName string         `gorm:"not null" json:"companyName"`
	// LastSeenAt is updated every time someone logs in from this company via
	// the iframe flow (iframe_token + company_id, or iframe_token +
	// experience_id resolved to this company).
	LastSeenAt time.Time      `json:"lastSeenAt"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
