package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CalendarEventTicketType struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	CalendarEventID   uint           `gorm:"not null;index" json:"calendarEventId"`
	CalendarEvent     *CalendarEvent `gorm:"foreignKey:CalendarEventID;constraint:OnDelete:CASCADE" json:"-"`
	Name              string         `json:"name"`
	Description       *string        `json:"description"`
	PriceSats         *int64         `gorm:"type:bigint" json:"priceSats"`
	PriceFiat         *int64         `gorm:"type:bigint" json:"priceFiat"`
	FiatCurrency      *OrderCurrency `gorm:"type:varchar(8)" json:"fiatCurrency"`
	PriceCurrency     *OrderCurrency `gorm:"type:varchar(10)" json:"priceCurrency"`
	SellCurrencies    datatypes.JSON `gorm:"type:jsonb" json:"sellCurrencies"`
	PriceAmount       *int64         `json:"priceAmount"`
	PriceAmountForBTC *int64         `json:"priceAmountForBTC"`
	MaxCapacity       *uint          `json:"maxCapacity"`
	SellStartDate     *time.Time     `json:"sellStartDate"`
	SellEndDate       *time.Time     `json:"sellEndDate"`
	CreatedByID       *uint          `json:"-"`
	CreatedBy         *Account       `gorm:"foreignKey:CreatedByID;constraint:OnDelete:SET NULL" json:"-"`
	CreatedAt         time.Time      `json:"createdAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	IsHidden          bool           `gorm:"default:false" json:"isHidden"`
}

type CalendarEventTicketTypePromotion struct {
	ID                 uint                     `gorm:"primaryKey" json:"id"`
	TicketTypeID       uint                     `gorm:"not null;index" json:"ticketTypeId"`
	TicketType         *CalendarEventTicketType `gorm:"foreignKey:TicketTypeID;constraint:OnDelete:CASCADE" json:"-"`
	Name               string                   `gorm:"size:30;not null" json:"name"`
	DiscountPercentage uint                     `gorm:"not null" json:"discountPercentage"`
	EndDate            *time.Time               `json:"endDate,omitempty"`
	MaxQuantity        *uint                    `json:"maxQuantity,omitempty"`
	RedeemedQuantity   uint                     `gorm:"default:0" json:"redeemedQuantity"`
	IsActive           bool                     `gorm:"default:true" json:"isActive"`
	CreatedAt          time.Time                `json:"createdAt"`
	UpdatedAt          time.Time                `json:"updatedAt"`
	DeletedAt          gorm.DeletedAt           `gorm:"index" json:"-"`
}
