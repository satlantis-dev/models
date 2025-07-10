package models

import (
	"time"

	"gorm.io/gorm"
)

type Collection struct {
	ID          uint                 `gorm:"primaryKey" json:"id"`
	AccountID   uint                 `gorm:"index;uniqueIndex:idx_accountid_name" json:"accountId"`
	Account     *Account             `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	Name        string               `gorm:"type:text;not null;uniqueIndex:idx_accountid_name" json:"name"`
	Description *string              `gorm:"type:text" json:"description,omitempty"`
	Cover       *string              `gorm:"type:text" json:"cover,omitempty"`
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
	CollectionID   uint               `gorm:"primaryKey;autoIncrement:false;not null;uniqueIndex:idx_collectionid_locationid" json:"collectionId"`
	Collection     *Collection        `gorm:"foreignKey:CollectionID;references:ID" json:"collection,omitempty"`
	GoogleID       string             `gorm:"primaryKey;type:text" json:"googleId"`
	SourceLocation *SourceLocationDTO `gorm:"foreignKey:GoogleID;references:GoogleId" json:"sourceLocation,omitempty"`
	Location       *LocationDTO       `gorm:"foreignKey:GoogleID;references:GoogleID" json:"location,omitempty"`
	SeqNum         int                `gorm:"default:0;not null" json:"seqNum"`
	Blurb          *string            `gorm:"type:text" json:"blurb,omitempty"`
}

func (CollectionLocation) TableName() string {
	return "collection_locations"
}
