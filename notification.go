package models

import (
	"time"
)

const (
	NotificationTypeEventAnnouncement               = "calendar_event_announcement"
	NotificationTypeEventCohostInvitation           = "calendar_event_cohost_invitation"
	NotificationTypeEventCohostInvitationAccepted   = "calendar_event_cohost_invitation_accepted"
	NotificationTypeEventCohostInvitationDeclined   = "calendar_event_cohost_invitation_declined"
	NotificationTypeEventCommentMention             = "calendar_event_comment_mention"
	NotificationTypeEventDiscussion                 = "calendar_event_discussion"
	NotificationTypeEventDiscussionMention          = "calendar_event_discussion_mention"
	NotificationTypeEventRSVP                       = "calendar_event_rsvp"
	NotificationTypeEventUpdate                     = "calendar_event_update"
	NotificationTypeEventRsvpGatedRejected          = "calendar_event_rsvp_gated_rejected"
	NotificationTypeCalendarInvitation              = "calendar_invitation"
	NotificationTypeCalendarInvitationAccepted      = "calendar_invitation_accepted"
	NotificationTypeCalendarInvitationDeclined      = "calendar_invitation_declined"
	NotificationTypeCalendarSubscription            = "calendar_subscription"
	NotificationTypeCalendarSubscriptionEventAdded  = "calendar_subscription_event_added"
	NotificationTypeCalendarEventSubmission         = "calendar_event_submission"
	NotificationTypeCalendarEventSubmissionAccepted = "calendar_event_submission_accepted"
	NotificationTypeCalendarEventSubmissionDeclined = "calendar_event_submission_declined"
	NotificationTypeCollectionImportSuccessful      = "collection_import_successful"
	NotificationTypeCollectionInvitation            = "collection_invitation"
	NotificationTypeCollectionInvitationAccepted    = "collection_invitation_accepted"
	NotificationTypeCollectionInvitationDeclined    = "collection_invitation_declined"
	NotificationTypeCollectionShared                = "collection_shared"
	NotificationTypeCollectionSaved                 = "collection_saved"
	NotificationTypeComment                         = "comment"
	NotificationTypeFollow                          = "follow"
	NotificationTypeLikeMediaNote                   = "like-media-note"
	NotificationTypeMention                         = "mention"
	NotificationTypeReport                          = "report"
	NotificationTypeVerifyEmail                     = "verify_email"
)

type Notification struct {
	ID                 uint                    `gorm:"primaryKey" json:"id"`
	CreatorAccountID   uint                    `gorm:"index" json:"creatorAccountId"` // Account that created the event
	CreatorAccount     AccountDTO              `gorm:"foreignKey:CreatorAccountID;references:ID;constraint:OnDelete:CASCADE" json:"creatorAccount"`
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
