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

type Verb string

const (
	// account related
	VerbUpdateAccount Verb = "updated_account"
	VerbFollowAccount Verb = "followed_account"
	// note related
	VerbCreateNote Verb = "created_note"
	VerbRepostNote Verb = "reposted_note"
	VerbReplyNote  Verb = "replied_to_note"
	VerbLikeNote   Verb = "liked_note"
	// destination related
	VerbFollowDestination Verb = "followed_destination"
	// collection related
	VerbCreateCollection Verb = "created_collection"
	VerbSaveCollection   Verb = "saved_collection"
	// location related
	VerbSuggestLocation  Verb = "suggested_location"
	VerbClaimLocation    Verb = "claimed_location"
	VerbReviewLocation   Verb = "reviewed_location"
	VerbBookmarkLocation Verb = "bookmarked_location"
	// calendar related
	VerbCreateCalendar Verb = "created_calendar"
	// event related
	VerbCreatedEvent  Verb = "created_event"
	VerbUpdatedEvent  Verb = "updated_event"
	VerbRSVPEvent     Verb = "rsvped_to_event"
	VerbBookmarkEvent Verb = "bookmarked_event"
)

type Activity struct {
	ID         uint                    `gorm:"primaryKey" json:"id"`
	AccountID  uint                    `gorm:"not null;index" json:"accountId"`
	Account    *Account                `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	ObjectType ObjectType              `gorm:"not null;index" json:"objectType"`
	ObjectID   uint                    `gorm:"not null" json:"objectId"`
	Verb       Verb                    `gorm:"not null" json:"verb"`
	Details    *map[string]interface{} `gorm:"type:jsonb;serializer:json" json:"details,omitempty"`
	CreatedAt  time.Time               `gorm:"autoCreateTime" json:"createdAt"`
}

func (Activity) TableName() string {
	return "activities"
}
