package models

import "github.com/lib/pq"

type LocationCategory struct {
	Name       string         `gorm:"primaryKey" json:"name"`
	Eligible   bool           `gorm:"type:boolean" json:"eligible"`
	Primary    bool           `gorm:"type:boolean" json:"primary"`
	Highlights pq.StringArray `gorm:"type:text[]" json:"highlights"`
}

type LocationTag struct {
	ID               uint             `gorm:"primaryKey" json:"id"`
	Category         string           `json:"category"`
	LocationCategory LocationCategory `gorm:"foreignKey:Category;references:Name" json:"-"`
	Key              string           `gorm:"type:text" json:"key"`
	Value            string           `gorm:"type:text" json:"value"`
	OsmPull          bool             `gorm:"type:boolean" json:"osmPull"`
	Eligible         bool             `gorm:"type:boolean" json:"eligible"`
	Locations        []Location       `gorm:"many2many:location_location_tags" json:"locations"`
	Section          *string          `gorm:"type:text" json:"section"`
}
