package models

import (
	"time"

	"gorm.io/gorm"
)

type Community struct {
	ID          uint                   `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt   time.Time              `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time              `gorm:"autoUpdateTime" json:"-"`
	DeletedAt   *gorm.DeletedAt        `gorm:"index" json:"-"`
	Newsletters *[]CommunityNewsletter `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"newsletters,omitempty"`
	Members     *[]CommunityMember     `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"members,omitempty"`
	Calendar    *Calendar              `json:"calendar,omitempty"`
}

func (Community) TableName() string {
	return "communities"
}
