package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type LightningContact struct {
	ID               uint            `gorm:"primaryKey" json:"id"`
	AccountID        uint            `gorm:"not null;uniqueIndex:idx_lightning_contact_account_address,priority:1" json:"accountId"`
	Account          *Account        `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"-"`
	Name             string          `gorm:"type:text;not null" json:"name"`
	LightningAddress string          `gorm:"type:text;not null;uniqueIndex:idx_lightning_contact_account_address,priority:2" json:"lightningAddress"`
	ContactAccountID *uint           `gorm:"index" json:"contactAccountId,omitempty"`
	ContactAccount   *Account        `gorm:"foreignKey:ContactAccountID;constraint:OnDelete:SET NULL" json:"-"`
	Metadata         *datatypes.JSON `gorm:"type:jsonb" json:"metadata,omitempty"`
	CreatedAt        time.Time       `json:"createdAt"`
	UpdatedAt        time.Time       `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt  `gorm:"index" json:"-"`
}

type LightningContactDTO struct {
	ID               uint            `json:"id"`
	Name             string          `json:"name"`
	LightningAddress string          `json:"lightningAddress"`
	ContactAccountID *uint           `json:"contactAccountId,omitempty"`
	Metadata         *datatypes.JSON `json:"metadata,omitempty"`
	CreatedAt        time.Time       `json:"createdAt"`
	UpdatedAt        time.Time       `json:"updatedAt"`
}

func (c *LightningContact) ToDTO() LightningContactDTO {
	return LightningContactDTO{
		ID:               c.ID,
		Name:             c.Name,
		LightningAddress: c.LightningAddress,
		ContactAccountID: c.ContactAccountID,
		Metadata:         c.Metadata,
		CreatedAt:        c.CreatedAt,
		UpdatedAt:        c.UpdatedAt,
	}
}
