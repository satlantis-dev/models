package models

import (
	"time"

	"gorm.io/gorm"
)

type Collection struct {
	ID           uint                 `gorm:"primaryKey" json:"id"`
	AccountID    uint                 `gorm:"index;uniqueIndex:idx_accountid_name" json:"accountId"`
	Account      *Account             `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	Name         string               `gorm:"type:text;not null;uniqueIndex:idx_accountid_name" json:"name"`
	Description  *string              `gorm:"type:text" json:"description,omitempty"`
	Cover        *string              `gorm:"type:text" json:"cover,omitempty"`
	CreatedAt    time.Time            `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time            `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt       `gorm:"index" json:"-"`
	IsPublic     bool                 `gorm:"default:true" json:"isPublic"`
	Locations    []CollectionLocation `json:"locations"`
	NumLocations int                  `gorm:"-" json:"numLocations"`
	NumSaves     int                  `gorm:"-" json:"numSaves"`
	Contributors []AccountMiniDTO     `gorm:"-" json:"contributors,omitempty"`
	LocationTags []LocationTag        `gorm:"many2many:collection_location_tags" json:"locationTags"`
	PlaceID      *uint                `json:"placeId"`
}

func (Collection) TableName() string {
	return "collections"
}

type CollectionLocation struct {
	CollectionID uint         `gorm:"primaryKey;autoIncrement:false;not null;uniqueIndex:idx_collectionid_locationid" json:"collectionId"`
	Collection   *Collection  `gorm:"foreignKey:CollectionID;references:ID" json:"collection,omitempty"`
	GoogleID     string       `gorm:"primaryKey;type:text;not null;uniqueIndex:idx_collectionid_locationid" json:"googleId"`
	Location     *LocationDTO `gorm:"-" json:"location,omitempty"`
	SeqNum       int          `gorm:"default:0;not null" json:"seqNum"`
	Blurb        *string      `gorm:"type:text" json:"blurb,omitempty"`
}

func (CollectionLocation) TableName() string {
	return "collection_locations"
}
