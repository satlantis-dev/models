package models

import (
	"time"
)

const (
	NotificationTypeCalendarEventAnnouncement             = "calendar_event_announcement"
	NotificationTypeCalendarEventCohostInvitation         = "calendar_event_cohost_invitation"
	NotificationTypeCalendarEventCohostInvitationAccepted = "calendar_event_cohost_invitation_accepted"
	NotificationTypeCalendarEventCohostInvitationDeclined = "calendar_event_cohost_invitation_declined"
	NotificationTypeCalendarEventCommentMention           = "calendar_event_comment_mention"
	NotificationTypeCalendarEventDiscussion               = "calendar_event_discussion"
	NotificationTypeCalendarEventDiscussionMention        = "calendar_event_discussion_mention"
	NotificationTypeCalendarEventRSVP                     = "calendar_event_rsvp"
	NotificationTypeCalendarEventUpdate                   = "calendar_event_update"
	NotificationTypeCollectionImportSuccessful            = "collection_import_successful"
	NotificationTypeCollectionInvitation                  = "collection_invitation"
	NotificationTypeCollectionSaved                       = "collection_saved"
	NotificationTypeComment                               = "comment"
	NotificationTypeFollow                                = "follow"
	NotificationTypeLikeMediaNote                         = "like-media-note"
	NotificationTypeMention                               = "mention"
	NotificationTypeReport                                = "report"
	NotificationTypeVerifyEmail                           = "verify_email"
)

type Notification struct {
	ID                 uint                    `gorm:"primaryKey" json:"id"`
	CreatorAccountID   uint                    `gorm:"index" json:"creatorAccountId"` // Account that created the event
	CreatorAccount     AccountDTO              `gorm:"foreignKey:CreatorAccountID;references:ID" json:"creatorAccount"`
	RecipientAccountID uint                    `gorm:"index" json:"-"` // Account that will be notified
	Type               string                  `gorm:"type:varchar(255)" json:"type"`
	Action             string                  `gorm:"type:text" json:"action"`
	ImageURL           string                  `gorm:"type:text" json:"imageUrl"`
	Link               string                  `gorm:"type:text" json:"link"`
	Message            string                  `gorm:"type:text" json:"message"`
	IsRead             bool                    `gorm:"default:false" json:"isRead"`
	Payload            *map[string]interface{} `gorm:"type:jsonb;serializer:json" json:"payload,omitempty"`
	CreatedAt          time.Time               `json:"createdAt"`
	UpdatedAt          time.Time               `json:"-"`
}
