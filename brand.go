package models

type Brand struct {
	Name    string `gorm:"primaryKey" json:"name"`
	Type    string `json:"type"`
	Website string `json:"website"`
	Logo    string `json:"logo"`
}
