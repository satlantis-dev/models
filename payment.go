package models

type OrderCurrency string

const (
	OrderCurrencyBTC OrderCurrency = "BTC"
	OrderCurrencyUSD OrderCurrency = "USD"
	OrderCurrencyEUR OrderCurrency = "EUR"
	OrderCurrencyCAD OrderCurrency = "CAD"
	OrderCurrencyGBP OrderCurrency = "GBP"
	OrderCurrencyAUD OrderCurrency = "AUD"
	OrderCurrencyNZD OrderCurrency = "NZD"
)

type PaymentMethod string

const (
	PaymentMethodLightning PaymentMethod = "lightning"
	PaymentMethodStripe    PaymentMethod = "stripe"
	PaymentMethodOnchain   PaymentMethod = "onchain"
	PaymentMethodOffline   PaymentMethod = "offline"
)

type OfflinePaymentMethod string

const (
	OfflinePaymentMethodCash         OfflinePaymentMethod = "cash"
	OfflinePaymentMethodCard         OfflinePaymentMethod = "card"
	OfflinePaymentMethodBankTransfer OfflinePaymentMethod = "transfer"
	OfflinePaymentMethodCheque       OfflinePaymentMethod = "cheque"
	OfflinePaymentMethodCrypto       OfflinePaymentMethod = "crypto"
	OfflinePaymentMethodPix          OfflinePaymentMethod = "pix"
	OfflinePaymentMethodAppAlipay    OfflinePaymentMethod = "alipay"
	OfflinePaymentMethodAppRevolut   OfflinePaymentMethod = "revolut"
	OfflinePaymentMethodAppWechat    OfflinePaymentMethod = "wechat"
	OfflinePaymentMethodAppWise      OfflinePaymentMethod = "wise"
	OfflinePaymentMethodAppZelle     OfflinePaymentMethod = "zelle"
	OfflinePaymentMethodOther        OfflinePaymentMethod = "other"
)

// AllOfflinePaymentMethods lists every recognized OfflinePaymentMethod value,
// so callers validating user input can check against this instead of
// duplicating the list (and risking drift as new methods are added above).
var AllOfflinePaymentMethods = []OfflinePaymentMethod{
	OfflinePaymentMethodCash,
	OfflinePaymentMethodCard,
	OfflinePaymentMethodBankTransfer,
	OfflinePaymentMethodCheque,
	OfflinePaymentMethodCrypto,
	OfflinePaymentMethodPix,
	OfflinePaymentMethodAppAlipay,
	OfflinePaymentMethodAppRevolut,
	OfflinePaymentMethodAppWechat,
	OfflinePaymentMethodAppWise,
	OfflinePaymentMethodAppZelle,
	OfflinePaymentMethodOther,
}

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
