package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type CouponDiscountType string

const (
	CouponDiscountPercent CouponDiscountType = "percentage"
	CouponDiscountAmount  CouponDiscountType = "fixed_amount"
)

type CouponScope string

const (
	CouponScopeEvent    CouponScope = "event"
	CouponScopeCalendar CouponScope = "calendar"
)

type CalendarEventCoupon struct {
	ID                 uint               `gorm:"primaryKey" json:"id"`
	AccountID          uint               `json:"-"`
	Account            *Account           `gorm:"foreignKey:AccountID;constraint:OnDelete:SET NULL" json:"-"`
	Scope              CouponScope        `gorm:"type:varchar(16);not null" json:"scope"`
	CalendarEventID    *uint              `gorm:"index;uniqueIndex:idx_coupon_code_event" json:"calendarEventId,omitempty"`
	CalendarEvent      *CalendarEvent     `gorm:"foreignKey:CalendarEventID;constraint:OnDelete:CASCADE" json:"-"`
	CalendarID         *uint              `gorm:"index;uniqueIndex:idx_coupon_code_calendar" json:"calendarId,omitempty"`
	Calendar           *Calendar          `gorm:"foreignKey:CalendarID;constraint:OnDelete:CASCADE" json:"-"`
	TicketTypeIDs      pq.Int32Array      `gorm:"type:integer[]" json:"ticketTypeIds,omitempty"`
	Code               string             `gorm:"index;uniqueIndex:idx_coupon_code_event;uniqueIndex:idx_coupon_code_calendar;size:64;not null" json:"code"`
	Description        *string            `json:"description,omitempty"`
	DiscountType       CouponDiscountType `gorm:"type:varchar(16);not null" json:"discountType"`
	DiscountPercentage *uint              `json:"discountPercentage,omitempty"`
	DiscountAmount     *uint              `json:"discountAmount,omitempty"`
	DiscountCurrency   *OrderCurrency     `gorm:"type:varchar(8)" json:"discountCurrency,omitempty"`
	MaxRedemptions     *uint              `json:"maxRedemptions,omitempty"`
	Redemptions        uint               `gorm:"default:0" json:"redemptions"`
	StartsAt           time.Time          `json:"startsAt"`
	EndsAt             time.Time          `json:"endsAt"`
	IsSingleUse        bool               `gorm:"default:false" json:"isSingleUse"`
	IsActive           bool               `gorm:"default:true" json:"isActive"`
	CreatedAt          time.Time          `json:"createdAt"`
	UpdatedAt          time.Time          `json:"updatedAt"`
	DeletedAt          gorm.DeletedAt     `gorm:"index" json:"-"`
}

func (CalendarEventCoupon) TableName() string {
	return "calendar_event_coupons"
}
