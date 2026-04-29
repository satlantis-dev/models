package models

import (
	"time"

	"gorm.io/gorm"
)

type CommunityMembershipRequestType string

const (
	CommunityMembershipRequestTypeNew       CommunityMembershipRequestType = "new"
	CommunityMembershipRequestTypeUpgrade   CommunityMembershipRequestType = "upgrade"
	CommunityMembershipRequestTypeDowngrade CommunityMembershipRequestType = "downgrade"
	CommunityMembershipRequestTypeExtend    CommunityMembershipRequestType = "extend"
)

type CommunityMembershipRequestStatus string

const (
	CommunityMembershipRequestStatusPending   CommunityMembershipRequestStatus = "pending"
	CommunityMembershipRequestStatusAccepted  CommunityMembershipRequestStatus = "accepted"
	CommunityMembershipRequestStatusRejected  CommunityMembershipRequestStatus = "rejected"
	CommunityMembershipRequestStatusCancelled CommunityMembershipRequestStatus = "cancelled"
)

type CommunityMembershipPeriod string

const (
	CommunityMembershipPeriodMonthly CommunityMembershipPeriod = "monthly"
	CommunityMembershipPeriodAnnual  CommunityMembershipPeriod = "annual"
)

type CommunityMembershipRequest struct {
	ID                  uint                             `gorm:"primaryKey;autoIncrement" json:"id"`
	CommunityID         uint                             `gorm:"not null;index:idx_community_membership_application_lookup;uniqueIndex:idx_community_membership_pending_unique,where:status = 'pending' AND deleted_at IS NULL" json:"communityId"`
	Community           *Community                       `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"community,omitempty"`
	AccountID           uint                             `gorm:"not null;index:idx_community_membership_application_lookup;uniqueIndex:idx_community_membership_pending_unique,where:status = 'pending' AND deleted_at IS NULL" json:"accountId"`
	Account             *AccountDTO                      `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	Type                CommunityMembershipRequestType   `gorm:"type:varchar(32);not null;default:'new';index" json:"type"`
	Status              CommunityMembershipRequestStatus `gorm:"type:varchar(32);not null;default:'pending';index" json:"status"`
	CurrentTierID       *uint                            `gorm:"index" json:"currentTierId,omitempty"`
	CurrentTier         *CommunityMembershipTier         `gorm:"foreignKey:CurrentTierID;constraint:OnDelete:SET NULL;" json:"currentTier,omitempty"`
	RequestedTierID     uint                             `gorm:"not null;index" json:"requestedTierId"`
	RequestedTier       *CommunityMembershipTier         `gorm:"foreignKey:RequestedTierID;constraint:OnDelete:RESTRICT;" json:"requestedTier,omitempty"`
	Period              *CommunityMembershipPeriod       `gorm:"type:varchar(16)" json:"period,omitempty"`
	SubscriptionID      *uint                            `gorm:"index" json:"subscriptionId,omitempty"`
	Subscription        *CommunityMembershipSubscription `gorm:"foreignKey:SubscriptionID;constraint:OnDelete:SET NULL;" json:"subscription,omitempty"`
	RegistrationAnswers *map[string]interface{}          `gorm:"type:jsonb;serializer:json" json:"registrationAnswers,omitempty"`
	ReviewedByAccountID *uint                            `gorm:"index" json:"reviewedByAccountId,omitempty"`
	ReviewedByAccount   *AccountDTO                      `gorm:"foreignKey:ReviewedByAccountID;constraint:OnDelete:SET NULL;" json:"reviewedByAccount,omitempty"`
	ReviewNotes         *string                          `gorm:"type:text" json:"reviewNotes,omitempty"`
	ReviewedAt          *time.Time                       `gorm:"type:timestamptz" json:"reviewedAt,omitempty"`
	CreatedAt           time.Time                        `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt           time.Time                        `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt           *gorm.DeletedAt                  `gorm:"index" json:"-"`
}

func (CommunityMembershipRequest) TableName() string {
	return "community_membership_requests"
}
