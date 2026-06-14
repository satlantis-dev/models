package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountCommunityRoleType string

const (
	AccountCommunityRoleAdmin AccountCommunityRoleType = "admin"
)

type AccountCommunityRole struct {
	AccountID            uint                     `gorm:"not null;index;uniqueIndex:idx_account_community_role" json:"accountId"`
	Account              *AccountDTO              `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	CommunityID          uint                     `gorm:"not null;index;uniqueIndex:idx_account_community_role" json:"communityId"`
	Community            *Community               `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"community,omitempty"`
	Type                 AccountCommunityRoleType `gorm:"type:varchar(32);not null;uniqueIndex:idx_account_community_role" json:"type"`
	InvitationReceivedAt *time.Time               `gorm:"type:timestamptz" json:"invitationReceivedAt,omitempty"`
	InvitationAcceptedAt *time.Time               `gorm:"type:timestamptz" json:"invitationAcceptedAt,omitempty"`
	InvitationDeclinedAt *time.Time               `gorm:"type:timestamptz" json:"invitationDeclinedAt,omitempty"`
	CreatedAt            time.Time                `json:"createdAt"`
	UpdatedAt            time.Time                `json:"updatedAt"`
	DeletedAt            gorm.DeletedAt           `gorm:"index" json:"-"`
}

func (AccountCommunityRole) TableName() string {
	return "account_community_roles"
}
