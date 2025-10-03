package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
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

type CalendarEventTicketType struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	CalendarEventID uint           `gorm:"not null;index" json:"calendarEventId"`
	Name            string         `json:"name"`
	Description     *string        `json:"description"`
	PriceSats       *int64         `gorm:"type:bigint" json:"priceSats"`
	PriceFiat       *int64         `gorm:"type:bigint" json:"priceFiat"`
	FiatCurrency    *OrderCurrency `gorm:"type:varchar(8)" json:"fiatCurrency"`
	MaxCapacity     *uint          `json:"maxCapacity"`
	SellStartDate   *time.Time     `json:"sellStartDate"`
	SellEndDate     *time.Time     `json:"sellEndDate"`
	CreatedByID     *uint          `json:"-"`
	CreatedBy       *Account       `json:"-"`
	CreatedAt       time.Time      `json:"createdAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

type CalendarEventTicketOrder struct {
	ID              uint                    `gorm:"primaryKey" json:"id"`
	CalendarEventID uint                    `gorm:"not null;index" json:"calendarEventId"`
	AccountID       uint                    `gorm:"not null;index" json:"accountId"`
	TotalPrice      int64                   `gorm:"type:bigint" json:"totalPrice"`
	Currency        OrderCurrency           `gorm:"type:varchar(8)" json:"currency"`
	RefundedAmount  int64                   `gorm:"type:bigint;default:0" json:"refundedAmount"`
	Status          OrderStatus             `gorm:"type:varchar(32);default:'pending'" json:"status"`
	RsvpData        *map[string]interface{} `gorm:"type:jsonb" json:"rsvpData,omitempty"`
	CreatedAt       time.Time               `json:"-"`
	UpdatedAt       time.Time               `json:"-"`
	DeletedAt       gorm.DeletedAt          `gorm:"index" json:"-"`
}

type CalendarEventTicketOrderItem struct {
	ID             uint                     `gorm:"primaryKey" json:"id"`
	OrderID        uint                     `gorm:"not null;index" json:"orderId"`
	Order          CalendarEventTicketOrder `json:"order"`
	TicketTypeID   uint                     `gorm:"not null;index" json:"ticketTypeId"`
	TicketType     CalendarEventTicketType  `json:"ticketType"`
	Quantity       uint                     `json:"quantity"`
	PriceEach      int64                    `gorm:"type:bigint" json:"priceEach"`
	Currency       OrderCurrency            `gorm:"type:varchar(8)" json:"currency"`
	RefundedAmount int64                    `gorm:"type:bigint;default:0" json:"refundedAmount"`
	Status         OrderStatus              `gorm:"type:varchar(32);default:'pending'" json:"status"`
	CreatedAt      time.Time                `json:"-"`
	UpdatedAt      time.Time                `json:"-"`
	DeletedAt      gorm.DeletedAt           `gorm:"index" json:"-"`
}

type CalendarEventTicket struct {
	ID          uint                         `gorm:"primaryKey" json:"id"`
	OrderItemID uint                         `gorm:"not null;index" json:"orderItemId"`
	OrderItem   CalendarEventTicketOrderItem `json:"orderItem"`
	AccountID   uint                         `gorm:"not null;index" json:"accountId"`
	RsvpID      *uint                        `gorm:"index" json:"rsvpId,omitempty"`
	RSVP        *CalendarEventRSVP           `gorm:"foreignKey:RsvpID" json:"rsvp,omitempty"`
	Status      TicketStatus                 `gorm:"type:varchar(32);default:'active'" json:"status"`
	Code        string                       `gorm:"uniqueIndex;size:64" json:"code"`
	CreatedAt   time.Time                    `json:"-"`
	UpdatedAt   time.Time                    `json:"-"`
	DeletedAt   gorm.DeletedAt               `gorm:"index" json:"-"`
}

type CalendarEventTicketOrderPayment struct {
	ID                      uint                     `gorm:"primaryKey" json:"id"`
	OrderID                 uint                     `gorm:"uniqueIndex;not null" json:"orderId"`
	Order                   CalendarEventTicketOrder `json:"order"`
	PaymentMethod           PaymentMethod            `gorm:"type:varchar(32);not null" json:"paymentMethod"`
	Status                  PaymentStatus            `gorm:"type:varchar(32);default:'pending'" json:"status"`
	Amount                  int64                    `gorm:"not null" json:"amount"` // Cents for fiat, sats for BTC/Lightning
	Currency                OrderCurrency            `gorm:"type:varchar(8);not null" json:"currency"`
	LightningPaymentHash    *string                  `gorm:"uniqueIndex;size:64" json:"lightningPaymentHash,omitempty"`
	LightningPaymentRequest *string                  `gorm:"type:text" json:"lightningPaymentRequest,omitempty"`
	LightningPreimage       *string                  `gorm:"size:64" json:"lightningPreimage,omitempty"`
	LightningProvider       *string                  `gorm:"size:32" json:"lightningProvider,omitempty"`
	LightningProviderTxID   *string                  `gorm:"index" json:"lightningProviderTxId,omitempty"`
	Metadata                *datatypes.JSON          `gorm:"type:jsonb" json:"metadata,omitempty"`
	PaidAt                  *time.Time               `json:"paidAt,omitempty"`
	ExpiredAt               *time.Time               `json:"expiredAt,omitempty"`
	ExpiresAt               *time.Time               `json:"expiresAt,omitempty"`
	FailedAt                *time.Time               `json:"failedAt,omitempty"`
	RefundedAt              *time.Time               `json:"refundedAt,omitempty"`
	CancelledAt             *time.Time               `json:"cancelledAt,omitempty"`
	CreatedAt               time.Time                `json:"createdAt"`
	UpdatedAt               time.Time                `json:"updatedAt"`
	DeletedAt               gorm.DeletedAt           `gorm:"index" json:"-"`
}
