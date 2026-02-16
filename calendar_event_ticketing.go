package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type OrderCurrency string

var (
	OrderCurrencyBTC OrderCurrency = "BTC"
	OrderCurrencyUSD OrderCurrency = "USD"
)

type OrderStatus string

var (
	OrderPending  OrderStatus = "pending"
	OrderPaid     OrderStatus = "paid"
	OrderCanceled OrderStatus = "cancelled"
	OrderRefunded OrderStatus = "refunded"
)

type TicketStatus string

var (
	TicketActive   TicketStatus = "active"
	TicketUsed     TicketStatus = "used"
	TicketRefunded TicketStatus = "refunded"
	TicketCanceled TicketStatus = "cancelled"
	TicketReissued TicketStatus = "reissued"
)

type PaymentMethod string

const (
	PaymentMethodLightning PaymentMethod = "lightning"
	PaymentMethodStripe    PaymentMethod = "stripe"
)

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "pending"
	PaymentPaid      PaymentStatus = "paid"
	PaymentExpired   PaymentStatus = "expired"
	PaymentFailed    PaymentStatus = "failed"
	PaymentRefunded  PaymentStatus = "refunded"
	PaymentCancelled PaymentStatus = "cancelled"
)

// RefundStatus represents the status of a refund
type RefundStatus string

const (
	RefundPending    RefundStatus = "pending"
	RefundProcessing RefundStatus = "processing"
	RefundCompleted  RefundStatus = "completed"
	RefundFailed     RefundStatus = "failed"
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
}

type CalendarEventTicketOrder struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	CalendarEventID uint           `gorm:"not null;index" json:"calendarEventId"`
	CalendarEvent   *CalendarEvent `gorm:"foreignKey:CalendarEventID;constraint:OnDelete:CASCADE" json:"-"`
	AccountID       uint           `gorm:"not null;index" json:"accountId"`
	Account         *Account       `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"-"`
	TotalPrice      int64          `gorm:"type:bigint" json:"totalPrice"`
	Currency        OrderCurrency  `gorm:"type:varchar(8)" json:"currency"`
	RefundedAmount  int64          `gorm:"type:bigint;default:0" json:"refundedAmount"`
	PriceCurrency   *OrderCurrency `gorm:"type:varchar(10)" json:"priceCurrency"`
	PriceAmount     *int64         `json:"priceAmount"`
	Status          OrderStatus    `gorm:"type:varchar(32);default:'pending'" json:"status"`
	RsvpData        datatypes.JSON `gorm:"type:jsonb" json:"rsvpData,omitempty"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

type CalendarEventTicketOrderItem struct {
	ID             uint                     `gorm:"primaryKey" json:"id"`
	OrderID        uint                     `gorm:"not null;index" json:"orderId"`
	Order          CalendarEventTicketOrder `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"order"`
	TicketTypeID   uint                     `gorm:"not null;index" json:"ticketTypeId"`
	TicketType     CalendarEventTicketType  `gorm:"foreignKey:TicketTypeID;constraint:OnDelete:CASCADE" json:"ticketType"`
	Quantity       uint                     `json:"quantity"`
	PriceEach      int64                    `gorm:"type:bigint" json:"priceEach"`
	Currency       OrderCurrency            `gorm:"type:varchar(8)" json:"currency"`
	RefundedAmount int64                    `gorm:"type:bigint;default:0" json:"refundedAmount"`
	PriceCurrency  *OrderCurrency           `gorm:"type:varchar(10)" json:"priceCurrency"`
	PriceAmount    *int64                   `json:"priceAmount"`
	Status         OrderStatus              `gorm:"type:varchar(32);default:'pending'" json:"status"`
	CreatedAt      time.Time                `json:"-"`
	UpdatedAt      time.Time                `json:"-"`
	DeletedAt      gorm.DeletedAt           `gorm:"index" json:"-"`
}

type CalendarEventTicket struct {
	ID          uint                         `gorm:"primaryKey" json:"id"`
	OrderItemID uint                         `gorm:"not null;index" json:"orderItemId"`
	OrderItem   CalendarEventTicketOrderItem `gorm:"foreignKey:OrderItemID;constraint:OnDelete:CASCADE" json:"orderItem"`
	AccountID   uint                         `gorm:"not null;index" json:"accountId"`
	Account     *Account                     `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"-"`
	RsvpID      *uint                        `gorm:"index" json:"rsvpId,omitempty"`
	RSVP        *CalendarEventRSVP           `gorm:"foreignKey:RsvpID;constraint:OnDelete:CASCADE" json:"rsvp,omitempty"`
	Status      TicketStatus                 `gorm:"type:varchar(32);default:'active'" json:"status"`
	Code        string                       `gorm:"uniqueIndex;size:64" json:"code"`
	CheckedInAt *time.Time                   `json:"checkedInAt,omitempty"`
	CreatedAt   time.Time                    `json:"-"`
	UpdatedAt   time.Time                    `json:"-"`
	DeletedAt   gorm.DeletedAt               `gorm:"index" json:"-"`
}

type CalendarEventTicketOrderPayment struct {
	ID                       uint                     `gorm:"primaryKey" json:"id"`
	OrderID                  uint                     `gorm:"uniqueIndex;not null" json:"orderId"`
	Order                    CalendarEventTicketOrder `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"order"`
	PaymentMethod            PaymentMethod            `gorm:"type:varchar(32);not null" json:"paymentMethod"`
	Status                   PaymentStatus            `gorm:"type:varchar(32);default:'pending'" json:"status"`
	Amount                   int64                    `gorm:"not null" json:"amount"` // Cents for fiat, sats for BTC/Lightning
	Currency                 OrderCurrency            `gorm:"type:varchar(8);not null" json:"currency"`
	ExchangeRate             *float64                 `json:"exchangeRate"`
	ExchangeRateSource       *string                  `gorm:"type:varchar(50)" json:"exchangeRateSource"`
	LightningPaymentHash     *string                  `gorm:"uniqueIndex;size:64" json:"lightningPaymentHash,omitempty"`
	LightningPaymentRequest  *string                  `gorm:"type:text" json:"lightningPaymentRequest,omitempty"`
	LightningPreimage        *string                  `gorm:"size:64" json:"lightningPreimage,omitempty"`
	LightningProvider        *string                  `gorm:"size:32" json:"lightningProvider,omitempty"`
	LightningProviderTxID    *string                  `gorm:"index" json:"lightningProviderTxId,omitempty"`
	PaymentProviderReference *string                  `gorm:"index" json:"paymentProviderReference"`
	Metadata                 *datatypes.JSON          `gorm:"type:jsonb" json:"metadata,omitempty"`
	PaidAt                   *time.Time               `json:"paidAt,omitempty"`
	ExpiredAt                *time.Time               `json:"expiredAt,omitempty"`
	ExpiresAt                *time.Time               `json:"expiresAt,omitempty"`
	FailedAt                 *time.Time               `json:"failedAt,omitempty"`
	RefundedAt               *time.Time               `json:"refundedAt,omitempty"`
	CancelledAt              *time.Time               `json:"cancelledAt,omitempty"`
	CreatedAt                time.Time                `json:"createdAt"`
	UpdatedAt                time.Time                `json:"updatedAt"`
	DeletedAt                gorm.DeletedAt           `gorm:"index" json:"-"`
}

// CalendarEventTicketOrderRefund represents a refund for a ticket order
type CalendarEventTicketOrderRefund struct {
	ID           uint                      `gorm:"primarykey" json:"id"`
	OrderID      uint                      `gorm:"not null;index" json:"orderId"`
	Order        *CalendarEventTicketOrder `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"-"`
	Amount       int64                     `gorm:"not null" json:"amount"`
	Fee          int64                     `gorm:"type:bigint;default:0" json:"fee"`
	Currency     string                    `gorm:"not null" json:"currency"`
	Status       RefundStatus              `gorm:"not null;default:'pending'" json:"status"`
	RefundMethod string                    `gorm:"not null" json:"refundMethod"` // "lightning", "stripe", etc.

	LightningAddress     *string `json:"lightningAddress,omitempty"`
	LightningPaymentHash *string `json:"lightningPaymentHash,omitempty"`
	LightningPreimage    *string `gorm:"size:64" json:"lightningPreimage,omitempty"`

	StripeRefundID        *string `json:"stripeRefundId,omitempty"`
	StripePaymentIntentID *string `json:"stripePaymentIntentId,omitempty"`

	Reason        *string         `gorm:"type:text" json:"reason,omitempty"`
	CreatedAt     time.Time       `json:"createdAt"`
	ProcessedAt   *time.Time      `json:"processedAt,omitempty"`
	FailedAt      *time.Time      `json:"failedAt,omitempty"`
	FailureReason *string         `gorm:"type:text" json:"failureReason,omitempty"`
	Metadata      *datatypes.JSON `gorm:"type:jsonb" json:"metadata,omitempty"`
	DeletedAt     gorm.DeletedAt  `gorm:"index" json:"-"`
}

type CouponDiscountType string

const (
	CouponDiscountPercent CouponDiscountType = "percent"
	CouponDiscountAmount  CouponDiscountType = "amount"
)

type CouponScope string

const (
	CouponScopeEvent    CouponScope = "event"
	CouponScopeCalendar CouponScope = "calendar"
)

type CalendarEventTicketCoupon struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	AccountID        *uint              `json:"-"`
	Account          *Account           `gorm:"foreignKey:AccountID;constraint:OnDelete:SET NULL" json:"-"`
	Scope            CouponScope        `gorm:"type:varchar(16);not null" json:"scope"`
	CalendarEventID  *uint              `gorm:"not null;index" json:"calendarEventId,omitempty"`
	CalendarEvent    *CalendarEvent     `gorm:"foreignKey:CalendarEventID;constraint:OnDelete:CASCADE" json:"-"`
	CalendarID       *uint              `gorm:"not null;index" json:"calendarId,omitempty"`
	Calendar         *Calendar          `gorm:"foreignKey:CalendarID;constraint:OnDelete:CASCADE" json:"-"`
	TicketTypeIDs    *[]uint            `json:"ticketTypeIds,omitempty"`
	Code             string             `gorm:"uniqueIndex;size:64;not null" json:"code"`
	Description      *string            `json:"description,omitempty"`
	DiscountType     CouponDiscountType `gorm:"type:varchar(16);not null" json:"discountType"`
	DiscountPercent  *uint              `json:"discountPercent,omitempty"`
	DiscountAmount   *uint              `json:"discountAmount,omitempty"`
	DiscountCurrency *OrderCurrency     `gorm:"type:varchar(8)" json:"discountCurrency,omitempty"`
	SingleUse        bool               `gorm:"default:false" json:"singleUse"`
	MaxRedemptions   *uint              `json:"maxRedemptions,omitempty"`
	Redemptions      uint               `gorm:"default:0" json:"redemptions"`
	StartsAt         time.Time          `json:"startsAt"`
	EndsAt           time.Time          `json:"endsAt"`
	IsActive         bool               `gorm:"default:true" json:"isActive"`
	CreatedAt        time.Time          `json:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt     `gorm:"index" json:"-"`
}
