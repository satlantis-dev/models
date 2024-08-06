package models

type NoteClosure struct {
	AncestorID   uint `gorm:"primaryKey;autoIncrement:false;uniqueIndex:unique_note_closure" json:"ancestorId"`
	DescendantID uint `gorm:"primaryKey;autoIncrement:false;uniqueIndex:unique_note_closure" json:"descendantId"`
	Depth        int  `json:"depth"`
}
