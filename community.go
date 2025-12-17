package models

import (
	"time"

	"gorm.io/gorm"
)

type Community struct {
	ID          uint                   `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID   uint                   `gorm:"index" json:"accountId"`
	Account     *Account               `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	CreatedAt   time.Time              `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time              `gorm:"autoUpdateTime" json:"-"`
	DeletedAt   *gorm.DeletedAt        `gorm:"index" json:"-"`
	Name        string                 `gorm:"type:text;not null" json:"name"`
	Description *string                `gorm:"type:text" json:"description,omitempty"`
	Banner      *string                `gorm:"type:text" json:"banner,omitempty"`
	Newsletters *[]CommunityNewsletter `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"newsletters,omitempty"`
	Members     *[]CommunityMember     `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"members,omitempty"`
	Calendar    *Calendar              `json:"calendar,omitempty"`
}

func (Community) TableName() string {
	return "communities"
}
