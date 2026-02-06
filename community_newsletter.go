package models

import (
	"time"

	"gorm.io/gorm"
)

type CommunityNewsletter struct {
	ID          uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	CommunityID uint            `gorm:"index" json:"communityId"`
	Community   *Community      `gorm:"constraint:OnDelete:CASCADE;" json:"community,omitempty"`
	Content     string          `gorm:"type:text" json:"content"`
	Title       string          `gorm:"type:varchar(255)" json:"title"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"-"`
}

func (CommunityNewsletter) TableName() string {
	return "community_newsletters"
}
