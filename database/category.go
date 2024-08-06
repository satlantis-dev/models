package database

type Category struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Description string `gorm:"type:text" json:"description"`
	Name        string `gorm:"type:text" json:"name"`
}
