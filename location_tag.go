package models

type LocationTag struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Category  string     `gorm:"type:text" json:"category"`
	Key       string     `gorm:"type:text" json:"key"`
	Value     string     `gorm:"type:text" json:"value"`
	Eligible  bool       `gorm:"type:boolean" json:"eligible"`
	Locations []Location `gorm:"many2many:location_location_tags" json:"locations"`
}
