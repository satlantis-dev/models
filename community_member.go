package models

import (
	"time"

	"gorm.io/gorm"
)

type CommunityMemberEngagementStage string

const (
	CommunityMemberEngagementStageUnknown             CommunityMemberEngagementStage = "unknown"
	CommunityMemberEngagementStageImportedContact     CommunityMemberEngagementStage = "imported_contact"
	CommunityMemberEngagementStageEventAttendee       CommunityMemberEngagementStage = "event_attendee"
	CommunityMemberEngagementStageCalendarSubscriber  CommunityMemberEngagementStage = "calendar_subscriber"
	CommunityMemberEngagementStageMembershipRequested CommunityMemberEngagementStage = "membership_requested"
	CommunityMemberEngagementStageMember              CommunityMemberEngagementStage = "member"
)

type CommunityMember struct {
	ID                  uint                               `gorm:"primaryKey;autoIncrement" json:"id"`
	CommunityID         uint                               `gorm:"not null;index;uniqueIndex:idx_community_account" json:"communityId"`
	Community           *Community                         `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"community,omitempty"`
	AccountID           uint                               `gorm:"not null;index;uniqueIndex:idx_community_account" json:"accountId"`
	Account             *AccountDTO                        `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	TierID              *uint                              `gorm:"index" json:"tierId,omitempty"`
	Tier                *CommunityMembershipTier           `gorm:"foreignKey:TierID;constraint:OnDelete:SET NULL;" json:"tier,omitempty"`
	OpenSubscriptions   *[]CommunityMembershipSubscription `gorm:"-" json:"openSubscriptions,omitempty"`
	OpenRequests        *[]CommunityMembershipRequest      `gorm:"-" json:"openRequests,omitempty"`
	RegistrationAnswers *RegistrationAnswersPayload        `gorm:"type:jsonb;serializer:json" json:"registrationAnswers,omitempty"`
	StartDate           *time.Time                         `gorm:"type:timestamptz" json:"startDate,omitempty"`
	ExpiryDate          *time.Time                         `gorm:"type:timestamptz;index" json:"expiryDate,omitempty"`
	IsExpired           bool                               `gorm:"-" json:"isExpired"`
	IsInvited           bool                               `gorm:"not null;default:false" json:"isInvited"`
	IsBanned            bool                               `gorm:"not null;default:false" json:"isBanned"`
	EngagementStage     CommunityMemberEngagementStage     `gorm:"type:varchar(32);not null;default:'unknown';index" json:"engagementStage"`
	CreatedAt           time.Time                          `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt           time.Time                          `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt           *gorm.DeletedAt                    `gorm:"index" json:"-"`
}

func (CommunityMember) TableName() string {
	return "community_members"
}

type CommunityMemberDTO struct {
	ID                  uint                               `json:"id"`
	CommunityID         uint                               `json:"communityId"`
	AccountID           uint                               `json:"accountId"`
	Account             *AccountDTO                        `json:"account,omitempty"`
	TierID              *uint                              `json:"tierId,omitempty"`
	Tier                *CommunityMembershipTier           `json:"tier,omitempty"`
	OpenSubscriptions   *[]CommunityMembershipSubscription `json:"openSubscriptions,omitempty"`
	OpenRequests        *[]CommunityMembershipRequest      `json:"openRequests,omitempty"`
	RegistrationAnswers *RegistrationAnswersPayload        `json:"registrationAnswers,omitempty"`
	StartDate           *time.Time                         `json:"startDate,omitempty"`
	ExpiryDate          *time.Time                         `json:"expiryDate,omitempty"`
	IsExpired           bool                               `json:"isExpired"`
	IsInvited           bool                               `json:"isInvited"`
	EngagementStage     CommunityMemberEngagementStage     `json:"engagementStage"`
}

func (m CommunityMember) ToDTO() CommunityMemberDTO {
	return CommunityMemberDTO{
		ID:                  m.ID,
		CommunityID:         m.CommunityID,
		AccountID:           m.AccountID,
		Account:             m.Account,
		TierID:              m.TierID,
		Tier:                m.Tier,
		OpenSubscriptions:   m.OpenSubscriptions,
		OpenRequests:        m.OpenRequests,
		RegistrationAnswers: m.RegistrationAnswers,
		StartDate:           m.StartDate,
		ExpiryDate:          m.ExpiryDate,
		IsExpired:           m.IsExpired,
		IsInvited:           m.IsInvited,
		EngagementStage:     m.EngagementStage,
	}
}

func (CommunityMemberDTO) TableName() string {
	return "community_members"
}

type CommunityMemberMiniDTO struct {
	ID          uint                     `json:"id"`
	CommunityID uint                     `json:"communityId"`
	AccountID   uint                     `json:"accountId"`
	Account     *SearchAccountDTO        `json:"account,omitempty"`
	TierID      *uint                    `json:"tierId,omitempty"`
	Tier        *CommunityMembershipTier `json:"tier,omitempty"`
}

func (m CommunityMember) ToMiniDTO() CommunityMemberMiniDTO {
	var searchAccountDTO *SearchAccountDTO
	if m.Account != nil {
		dto := m.Account.ToSearchAccountDTO()
		searchAccountDTO = &dto
	}
	return CommunityMemberMiniDTO{
		ID:          m.ID,
		CommunityID: m.CommunityID,
		AccountID:   m.AccountID,
		Account:     searchAccountDTO,
		TierID:      m.TierID,
		Tier:        m.Tier,
	}
}

func (CommunityMemberMiniDTO) TableName() string {
	return "community_members"
}
