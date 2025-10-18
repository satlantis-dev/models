package models

type VanityPath struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Path       string     `gorm:"type:text;uniqueIndex" json:"path"`
	ObjectType ObjectType `gorm:"type:text" json:"objectType"`
	ObjectRef  string     `json:"objectRef"`
}

func (VanityPath) TableName() string {
	return "vanity_paths"
}
