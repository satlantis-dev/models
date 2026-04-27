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
	OrderCurrencyEUR OrderCurrency = "EUR"
	OrderCurrencyCAD OrderCurrency = "CAD"
	OrderCurrencyGBP OrderCurrency = "GBP"
	OrderCurrencyAUD OrderCurrency = "AUD"
)

type OrderStatus string

var (
	OrderPending   OrderStatus = "pending"
	OrderPaid      OrderStatus = "paid"
	OrderCancelled OrderStatus = "cancelled"
	OrderRefunded  OrderStatus = "refunded"
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

type CalendarEventTicketOrderPayment struct {
	ID                       uint                      `gorm:"primaryKey" json:"id"`
	OrderID                  uint                      `gorm:"uniqueIndex;not null" json:"orderId"`
	Order                    *CalendarEventTicketOrder `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"order,omitempty"`
	PaymentMethod            PaymentMethod             `gorm:"type:varchar(32);not null" json:"paymentMethod"`
	Status                   PaymentStatus             `gorm:"type:varchar(32);default:'pending'" json:"status"`
	Amount                   int64                     `gorm:"not null" json:"amount"` // Cents for fiat, sats for BTC/Lightning
	Currency                 OrderCurrency             `gorm:"type:varchar(8);not null" json:"currency"`
	ExchangeRate             *float64                  `json:"exchangeRate"`
	ExchangeRateSource       *string                   `gorm:"type:varchar(50)" json:"exchangeRateSource"`
	LightningPaymentHash     *string                   `gorm:"uniqueIndex;size:64" json:"lightningPaymentHash,omitempty"`
	LightningPaymentRequest  *string                   `gorm:"type:text" json:"lightningPaymentRequest,omitempty"`
	LightningPreimage        *string                   `gorm:"size:64" json:"lightningPreimage,omitempty"`
	LightningProvider        *string                   `gorm:"size:32" json:"lightningProvider,omitempty"`
	LightningProviderTxID    *string                   `gorm:"index" json:"lightningProviderTxId,omitempty"`
	PaymentProviderReference *string                   `gorm:"index" json:"paymentProviderReference"`
	Metadata                 *datatypes.JSON           `gorm:"type:jsonb" json:"metadata,omitempty"`
	PaidAt                   *time.Time                `json:"paidAt,omitempty"`
	ExpiredAt                *time.Time                `json:"expiredAt,omitempty"`
	ExpiresAt                *time.Time                `json:"expiresAt,omitempty"`
	FailedAt                 *time.Time                `json:"failedAt,omitempty"`
	RefundedAt               *time.Time                `json:"refundedAt,omitempty"`
	CancelledAt              *time.Time                `json:"cancelledAt,omitempty"`
	CreatedAt                time.Time                 `json:"createdAt"`
	UpdatedAt                time.Time                 `json:"updatedAt"`
	DeletedAt                gorm.DeletedAt            `gorm:"index" json:"-"`
}

// CalendarEventTicketOrderRefund represents a refund for a ticket order
type CalendarEventTicketOrderRefund struct {
	ID           uint                      `gorm:"primarykey" json:"id"`
	OrderID      uint                      `gorm:"not null;index" json:"orderId"`
	Order        *CalendarEventTicketOrder `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"-"`
	TicketID     *uint                     `gorm:"index" json:"ticketId,omitempty"`
	Ticket       *CalendarEventTicket      `gorm:"foreignKey:TicketID;constraint:OnDelete:SET NULL" json:"ticket,omitempty"`
	Amount       int64                     `gorm:"not null" json:"amount"`
	Fee          int64                     `gorm:"type:bigint;default:0" json:"fee"`
	Currency     OrderCurrency             `gorm:"not null" json:"currency"`
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

type CalendarEventTicketOrder struct {
	ID              uint                           `gorm:"primaryKey" json:"id"`
	Code            string                         `gorm:"uniqueIndex;size:64" json:"code"`
	CalendarEventID uint                           `gorm:"not null;index" json:"calendarEventId"`
	CalendarEvent   *CalendarEvent                 `gorm:"foreignKey:CalendarEventID;constraint:OnDelete:CASCADE" json:"calendarEvent,omitempty"`
	AccountID       uint                           `gorm:"not null;index" json:"accountId"`
	Account         *Account                       `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"-"`
	TotalPrice      int64                          `gorm:"type:bigint" json:"totalPrice"`
	Currency        OrderCurrency                  `gorm:"type:varchar(8)" json:"currency"`
	RefundedAmount  int64                          `gorm:"type:bigint;default:0" json:"refundedAmount"`
	PriceCurrency   *OrderCurrency                 `gorm:"type:varchar(10)" json:"priceCurrency"`
	PriceAmount     *int64                         `json:"priceAmount"`
	Status          OrderStatus                    `gorm:"type:varchar(32);default:'pending'" json:"status"`
	RsvpData        datatypes.JSON                 `gorm:"type:jsonb" json:"rsvpData,omitempty"`
	Items           []CalendarEventTicketOrderItem `gorm:"foreignKey:OrderID" json:"items,omitempty"`
	CreatedAt       time.Time                      `json:"-"`
	UpdatedAt       time.Time                      `json:"-"`
	DeletedAt       gorm.DeletedAt                 `gorm:"index" json:"-"`
}

type CalendarEventTicketOrderItem struct {
	ID              uint                              `gorm:"primaryKey" json:"id"`
	OrderID         uint                              `gorm:"not null;index" json:"orderId"`
	Order           *CalendarEventTicketOrder         `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"order,omitempty"`
	TicketTypeID    uint                              `gorm:"not null;index" json:"ticketTypeId"`
	TicketType      *CalendarEventTicketType          `gorm:"foreignKey:TicketTypeID;constraint:OnDelete:CASCADE" json:"ticketType,omitempty"`
	Quantity        uint                              `json:"quantity"`
	PriceEach       int64                             `gorm:"type:bigint" json:"priceEach"`
	Currency        OrderCurrency                     `gorm:"type:varchar(8)" json:"currency"`
	RefundedAmount  int64                             `gorm:"type:bigint;default:0" json:"refundedAmount"`
	PriceCurrency   *OrderCurrency                    `gorm:"type:varchar(10)" json:"priceCurrency"`
	PriceAmount     *int64                            `json:"priceAmount"`
	Status          OrderStatus                       `gorm:"type:varchar(32);default:'pending'" json:"status"`
	PromotionID     *uint                             `gorm:"index" json:"promotionId,omitempty"`
	Promotion       *CalendarEventTicketTypePromotion `gorm:"foreignKey:PromotionID;constraint:OnDelete:SET NULL" json:"promotion,omitempty"`
	OriginalPrice   *int64                            `gorm:"type:bigint" json:"originalPrice,omitempty"`
	DiscountPercent *uint                             `json:"discountPercent,omitempty"`
	CreatedAt       time.Time                         `json:"-"`
	UpdatedAt       time.Time                         `json:"-"`
	DeletedAt       gorm.DeletedAt                    `gorm:"index" json:"-"`
}

type CalendarEventTicket struct {
	ID          uint                          `gorm:"primaryKey" json:"id"`
	OrderItemID uint                          `gorm:"not null;index" json:"orderItemId"`
	OrderItem   *CalendarEventTicketOrderItem `gorm:"foreignKey:OrderItemID;constraint:OnDelete:CASCADE" json:"orderItem,omitempty"`
	AccountID   uint                          `gorm:"not null;index" json:"accountId"`
	Account     *Account                      `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"-"`
	RsvpID      *uint                         `gorm:"index" json:"rsvpId,omitempty"`
	RSVP        *CalendarEventRSVP            `gorm:"foreignKey:RsvpID;constraint:OnDelete:CASCADE" json:"rsvp,omitempty"`
	Status      TicketStatus                  `gorm:"type:varchar(32);default:'active'" json:"status"`
	Code        string                        `gorm:"uniqueIndex;size:64" json:"code"`
	CheckedInAt *time.Time                    `json:"checkedInAt,omitempty"`
	CreatedAt   time.Time                     `json:"-"`
	UpdatedAt   time.Time                     `json:"-"`
	DeletedAt   gorm.DeletedAt                `gorm:"index" json:"-"`
}

type CalendarEventTicketOrderPaymentDTO struct {
	ID                       uint                      `json:"id"`
	OrderID                  uint                      `json:"orderId"`
	Order                    *CalendarEventTicketOrder `json:"order,omitempty"`
	PaymentMethod            PaymentMethod             `json:"paymentMethod"`
	Status                   PaymentStatus             `json:"status"`
	Amount                   int64                     `json:"amount"` // Cents for fiat, sats for BTC/Lightning
	Currency                 OrderCurrency             `json:"currency"`
	ExchangeRate             *float64                  `json:"exchangeRate"`
	LightningPreimage        *string                   `json:"lightningPreimage,omitempty"`
	LightningProvider        *string                   `json:"lightningProvider,omitempty"`
	LightningProviderTxID    *string                   `json:"lightningProviderTxId,omitempty"`
	PaymentProviderReference *string                   `json:"paymentProviderReference"`
	PaidAt                   *time.Time                `json:"paidAt,omitempty"`
	ExpiredAt                *time.Time                `json:"expiredAt,omitempty"`
	ExpiresAt                *time.Time                `json:"expiresAt,omitempty"`
	FailedAt                 *time.Time                `json:"failedAt,omitempty"`
	RefundedAt               *time.Time                `json:"refundedAt,omitempty"`
	CancelledAt              *time.Time                `json:"cancelledAt,omitempty"`
}

func (CalendarEventTicketOrderPaymentDTO) TableName() string {
	return "calendar_event_ticket_order_payments"
}

// ToDTO - Convert CalendarEventTicketOrderPayment to CalendarEventTicketOrderPaymentDTO.
func (payment CalendarEventTicketOrderPayment) ToDTO() *CalendarEventTicketOrderPaymentDTO {
	return &CalendarEventTicketOrderPaymentDTO{
		ID:                       payment.ID,
		OrderID:                  payment.OrderID,
		Order:                    payment.Order,
		PaymentMethod:            payment.PaymentMethod,
		Status:                   payment.Status,
		Amount:                   payment.Amount,
		Currency:                 payment.Currency,
		ExchangeRate:             payment.ExchangeRate,
		LightningPreimage:        payment.LightningPreimage,
		LightningProvider:        payment.LightningProvider,
		LightningProviderTxID:    payment.LightningProviderTxID,
		PaymentProviderReference: payment.PaymentProviderReference,
		PaidAt:                   payment.PaidAt,
		ExpiredAt:                payment.ExpiredAt,
		ExpiresAt:                payment.ExpiresAt,
		RefundedAt:               payment.RefundedAt,
		CancelledAt:              payment.CancelledAt,
	}
}
