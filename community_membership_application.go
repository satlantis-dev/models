package models

import (
	"time"

	"gorm.io/gorm"
)

type CommunityMembershipApplicationType string

const (
	CommunityMembershipApplicationTypeNew     CommunityMembershipApplicationType = "new"
	CommunityMembershipApplicationTypeUpgrade CommunityMembershipApplicationType = "upgrade"
)

type CommunityMembershipApplicationStatus string

const (
	CommunityMembershipApplicationStatusPending  CommunityMembershipApplicationStatus = "pending"
	CommunityMembershipApplicationStatusApproved CommunityMembershipApplicationStatus = "approved"
	CommunityMembershipApplicationStatusRejected CommunityMembershipApplicationStatus = "rejected"
	CommunityMembershipApplicationStatusCanceled CommunityMembershipApplicationStatus = "canceled"
)

type CommunityMembershipApplication struct {
	ID                  uint                                 `gorm:"primaryKey;autoIncrement" json:"id"`
	CommunityID         uint                                 `gorm:"not null;index:idx_community_membership_application_lookup" json:"communityId"`
	Community           *Community                           `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"community,omitempty"`
	AccountID           uint                                 `gorm:"not null;index:idx_community_membership_application_lookup" json:"accountId"`
	Account             *AccountDTO                          `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	Type                CommunityMembershipApplicationType   `gorm:"type:varchar(32);not null;default:'new';index" json:"type"`
	Status              CommunityMembershipApplicationStatus `gorm:"type:varchar(32);not null;default:'pending';index" json:"status"`
	CurrentTierID       *uint                                `gorm:"index" json:"currentTierId,omitempty"`
	CurrentTier         *CommunityMembershipTier             `gorm:"foreignKey:CurrentTierID;constraint:OnDelete:SET NULL;" json:"currentTier,omitempty"`
	RequestedTierID     uint                                 `gorm:"not null;index" json:"requestedTierId"`
	RequestedTier       *CommunityMembershipTier             `gorm:"foreignKey:RequestedTierID;constraint:OnDelete:RESTRICT;" json:"requestedTier,omitempty"`
	RegistrationAnswers *map[string]interface{}              `gorm:"type:jsonb;serializer:json" json:"registrationAnswers,omitempty"`
	ReviewedByAccountID *uint                                `gorm:"index" json:"reviewedByAccountId,omitempty"`
	ReviewedByAccount   *AccountDTO                          `gorm:"foreignKey:ReviewedByAccountID;constraint:OnDelete:SET NULL;" json:"reviewedByAccount,omitempty"`
	ReviewNotes         *string                              `gorm:"type:text" json:"reviewNotes,omitempty"`
	ReviewedAt          *time.Time                           `gorm:"type:timestamptz" json:"reviewedAt,omitempty"`
	CreatedAt           time.Time                            `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt           time.Time                            `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt           *gorm.DeletedAt                      `gorm:"index" json:"-"`
}

func (CommunityMembershipApplication) TableName() string {
	return "community_membership_applications"
}
