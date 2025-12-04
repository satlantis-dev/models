package models

type CalendarEventTag struct {
	ID        uint   `gorm:"primaryKey" json:"id,omitempty"`
	Name      string `gorm:"type:text;not null;uniqueIndex" json:"name,omitempty"`
	NumEvents *int   `gorm:"default:0" json:"numEvents,omitempty"`
}

func (CalendarEventTag) TableName() string {
	return "calendar_event_tags"
}
