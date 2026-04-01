package models

import (
	"time"

	"gorm.io/gorm"
)

type CommunityMembershipTier struct {
	ID          uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	CommunityID uint            `gorm:"not null;uniqueIndex:idx_community_tier_name" json:"communityId"`
	Community   *Community      `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"community,omitempty"`
	Name        string          `gorm:"type:text;not null;uniqueIndex:idx_community_tier_name" json:"name"`
	Description *string         `gorm:"type:text" json:"description,omitempty"`
	Rank        int             `gorm:"not null;default:0" json:"rank"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"-"`
}

func (CommunityMembershipTier) TableName() string {
	return "community_membership_tiers"
}
