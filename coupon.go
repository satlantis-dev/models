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
	CouponScopeEvent     CouponScope = "event"
	CouponScopeCalendar  CouponScope = "calendar"
	CouponScopeCommunity CouponScope = "community"
	CouponScopePlan      CouponScope = "plan"
)

type Coupon struct {
	ID                 uint               `gorm:"primaryKey" json:"id"`
	AccountID          *uint              `gorm:"index" json:"-"`
	Account            *Account           `gorm:"foreignKey:AccountID;constraint:OnDelete:SET NULL" json:"-"`
	Scope              CouponScope        `gorm:"type:varchar(16);not null" json:"scope"`
	CalendarEventID    *uint              `gorm:"index;uniqueIndex:idx_generic_coupon_code_event" json:"calendarEventId,omitempty"`
	CalendarEvent      *CalendarEvent     `gorm:"foreignKey:CalendarEventID;constraint:OnDelete:CASCADE" json:"-"`
	CalendarID         *uint              `gorm:"index;uniqueIndex:idx_generic_coupon_code_calendar" json:"calendarId,omitempty"`
	Calendar           *Calendar          `gorm:"foreignKey:CalendarID;constraint:OnDelete:CASCADE" json:"-"`
	CommunityID        *uint              `gorm:"index;uniqueIndex:idx_generic_coupon_code_community" json:"communityId,omitempty"`
	Community          *Community         `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE" json:"-"`
	TicketTypeIDs      pq.Int32Array      `gorm:"type:integer[]" json:"ticketTypeIds,omitempty"` // Optional list of CalendarEventTicketType IDs the coupon applies to; empty/nil applies to all ticket types (event/calendar scope only)
	TierIDs            pq.Int32Array      `gorm:"type:integer[]" json:"tierIds,omitempty"`       // Optional list of CommunityMembershipTier IDs the coupon applies to; empty/nil applies to all tiers (community scope only)
	PlanIDs            pq.Int32Array      `gorm:"type:integer[]" json:"planIds,omitempty"`       // Optional list of Plan IDs the coupon applies to; empty/nil applies to all plans (plan scope only, which is platform-wide by default)
	Code               string             `gorm:"index;uniqueIndex:idx_generic_coupon_code_event;uniqueIndex:idx_generic_coupon_code_calendar;uniqueIndex:idx_generic_coupon_code_community;uniqueIndex:idx_generic_coupon_code_plan,where:scope = 'plan';size:64;not null" json:"code"`
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

func (Coupon) TableName() string {
	return "coupons"
}

// CouponRedemption records a single use of a Coupon. TargetScope/TargetID
// identify what was actually paid for - a CalendarEventTicketOrder ID (event
// or calendar scope), a CommunityMembershipPayment ID (community scope), or
// a PlanPayment ID (plan scope) - as a generic polymorphic pair rather than
// one nullable FK per target type, so a new coupon-consuming domain doesn't
// need its own column added here. There is deliberately no DB-level foreign
// key to the target row (it lives in a different table depending on scope);
// callers are responsible for loading the actual target when needed.
type CouponRedemption struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	CouponID       uint           `gorm:"not null;index:idx_generic_coupon_account_redemption,priority:1" json:"couponId"`
	Coupon         *Coupon        `gorm:"foreignKey:CouponID;constraint:OnDelete:CASCADE" json:"-"`
	AccountID      uint           `gorm:"not null;index:idx_generic_coupon_account_redemption,priority:2" json:"accountId"`
	Account        *Account       `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"-"`
	TargetScope    CouponScope    `gorm:"type:varchar(16);not null;uniqueIndex:idx_coupon_redemption_target,priority:2" json:"targetScope"`
	TargetID       uint           `gorm:"not null;uniqueIndex:idx_coupon_redemption_target,priority:1" json:"targetId"`
	DiscountAmount int64          `gorm:"type:bigint;not null" json:"discountAmount"`
	Currency       OrderCurrency  `gorm:"type:varchar(8);not null" json:"currency"`
	RedeemedAt     time.Time      `json:"redeemedAt"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

func (CouponRedemption) TableName() string {
	return "coupon_redemptions"
}
