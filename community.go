package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Community struct {
	ID                     uint                          `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID              uint                          `gorm:"index" json:"accountId"`
	Account                *Account                      `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	CreatedAt              time.Time                     `gorm:"autoCreateTime" json:"-"`
	UpdatedAt              time.Time                     `gorm:"autoUpdateTime" json:"-"`
	DeletedAt              *gorm.DeletedAt               `gorm:"index" json:"-"`
	Name                   string                        `gorm:"type:text;not null" json:"name"`
	Bio                    *string                       `gorm:"type:text" json:"bio,omitempty"`
	Blurb                  *string                       `gorm:"type:text" json:"blurb,omitempty"`
	Description            *string                       `gorm:"type:text" json:"description,omitempty"`
	Banner                 *string                       `gorm:"type:text" json:"banner,omitempty"`
	ContactEmail           *string                       `gorm:"type:text" json:"contactEmail"`
	Logo                   *string                       `gorm:"type:text" json:"logo,omitempty"`
	Notice                 *string                       `gorm:"type:text" json:"notice,omitempty"`
	Newsletters            *[]CommunityNewsletter        `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"newsletters,omitempty"`
	Admins                 *[]AccountMiniDTO             `gorm:"-" json:"admins,omitempty"`
	Members                *[]CommunityMember            `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"members,omitempty"`
	Requests               *[]CommunityMembershipRequest `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"requests,omitempty"`
	MemberCount            *int                          `gorm:"-" json:"memberCount,omitempty"`
	Tiers                  *[]CommunityMembershipTier    `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"tiers,omitempty"`
	Calendars              *[]Calendar                   `gorm:"foreignKey:CommunityID;constraint:OnDelete:SET NULL" json:"calendars,omitempty"`
	FAQ                    *[]CommunityFAQ               `gorm:"type:jsonb;serializer:json" json:"faq,omitempty"`
	Gallery                *[]CommunityGalleryImage      `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"gallery,omitempty"`
	SocialLinks            *datatypes.JSON               `gorm:"type:jsonb" json:"socialLinks,omitempty"`
	ChatLinks              *datatypes.JSON               `gorm:"type:jsonb" json:"chatLinks,omitempty"`
	ThemeID                *uint                         `gorm:"index" json:"themeId,omitempty"`
	Theme                  *Theme                        `gorm:"foreignKey:ThemeID;constraint:OnDelete:SET NULL;" json:"theme,omitempty"`
	WhopID                 *string                       `gorm:"type:varchar(20);uniqueIndex" json:"whopId,omitempty"`
	AccountStripeConnectID *uint                         `gorm:"index" json:"accountStripeConnectId,omitempty"`
	AccountStripeConnect   *AccountStripeConnect         `gorm:"foreignKey:AccountStripeConnectID;constraint:OnDelete:SET NULL;" json:"accountStripeConnect,omitempty"`
}

type CommunityFAQ struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (Community) TableName() string {
	return "communities"
}

type CommunityDTO struct {
	ID          uint        `json:"id"`
	AccountID   uint        `json:"accountId"`
	Name        string      `json:"name"`
	Bio         *string     `json:"bio,omitempty"`
	Description *string     `json:"description,omitempty"`
	Banner      *string     `json:"banner,omitempty"`
	Logo        *string     `json:"logo,omitempty"`
	MemberCount *int        `json:"memberCount"`
	Calendars   *[]Calendar `json:"calendars,omitempty"`
	WhopID      *string     `json:"whopId,omitempty"`
}

func (CommunityDTO) TableName() string {
	return "communities"
}

func (c *Community) ToDTO() *CommunityDTO {
	return &CommunityDTO{
		ID:          c.ID,
		AccountID:   c.AccountID,
		Name:        c.Name,
		Bio:         c.Bio,
		Description: c.Description,
		Banner:      c.Banner,
		Logo:        c.Logo,
		MemberCount: c.MemberCount,
		Calendars:   c.Calendars,
		WhopID:      c.WhopID,
	}
}
