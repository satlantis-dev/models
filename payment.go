package models

type PaymentMethod string

const (
	PaymentMethodLightning PaymentMethod = "lightning"
	PaymentMethodStripe    PaymentMethod = "stripe"
	PaymentMethodOnchain   PaymentMethod = "onchain"
	PaymentMethodOffline   PaymentMethod = "offline"
)

type OfflinePaymentMethod string

const (
	OfflinePaymentMethodCash    OfflinePaymentMethod = "cash"
	OfflinePaymentMethodCheck   OfflinePaymentMethod = "check"
	OfflinePaymentMethodCrypto  OfflinePaymentMethod = "crypto"
	OfflinePaymentMethodRevolut OfflinePaymentMethod = "revolut"
	OfflinePaymentMethodAlipay  OfflinePaymentMethod = "alipay"
	OfflinePaymentMethodWechat  OfflinePaymentMethod = "wechat"
	OfflinePaymentMethodPix     OfflinePaymentMethod = "pix"
	OfflinePaymentMethodOther   OfflinePaymentMethod = "other"
)

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "pending"
	PaymentSubmitted PaymentStatus = "submitted"
	PaymentPaid      PaymentStatus = "paid"
	PaymentExpired   PaymentStatus = "expired"
	PaymentFailed    PaymentStatus = "failed"
	PaymentRefunded  PaymentStatus = "refunded"
	PaymentCancelled PaymentStatus = "cancelled"
)

type RefundStatus string

const (
	RefundPending    RefundStatus = "pending"
	RefundProcessing RefundStatus = "processing"
	RefundCompleted  RefundStatus = "completed"
	RefundFailed     RefundStatus = "failed"
)

type RefundMethod string

const (
	RefundMethodLightning       RefundMethod = "lightning"
	RefundMethodSatlantisWallet RefundMethod = "satlantis_wallet"
	RefundMethodStripe          RefundMethod = "stripe"
	RefundMethodOffline         RefundMethod = "offline"
)
