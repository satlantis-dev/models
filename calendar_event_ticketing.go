package models

import (
	"gorm.io/gorm"
	"time"
)

type CalendarEventTicketType struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	CalendarEventID uint           `gorm:"not null;index" json:"calendarEventId"`
	Name            string         `json:"name"`
	Description     *string        `json:"description"`
	PriceSats       *int64         `gorm:"type:bigint" json:"priceSats"`
	PriceFiat       *int64         `gorm:"type:bigint" json:"priceFiat"`
	MaxCapacity     *uint          `json:"maxCapacity"`
	SellStartDate   *time.Time     `json:"sellStartDate"`
	SellEndDate     *time.Time     `json:"sellEndDate"`
	CreatedByID     *uint          `json:"-"`
	CreatedBy       *Account       `json:"-"`
	CreatedAt       time.Time      `json:"createdAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
