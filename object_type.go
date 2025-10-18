package models

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
