package models

// LocationLocationTag is the join model for Location.LocationTags / LocationTag.Locations. It exists
// to index location_id: the primary key leads with location_tag_id, so loading a location's tags
// (WHERE location_id IN ...) had no index of its own. Registered in SetupJoinTables.
type LocationLocationTag struct {
	LocationTagID uint         `gorm:"primaryKey" json:"locationTagId"`
	LocationTag   *LocationTag `gorm:"foreignKey:LocationTagID;constraint:OnDelete:CASCADE" json:"locationTag,omitempty"`
	LocationID    uint         `gorm:"primaryKey;index" json:"locationId"`
	Location      *Location    `gorm:"foreignKey:LocationID;constraint:OnDelete:CASCADE" json:"location,omitempty"`
}

func (LocationLocationTag) TableName() string {
	return "location_location_tags"
}
