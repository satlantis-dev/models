package database

import (
	"fmt"

	"gorm.io/gorm"
)

type AccountPlaceRoleType int

const (
	Follower AccountPlaceRoleType = iota + 1
	Visitor
	Inhabitant
	Ambassador
)

type AccountPlaceRole struct {
	AccountID         uint                 `gorm:"index;primaryKey" json:"accountId"`
	Account           AccountDTO           `gorm:"foreignKey:AccountID" json:"account"`
	PlaceID           uint                 `gorm:"index;primaryKey" json:"placeId"`
	Place             *Place               `gorm:"foreignKey:PlaceID" json:"place,omitempty"`
	AmbassadorRequest bool                 `json:"ambassadorRequest"`
	Type              AccountPlaceRoleType `gorm:"not null" json:"type"`
}

// CreateOrUpdateAccountPlaceRole - Create or update a AccountPlaceRole.
func CreateOrUpdateAccountPlaceRole(db *gorm.DB, accountPlaceRole AccountPlaceRole) (*AccountPlaceRole, error) {
	if err := db.Save(&accountPlaceRole).Error; err != nil {
		return nil, err
	}
	return &accountPlaceRole, nil
}

// DeleteAccountPlaceRole - Delete a AccountPlaceRole.
func DeleteAccountPlaceRole(db *gorm.DB, accountID uint, placeID uint) error {
	if err := db.Where("account_id = ? AND place_id = ?", accountID, placeID).Delete(&AccountPlaceRole{}).Error; err != nil {
		return err
	}
	return nil
}

// GetAccountPlaceRoles
func GetAccountPlaceRolesByRoleType(db *gorm.DB, roleType AccountPlaceRoleType) (*[]AccountPlaceRole, error) {
	var accountPlaceRoles []AccountPlaceRole
	if err := db.Where("type = ?", roleType).Preload("Place").Preload("Place.Event").Preload("Place.Event.Tags").Find(&accountPlaceRoles).Error; err != nil {
		return nil, err
	}
	return &accountPlaceRoles, nil
}

// GetAccountPlaceRoleByAccountIDAndPlaceID - Get a AccountPlaceRole by account ID and place ID.
func GetAccountPlaceRoleByAccountIDAndPlaceIDAndType(db *gorm.DB, accountID uint, placeID uint, accountPlaceRoleType AccountPlaceRoleType) (*AccountPlaceRole, error) {
	var accountPlaceRole AccountPlaceRole
	if err := db.Where("account_id = ? AND place_id = ? AND type = ?", accountID, placeID, accountPlaceRoleType).First(&accountPlaceRole).Error; err != nil {
		return nil, err
	}
	return &accountPlaceRole, nil
}

// GetAccountPlaceRoleByAccountIDAndPlaceID
func GetAccountPlaceRoleByAccountIDAndPlaceID(db *gorm.DB, accountID uint, placeID uint) (*AccountPlaceRole, error) {
	var accountPlaceRole AccountPlaceRole
	if err := db.Where("account_id = ? AND place_id = ?", accountID, placeID).First(&accountPlaceRole).Error; err != nil {
		return nil, err
	}
	return &accountPlaceRole, nil
}

// GetResidenciesByAccountIDandType - Get residencies by account ID and type.
func GetResidenciesByAccountIDandType(db *gorm.DB, accountID uint, accountPlaceRoleType AccountPlaceRoleType) ([]AccountPlaceRole, error) {
	var accountPlaceRoles []AccountPlaceRole
	if err := db.Where("account_id = ? AND type = ?", accountID, accountPlaceRoleType).Find(&accountPlaceRoles).Error; err != nil {
		return nil, err
	}
	return accountPlaceRoles, nil
}

// Raise an ambassador request
func RequestAmbassador(db *gorm.DB, accountID uint, placeID uint) error {
	if err := db.Model(&AccountPlaceRole{}).Where("account_id = ? AND place_id = ?", accountID, placeID).Update("ambassador_request", true).Error; err != nil {
		return err
	}
	return nil
}

// Decline an ambassador request
func DeclineAmbassador(db *gorm.DB, accountID uint, placeID uint) error {
	if err := db.Model(&AccountPlaceRole{}).Where("account_id = ? AND place_id = ?", accountID, placeID).Update("ambassador_request", false).Error; err != nil {
		return err
	}
	return nil
}

// ActivateAmbassador - Activate an ambassador.
func ActivateAmbassador(db *gorm.DB, pubKey, placeOSMRef string) error {
	// Construct the raw SQL query
	query := `
		UPDATE account_place_roles
		SET type = ?, ambassador_request = FALSE
		FROM accounts, places
		WHERE accounts.id = account_place_roles.account_id
		AND places.id = account_place_roles.place_id
		AND accounts.pub_key = ?
		AND places.osm_ref = ?
		`

	fmt.Println(query)
	// Execute the raw SQL query
	if err := db.Exec(query, Ambassador, pubKey, placeOSMRef).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// DeactivateAmbassador - Deactivate an ambassador.
func DeactivateAmbassador(db *gorm.DB, pubKey, placeOSMRef string) error {
	// Construct the raw SQL query
	query := `
		UPDATE account_place_roles
		SET type = ?, ambassador_request = FALSE
		FROM accounts, places
		WHERE accounts.id = account_place_roles.account_id
		AND places.id = account_place_roles.place_id
		AND accounts.pub_key = ?
		AND places.osm_ref = ?
		`

	// Execute the raw SQL query
	if err := db.Exec(query, Inhabitant, pubKey, placeOSMRef).Error; err != nil {
		return err
	}

	return nil
}

// GetAccountPlaceRolesByPlaceID - Get account place roles by place ID.
func GetAccountPlaceRolesByPlaceID(db *gorm.DB, placeID uint) (*[]AccountPlaceRole, error) {
	var accountPlaceRoles []AccountPlaceRole
	if err := db.Preload("Account").Where("place_id = ?", placeID).Find(&accountPlaceRoles).Error; err != nil {
		return nil, err
	}
	return &accountPlaceRoles, nil
}
