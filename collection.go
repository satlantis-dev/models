package models

import (
	"time"

	"gorm.io/gorm"
)

type Collection struct {
	ID          uint                 `gorm:"primaryKey" json:"id"`
	AccountID   uint                 `gorm:"index;uniqueIndex:idx_accountid_name" json:"accountId"`
	Account     Account              `gorm:"foreignKey:AccountID" json:"account"`
	Name        string               `gorm:"type:text;not null;uniqueIndex:idx_accountid_name" json:"name"`
	Description string               `gorm:"type:text" json:"description"`
	Cover       string               `gorm:"type:text" json:"cover"`
	CreatedAt   time.Time            `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time            `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt       `gorm:"index" json:"-"`
	IsPublic    bool                 `gorm:"default:true" json:"isPublic"`
	Locations   []CollectionLocation `json:"locations"`
}

func (Collection) TableName() string {
	return "collections"
}

type CollectionLocation struct {
	CollectionID uint       `gorm:"primaryKey;autoIncrement:false" json:"collectionId"`
	Collection   Collection `gorm:"foreignKey:CollectionID;references:ID" json:"collection"`
	LocationID   uint       `gorm:"primaryKey;autoIncrement:false" json:"locationId"`
	Location     Location   `gorm:"foreignKey:LocationID;references:ID" json:"location"`
	SeqNum       int        `gorm:"default:0" json:"seqNum"`
}
