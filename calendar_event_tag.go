package models

type CalendarEventTag struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:text;not null;uniqueIndex" json:"name"`
	NumEvents int    `gorm:"" json:"numEvents"`
}

func (CalendarEventTag) TableName() string {
	return "calendar_event_tags"
}
