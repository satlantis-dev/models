package models

import (
	"time"

	"gorm.io/gorm"
)

type CommunityNewsletter struct {
	ID           uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	CommunityID  uint            `gorm:"index" json:"communityId"`
	Community    *Community      `gorm:"constraint:OnDelete:CASCADE;" json:"community,omitempty"`
	NoteID       uint            `gorm:"index" json:"noteId"`
	Note         *Note           `gorm:"constraint:OnDelete:CASCADE;" json:"note,omitempty"`
	CreatedAt    time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt    *gorm.DeletedAt `gorm:"index" json:"-"`
	EmailSubject *string         `gorm:"type:text" json:"emailSubject,omitempty"`
}

func (CommunityNewsletter) TableName() string {
	return "community_newsletters"
}
