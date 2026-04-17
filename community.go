package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Community struct {
	ID          uint                       `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID   uint                       `gorm:"index" json:"accountId"`
	Account     *Account                   `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	CreatedAt   time.Time                  `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time                  `gorm:"autoUpdateTime" json:"-"`
	DeletedAt   *gorm.DeletedAt            `gorm:"index" json:"-"`
	Name        string                     `gorm:"type:text;not null" json:"name"`
	Bio         *string                    `gorm:"type:text" json:"bio,omitempty"`
	Description *string                    `gorm:"type:text" json:"description,omitempty"`
	Banner      *string                    `gorm:"type:text" json:"banner,omitempty"`
	Logo        *string                    `gorm:"type:text" json:"logo,omitempty"`
	Newsletters *[]CommunityNewsletter     `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"newsletters,omitempty"`
	Members     *[]CommunityMember         `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"members,omitempty"`
	MemberCount *int                       `gorm:"-" json:"memberCount,omitempty"`
	Tiers       *[]CommunityMembershipTier `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"tiers,omitempty"`
	Calendars   *[]Calendar                `gorm:"foreignKey:CommunityID;constraint:OnDelete:SET NULL" json:"calendars,omitempty"`
	FAQ         *[]CommunityFAQ            `gorm:"type:jsonb;serializer:json" json:"faq,omitempty"`
	Gallery     *[]CommunityGalleryImage   `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE;" json:"gallery,omitempty"`
	SocialLinks datatypes.JSON             `gorm:"type:jsonb" json:"socialLinks"`
	ThemeID     *uint                      `gorm:"index" json:"themeId,omitempty"`
	Theme       *Theme                     `gorm:"foreignKey:ThemeID;constraint:OnDelete:SET NULL;" json:"theme,omitempty"`
}

type CommunityFAQ struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (Community) TableName() string {
	return "communities"
}
