package models

import "time"

type ObjectType string

const (
	ObjectTypeAccount     ObjectType = "account"
	ObjectTypeNote        ObjectType = "note"
	ObjectTypeDestination ObjectType = "destination"
	ObjectTypeCollection  ObjectType = "collection"
	ObjectTypeLocation    ObjectType = "location"
	ObjectTypeCalendar    ObjectType = "calendar"
	ObjectTypeEvent       ObjectType = "event"
)

type Action string

const (
	// account related
	ActionUpdateAccount Action = "updated_account"
	ActionFollowAccount Action = "followed_account"
	// note related
	ActionCreateNote Action = "created_note"
	ActionRepostNote Action = "reposted_note"
	ActionReplyNote  Action = "replied_to_note"
	ActionLikeNote   Action = "liked_note"
	// destination related
	ActionFollowDestination Action = "followed_destination"
	// collection related
	ActionCreateCollection Action = "created_collection"
	ActionSaveCollection   Action = "saved_collection"
	// location related
	ActionSuggestLocation  Action = "suggested_location"
	ActionClaimLocation    Action = "claimed_location"
	ActionReviewLocation   Action = "reviewed_location"
	ActionBookmarkLocation Action = "bookmarked_location"
	// calendar related
	ActionCreateCalendar Action = "created_calendar"
	// event related
	ActionCreateEvent   Action = "created_event"
	ActionUpdateEvent   Action = "updated_event"
	ActionRsvpEvent     Action = "rsvped_to_event"
	ActionBookmarkEvent Action = "bookmarked_event"
)

type Activity struct {
	ID         uint                    `gorm:"primaryKey" json:"id"`
	AccountID  uint                    `gorm:"not null;index" json:"accountId"`
	Account    *Account                `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	ObjectType ObjectType              `gorm:"not null;index" json:"objectType"`
	ObjectID   *uint                   `json:"objectId"`
	Action     Action                  `gorm:"not null" json:"action"`
	Details    *map[string]interface{} `gorm:"type:jsonb;serializer:json" json:"details,omitempty"`
	CreatedAt  time.Time               `gorm:"autoCreateTime" json:"createdAt"`
}

func (Activity) TableName() string {
	return "activities"
}
