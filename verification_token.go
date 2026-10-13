package models

import (
	"time"
)

// VerificationTokenPurpose is what a pending-auth token was issued for.
// Minting a token supersedes only the account's live tokens with the same
// purpose and scope, so unrelated links (an invite, a payment link, a login
// code) stay valid alongside each other.
type VerificationTokenPurpose string

const (
	// VerificationTokenPurposeLogin is an OTP login or email verification code.
	VerificationTokenPurposeLogin VerificationTokenPurpose = "login"
	// VerificationTokenPurposeEmailLink is a login link in a transactional
	// email (event RSVP verification/confirmation, calendar subscription).
	VerificationTokenPurposeEmailLink VerificationTokenPurpose = "email_link"
	// VerificationTokenPurposeInvite is a login link in an invitation
	// (community member/admin, event attendee/cohost/speaker, calendar contributor).
	VerificationTokenPurposeInvite VerificationTokenPurpose = "invite"
	// VerificationTokenPurposePaymentLink is a login link in a community
	// membership payment link email.
	VerificationTokenPurposePaymentLink VerificationTokenPurpose = "payment_link"
)

type VerificationToken struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `gorm:"index" json:"-"`
	Token       string     `gorm:"type:text" json:"token"`
	AccountID   uint       `gorm:"index;index:idx_verification_token_account_purpose_scope,priority:1" json:"accountId"`
	OTPHash     string     `gorm:"type:text" json:"otpHash"`
	OTPAttempts uint       `gorm:"default:0" json:"otpAttempts"`
	// Purpose and Scope decide which earlier tokens a new one supersedes.
	// Scope names what the token is about within its purpose (e.g.
	// "community:42", "event:7:cohost", "subscription:1873"); empty for login.
	// Tokens issued before these columns existed have both empty and are never
	// superseded, only expire.
	Purpose VerificationTokenPurpose `gorm:"type:varchar(32);not null;default:'';index:idx_verification_token_account_purpose_scope,priority:2" json:"-"`
	Scope   string                   `gorm:"type:varchar(64);not null;default:'';index:idx_verification_token_account_purpose_scope,priority:3" json:"-"`
}
