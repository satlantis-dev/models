package models

import "time"

// CommunityMemberAuditAction identifies what an admin-initiated change to a
// CommunityMember did.
type CommunityMemberAuditAction string

const (
	// CommunityMemberAuditActionTierSet covers every admin-driven tier
	// assignment that doesn't go through the member-facing request flow:
	// adding a member (with or without a tier) and a bulk tier upgrade.
	CommunityMemberAuditActionTierSet      CommunityMemberAuditAction = "tier_set"
	CommunityMemberAuditActionRemoved      CommunityMemberAuditAction = "removed"
	CommunityMemberAuditActionBanned       CommunityMemberAuditAction = "banned"
	CommunityMemberAuditActionMadeProspect CommunityMemberAuditAction = "made_prospect"
)

// CommunityMemberAuditLog records an admin-initiated change to a
// CommunityMember that bypasses CommunityMembershipRequest entirely (direct
// tier assignment, removal, ban, or demotion to prospect) - it's the audit
// trail for exactly those actions, so they remain reconstructable even
// though no request/subscription row exists for them.
type CommunityMemberAuditLog struct {
	ID                   uint                       `gorm:"primaryKey;autoIncrement" json:"id"`
	CommunityID          uint                       `gorm:"not null;index" json:"communityId"`
	AccountID            uint                       `gorm:"not null;index" json:"accountId"`
	Action               CommunityMemberAuditAction `gorm:"type:varchar(32);not null;index" json:"action"`
	OldTierID            *uint                      `gorm:"index" json:"oldTierId,omitempty"`
	OldTier              *CommunityMembershipTier   `gorm:"foreignKey:OldTierID;constraint:OnDelete:SET NULL;" json:"oldTier,omitempty"`
	NewTierID            *uint                      `gorm:"index" json:"newTierId,omitempty"`
	NewTier              *CommunityMembershipTier   `gorm:"foreignKey:NewTierID;constraint:OnDelete:SET NULL;" json:"newTier,omitempty"`
	PerformedByAccountID *uint                      `gorm:"index" json:"performedByAccountId,omitempty"`
	PerformedByAccount   *AccountDTO                `gorm:"foreignKey:PerformedByAccountID;constraint:OnDelete:SET NULL;" json:"performedByAccount,omitempty"`
	CreatedAt            time.Time                  `gorm:"autoCreateTime;index" json:"createdAt"`
}

func (CommunityMemberAuditLog) TableName() string {
	return "community_member_audit_logs"
}
