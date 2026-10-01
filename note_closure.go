package models

type NoteClosure struct {
	AncestorID   uint `gorm:"primaryKey;autoIncrement:false;uniqueIndex:unique_note_closure;index;index:idx_note_closures_ancestor_depth_descendant,priority:1" json:"ancestorId"`
	DescendantID uint `gorm:"primaryKey;autoIncrement:false;uniqueIndex:unique_note_closure;index;index:idx_note_closures_ancestor_depth_descendant,priority:3" json:"descendantId"`
	Depth        int  `gorm:"index:idx_note_closures_ancestor_depth_descendant,priority:2" json:"depth"`
}
