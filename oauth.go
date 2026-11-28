package models

import (
	"time"

	"github.com/lib/pq"
)

// OAuth 2.1 for MCP clients (Claude and other AI assistants). These tokens
// are separate from session JWTs: they are opaque, scoped, short-lived and
// only accepted on the api's /mcp route.

// OAuthClient is an MCP client registered through dynamic client
// registration (RFC 7591). All clients are public (PKCE, no secret).
type OAuthClient struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"-"`
	ClientID     string         `gorm:"type:varchar(64);not null;uniqueIndex" json:"clientId"`
	ClientName   string         `gorm:"type:text;not null;default:''" json:"clientName"`
	ClientURI    string         `gorm:"type:text;not null;default:''" json:"clientUri"`
	RedirectURIs pq.StringArray `gorm:"type:text[];not null" json:"redirectUris"`
}

func (OAuthClient) TableName() string {
	return "oauth_clients"
}

// OAuthGrant is one account's consent to one client, made on web-app's
// consent page. Revoking it (RevokedAt) kills every token issued under it.
type OAuthGrant struct {
	ID            uint         `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time    `json:"createdAt"`
	UpdatedAt     time.Time    `json:"-"`
	AccountID     uint         `gorm:"not null;index" json:"accountId"`
	Account       *AccountDTO  `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"-"`
	OAuthClientID uint         `gorm:"not null;index" json:"oauthClientId"`
	OAuthClient   *OAuthClient `gorm:"foreignKey:OAuthClientID;constraint:OnDelete:CASCADE;" json:"client,omitempty"`
	// Scopes is space-separated, as in OAuth (e.g. "events:read").
	Scopes string `gorm:"type:varchar(255);not null;default:''" json:"scopes"`
	// Resource is the RFC 8707 resource the tokens are bound to (the /mcp URL).
	Resource   string     `gorm:"type:text;not null;default:''" json:"resource"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	RevokedAt  *time.Time `gorm:"index" json:"-"`
}

func (OAuthGrant) TableName() string {
	return "oauth_grants"
}

type OAuthTokenKind string

const (
	OAuthTokenKindAuthorizationCode OAuthTokenKind = "authorization_code"
	OAuthTokenKindAccess            OAuthTokenKind = "access_token"
	OAuthTokenKindRefresh           OAuthTokenKind = "refresh_token"
)

// OAuthToken is an authorization code, access token or refresh token. Only
// the SHA-256 of the token is stored.
type OAuthToken struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	OAuthGrantID uint           `gorm:"not null;index" json:"oauthGrantId"`
	OAuthGrant   *OAuthGrant    `gorm:"foreignKey:OAuthGrantID;constraint:OnDelete:CASCADE;" json:"-"`
	Kind         OAuthTokenKind `gorm:"type:varchar(32);not null" json:"kind"`
	TokenHash    string         `gorm:"type:varchar(64);not null;uniqueIndex" json:"-"`
	ExpiresAt    time.Time      `gorm:"not null;index" json:"expiresAt"`
	// UsedAt marks a redeemed code or a rotated refresh token; presenting
	// either again revokes the grant (replay).
	UsedAt    *time.Time `json:"-"`
	RevokedAt *time.Time `json:"-"`
	// PKCE challenge (S256) and redirect URI, for authorization codes only.
	CodeChallenge string `gorm:"type:varchar(128);not null;default:''" json:"-"`
	RedirectURI   string `gorm:"type:text;not null;default:''" json:"-"`
}

func (OAuthToken) TableName() string {
	return "oauth_tokens"
}
