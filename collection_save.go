package models

import "time"

type CollectionSave struct {
	AccountID    uint        `gorm:"index;uniqueIndex:idx_collection_save_unique" json:"accountId"`
	Account      *Account    `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account,omitempty"`
	CollectionID uint        `gorm:"index;uniqueIndex:idx_collection_save_unique" json:"collectionId"`
	Collection   *Collection `gorm:"foreignKey:CollectionID;constraint:OnDelete:CASCADE;" json:"collection,omitempty"`
	CreatedAt    time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}

func (CollectionSave) TableName() string {
	return "collection_saves"
}
