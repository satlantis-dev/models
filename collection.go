package models

import "time"

type Collection struct {
	ID          uint                 `gorm:"primaryKey" json:"id"`
	AccountID   uint                 `gorm:"index" json:"accountId"`
	Account     AccountDTO           `gorm:"foreignKey:AccountID" json:"account"`
	Name        string               `gorm:"type:text" json:"name"`
	Description string               `gorm:"type:text" json:"description"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
	DeletedAt   *time.Time           `gorm:"index" json:"-,omitempty"`
	IsPublic    bool                 `gorm:"default:true" json:"isPublic"`
	Locations   []CollectionLocation `gorm:"foreignKey:CollectionID" json:"locations,omitempty"`
}

type CollectionLocation struct {
	CollectionID uint `gorm:"primaryKey" json:"collectionId"`
	LocationID   uint `gorm:"primaryKey" json:"locationId"`
}
