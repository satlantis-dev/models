package database

type Topic struct {
	ID          uint     `gorm:"primaryKey" json:"id"`
	CategoryID  uint     `gorm:"index" json:"-"`
	Category    Category `json:"-"`
	Description string   `gorm:"type:text" json:"description"`
	InFocus     bool     `json:"inFocus"`
	Name        string   `gorm:"type:text" json:"name"`
	Weight      uint     `json:"weight"`
}
