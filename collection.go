package models

import (
	"time"

	"gorm.io/gorm"
)

type Collection struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	AccountID   uint           `gorm:"index;uniqueIndex:idx_accountid_name" json:"accountId"`
	Account     Account        `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	Name        string         `gorm:"type:text;not null;uniqueIndex:idx_accountid_name" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	IsPublic    bool           `gorm:"default:true" json:"isPublic"`
	Locations   []Location     `gorm:"many2many:collection_locations;joinForeignKey:CollectionID;JoinReferences:LocationID" json:"locations,omitempty"`
}

func (Collection) TableName() string {
	return "collections"
}
