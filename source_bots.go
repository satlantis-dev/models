package models

import (
	"time"

	"github.com/lib/pq"
)

type SourceBots struct {
	Npub             string        `gorm:"primaryKey;" json:"npub"`
	Nsec             string        `gorm:"uniqueIndex;" json:"nsec"`
	Name             string        `gorm:"not null;" json:"name"`
	Type             string        `gorm:"not null;" json:"type"`
	ScopeInterests   pq.Int32Array `gorm:"type:integer[]" json:"scopeInterests"`
	ScopePlaces      pq.Int32Array `gorm:"type:integer[]" json:"scopePlaces"`
	AccountID        string        `gorm:"uniqueIndex" json:"accountId"`
	Account          Account       `json:"account" gorm:"foreignKey:AccountID"`
	PotentialFollows pq.Int32Array `gorm:"type:integer[]" json:"potentialFollows"`
	UpdatedOn        time.Time     `json:"updatedOn"`
}

func (SourceBots) TableName() string {
	return "source_bots"
}
