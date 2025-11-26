package models

import "github.com/lib/pq"

type Persona struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	Name        string        `gorm:"type:text;uniqueIndex;not null" json:"name"`
	Tribe       string        `gorm:"type:text" json:"tribe"`
	Emoji       string        `gorm:"type:varchar(7)" json:"emoji"`
	InterestIDs pq.Int32Array `gorm:"type:integer[]" json:"interestIds"`
}
