package models

type PlaceClosure struct {
	AncestorID   uint `gorm:"primaryKey;autoIncrement:false" json:"ancestorId"`
	DescendantID uint `gorm:"primaryKey;autoIncrement:false" json:"descendantId"`
	Depth        int  `json:"depth"`
}
