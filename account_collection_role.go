package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountCollectionRoleType int

const (
	CollectionOwner AccountCollectionRoleType = iota + 1
	CollectionContributor
	CollectionViewer
)

type AccountCollectionRole struct {
	AccountID    uint                      `gorm:"index;primaryKey" json:"accountId"`
	Account      *AccountDTO               `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	CollectionID uint                      `gorm:"index;primaryKey" json:"collectionId"`
	Collection   *Collection               `gorm:"foreignKey:CollectionID;constraint:OnDelete:CASCADE;" json:"collection,omitempty"`
	Type         AccountCollectionRoleType `gorm:"not null" json:"type"`
	CreatedAt    time.Time                 `json:"-"`
	UpdatedAt    time.Time                 `json:"-"`
	DeletedAt    gorm.DeletedAt            `gorm:"index" json:"-"`
}

func (AccountCollectionRole) TableName() string {
	return "account_collection_roles"
}
