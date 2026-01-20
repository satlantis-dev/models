package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type CalendarEventAnnouncement struct {
	ID                uint            `gorm:"primaryKey" json:"id"`
	CalendarEventID   uint            `json:"calendarEventId"`
	NoteID            uint            `json:"noteId"`
	Note              *Note           `gorm:"constraint:OnDelete:CASCADE;" json:"note,omitempty"`
	CreatedAt         time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt         time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt         *gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	ToNostr           bool            `gorm:"default:false" json:"toNostr"`
	ToDiscussion      bool            `gorm:"default:true" json:"toDiscussion"`
	ToEmail           bool            `gorm:"default:true" json:"toEmail"`
	ToNotification    bool            `gorm:"default:true" json:"toNotification"`
	EmailSubject      *string         `gorm:"type:text" json:"emailSubject,omitempty"`
	EmailRecipientIDs pq.Int32Array   `gorm:"type:integer[]" json:"emailRecipientIds"`
}
