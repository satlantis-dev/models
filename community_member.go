package models

import (
	"time"

	"gorm.io/gorm"
)

type CommunityMember struct {
	ID                        uint                               `gorm:"primaryKey;autoIncrement" json:"id"`
	CommunityID               uint                               `gorm:"not null;index;uniqueIndex:idx_community_account" json:"communityId"`
	Community                 *Community                         `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"community,omitempty"`
	AccountID                 uint                               `gorm:"not null;index;uniqueIndex:idx_community_account" json:"accountId"`
	Account                   *AccountDTO                        `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	TierID                    *uint                              `gorm:"index" json:"tierId,omitempty"`
	Tier                      *CommunityMembershipTier           `gorm:"foreignKey:TierID;constraint:OnDelete:SET NULL;" json:"tier,omitempty"`
	OpenSubscriptions         *[]CommunityMembershipSubscription `gorm:"-" json:"openSubscriptions,omitempty"`
	OpenRequests              *[]CommunityMembershipRequest      `gorm:"-" json:"openRequests,omitempty"`
	RegistrationAnswers       *RegistrationAnswersPayload        `gorm:"type:jsonb;serializer:json" json:"registrationAnswers,omitempty"`
	StartDate                 *time.Time                         `gorm:"type:timestamptz" json:"startDate,omitempty"`
	ExpiryDate                *time.Time                         `gorm:"type:timestamptz;index" json:"expiryDate,omitempty"`
	IsExpired                 bool                               `gorm:"-" json:"isExpired"`
	IsInvited                 bool                               `gorm:"not null;default:false" json:"isInvited"`
	IsBanned                  bool                               `gorm:"not null;default:false" json:"isBanned"`
	IsCommunityAdmin          bool                               `gorm:"not null;default:false" json:"isCommunityAdmin"`
	AdminInvitationReceivedAt *time.Time                         `json:"adminInvitationReceivedAt"`
	AdminInvitationAcceptedAt *time.Time                         `json:"adminInvitationAcceptedAt"`
	AdminInvitationDeclinedAt *time.Time                         `json:"adminInvitationDeclinedAt"`
	CreatedAt                 time.Time                          `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt                 time.Time                          `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt                 *gorm.DeletedAt                    `gorm:"index" json:"-"`
}

func (CommunityMember) TableName() string {
	return "community_members"
}
