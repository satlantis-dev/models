package models

import (
	"time"

	"gorm.io/gorm"
)

type CommunityMember struct {
	ID          uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	CommunityID uint            `gorm:"not null;index;uniqueIndex:idx_community_account" json:"communityId"`
	Community   *Community      `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"community,omitempty"`
	AccountID   uint            `gorm:"not null;index;uniqueIndex:idx_community_account" json:"accountId"`
	Account     *AccountDTO     `json:"account,omitempty"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"-"`
}

func (CommunityMember) TableName() string {
	return "community_members"
}
