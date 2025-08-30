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
	CollectionInvited
)

type AccountCollectionRole struct {
	AccountID    uint                      `gorm:"not null;index;uniqueIndex:idx_account_collection_role" json:"accountId"`
	Account      *Account                  `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	CollectionID uint                      `gorm:"not null;index;uniqueIndex:idx_account_collection_role" json:"collectionId"`
	Collection   *Collection               `gorm:"foreignKey:CollectionID;constraint:OnDelete:CASCADE;" json:"collection,omitempty"`
	Type         AccountCollectionRoleType `gorm:"not null" json:"type"`
	CreatedAt    time.Time                 `json:"-"`
	UpdatedAt    time.Time                 `json:"-"`
	DeletedAt    gorm.DeletedAt            `gorm:"index" json:"-"`
}

func (AccountCollectionRole) TableName() string {
	return "account_collection_roles"
}
