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
	VerbCreateAccount Verb = "create_account"
	VerbUpdateAccount Verb = "update_account"
	VerbFollowAccount Verb = "follow_account"
	// note related
	VerbCreateNote Verb = "create_note"
	VerbRepostNote Verb = "repost_note"
	VerbReplyNote  Verb = "reply_note"
	VerbLikeNote   Verb = "like_note"
	// destination related
	VerbAddDestination    Verb = "add_destination"
	VerbFollowDestination Verb = "follow_destination"
	// collection related
	VerbCreateCollection Verb = "create_collection"
	VerbAddToCollection  Verb = "addto_collection"
	VerbSaveCollection   Verb = "save_collection"
	// location related
	VerbAddLocation    Verb = "add_location"
	VerbReviewLocation Verb = "review_location"
	VerbClaimLocation  Verb = "claim_location"
	// calendar related
	VerbCreateCalendar Verb = "create_calendar"
	VerbAddToCalendar  Verb = "addto_calendar"
	// event related
	VerbCreateEvent Verb = "create_event"
	VerbUpdateEvent Verb = "update_event"
	VerbRSVPEvent   Verb = "rsvp_event"
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
