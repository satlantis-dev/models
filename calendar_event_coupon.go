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

type CalendarEventCouponRedemption struct {
	ID             uint                      `gorm:"primaryKey" json:"id"`
	CouponID       uint                      `gorm:"not null;index;uniqueIndex:idx_coupon_account_redemption,priority:1" json:"couponId"`
	Coupon         *CalendarEventCoupon      `gorm:"foreignKey:CouponID;constraint:OnDelete:CASCADE" json:"-"`
	AccountID      uint                      `gorm:"not null;index;uniqueIndex:idx_coupon_account_redemption,priority:2" json:"accountId"`
	Account        *Account                  `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"-"`
	OrderID        uint                      `gorm:"not null;uniqueIndex" json:"orderId"`
	Order          *CalendarEventTicketOrder `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"-"`
	DiscountAmount int64                     `gorm:"type:bigint;not null" json:"discountAmount"`
	Currency       OrderCurrency             `gorm:"type:varchar(8);not null" json:"currency"`
	RedeemedAt     time.Time                 `json:"redeemedAt"`
	CreatedAt      time.Time                 `json:"-"`
	UpdatedAt      time.Time                 `json:"-"`
	DeletedAt      gorm.DeletedAt            `gorm:"index" json:"-"`
}

func (CalendarEventCouponRedemption) TableName() string {
	return "calendar_event_coupon_redemptions"
}
