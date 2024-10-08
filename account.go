package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID                          uint                    `gorm:"primaryKey" json:"id"`
	CreatedAt                   time.Time               `json:"-"`
	About                       string                  `gorm:"type:text" json:"about"`
	AccountPlaceRoles           []AccountPlaceRole      `gorm:"foreignKey:AccountID" json:"accountPlaceRoles"`
	AuthDetails                 []AuthenticationDetail  `gorm:"foreignKey:AccountID" json:"authDetails"`
	Banner                      string                  `gorm:"type:text" json:"banner"`
	BusinessCategory            string                  `gorm:"default:NULL" json:"businessCategory"`
	ChatMemberships             []ChatMembership        `gorm:"foreignKey:AccountID" json:"chatMemberships"`
	CurrencyID                  *uint                   `gorm:"index" json:"currencyId"`
	Currency                    Currency                `json:"currency"`
	DisplayName                 string                  `gorm:"type:text" json:"displayName"`
	Email                       string                  `gorm:"default:NULL" json:"email"`
	EmailVerified               bool                    `json:"emailVerified"`
	Experiences                 []Experience            `gorm:"foreignKey:AccountID" json:"experiences"`
	Following                   []Follow                `gorm:"foreignkey:FollowerID" json:"following"`
	FollowedBy                  []Follow                `gorm:"foreignkey:FollowerID" json:"followedBy"`
	InfluenceScore              uint                    `json:"influenceScore"`
	Interests                   []Interest              `gorm:"many2many:account_interests" json:"interests"`
	IsAdmin                     bool                    `json:"isAdmin"`
	IsBlacklisted               bool                    `json:"isBlacklisted"`
	IsBusiness                  bool                    `json:"isBusiness"`
	LastSeen                    *time.Time              `json:"-"`
	LocationSetEventID          *uint                   `json:"locationSetEventId"`
	LocationSetEvent            Event                   `gorm:"foreignKey:LocationSetEventID" json:"locationSetEvent"`
	LocationRatings             []AccountLocationRating `gorm:"foreignKey:AccountID" json:"locationRatings"`
	Locations                   []LocationAccount       `gorm:"foreignKey:AccountID" json:"locations"`
	Lud06                       string                  `gorm:"default:NULL" json:"lud06"`
	Lud16                       string                  `gorm:"default:NULL" json:"lud16"`
	Name                        string                  `gorm:"type:text" json:"name"`
	Nip05                       string                  `gorm:"default:NULL" json:"nip05"`
	Notes                       []Note                  `gorm:"foreignKey:AccountID" json:"notes"`
	Npub                        string                  `gorm:"uniqueIndex;default:NULL" json:"npub"`
	Password                    string                  `gorm:"type:text" json:"password"`
	Picture                     string                  `gorm:"type:text" json:"picture"`
	Phone                       string                  `json:"phone"`
	PlaceRatings                []AccountPlaceRating    `gorm:"foreignKey:AccountID" json:"placeRatings"`
	PrivateKey                  string                  `json:"privateKey"`
	PubKey                      string                  `gorm:"uniqueIndex;default:NULL" json:"pubKey"`
	Relays                      []Relay                 `gorm:"foreignKey:AccountID" json:"relays"`
	ResetPasswordToken          *string                 `gorm:"type:text" json:"resetPasswordToken"`
	ResetPasswordTokenExpiresAt *time.Time              `json:"-"`
	SocialMediaList             []SocialMedia           `gorm:"foreignKey:AccountID" json:"socialMediaList"`
	Website                     string                  `gorm:"type:text" json:"website"`
	Username                    string                  `gorm:"uniqueIndex;default:NULL" json:"username"`
}

type AccountPortable struct {
	ID                 uint                    `json:"id"`
	About              string                  `json:"about"`
	AccountPlaceRoles  []AccountPlaceRole      `gorm:"foreignKey:AccountID" json:"accountPlaceRoles"`
	Banner             string                  `json:"banner"`
	ChatMemberships    []ChatMembership        `gorm:"foreignKey:AccountID" json:"chatMemberships"`
	CurrencyID         *uint                   `json:"currencyId"`
	DisplayName        string                  `json:"displayName"`
	Email              string                  `json:"email"`
	EmailVerified      bool                    `json:"emailVerified"`
	InfluenceScore     uint                    `json:"influenceScore"`
	Interests          []Interest              `gorm:"many2many:account_interests" json:"interests"`
	IsAdmin            bool                    `json:"isAdmin"`
	IsBusiness         bool                    `json:"isBusiness"`
	LastSeen           *time.Time              `json:"-"`
	LocationSetEventID *uint                   `json:"locationSetEventId"`
	LocationSetEvent   Event                   `gorm:"foreignKey:LocationSetEventID" json:"locationSetEvent"`
	LocationRatings    []AccountLocationRating `gorm:"foreignKey:AccountID" json:"locationRatings"`
	Locations          []LocationAccount       `gorm:"foreignKey:AccountID" json:"locations"`
	Lud06              string                  `gorm:"default:NULL" json:"lud06"`
	Lud16              string                  `gorm:"default:NULL" json:"lud16"`
	Name               string                  `json:"name"`
	Nip05              string                  `gorm:"default:NULL" json:"nip05"`
	Notes              []Note                  `gorm:"foreignKey:AccountID" json:"notes"`
	Npub               string                  `json:"npub"`
	Picture            string                  `json:"picture"`
	PlaceRatings       []AccountPlaceRating    `gorm:"foreignKey:AccountID" json:"placeRatings"`
	PubKey             string                  `json:"pubKey"`
	Website            string                  `json:"website"`
	Following          []AccountDTO            `json:"following"`
	FollowedBy         []AccountDTO            `json:"followedBy"`
}

type AccountDTO struct {
	ID         uint   `json:"id"`
	About      string `json:"about"`
	IsAdmin    bool   `json:"isAdmin"`
	IsBusiness bool   `json:"isBusiness"`
	Name       string `json:"name"`
	Nip05      string `gorm:"default:NULL" json:"nip05"`
	Npub       string `json:"npub"`
	Picture    string `json:"picture"`
	PubKey     string `json:"pubKey"`
}

type SearchAccountDTO struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	DisplayName    string `json:"display_name"`
	About          string `json:"about"`
	Picture        string `json:"picture"`
	Npub           string `json:"npub"`
	FollowersCount int64  `json:"followers_count"`
}

func (AccountDTO) TableName() string {
	return "accounts"
}

func (a *Account) ToDTO() AccountDTO {
	return AccountDTO{
		ID:         a.ID,
		About:      a.About,
		IsAdmin:    a.IsAdmin,
		IsBusiness: a.IsBusiness,
		Name:       a.Name,
		Nip05:      a.Nip05,
		Npub:       a.Npub,
		Picture:    a.Picture,
		PubKey:     a.PubKey,
	}
}

func (a *Account) ToPortableProfile(db *gorm.DB) (*AccountPortable, error) {
	following, err := a.GetFollowingAccounts(db, a.ID)
	if err != nil {
		return nil, err
	}

	followedBy, err := a.GetFollowedByAccounts(db, a.ID)
	if err != nil {
		return nil, err
	}

	return &AccountPortable{
		ID:                a.ID,
		About:             a.About,
		Banner:            a.Banner,
		AccountPlaceRoles: a.AccountPlaceRoles,
		ChatMemberships:   a.ChatMemberships,
		CurrencyID:        a.CurrencyID,
		DisplayName:       a.DisplayName,
		Email:             a.Email,
		EmailVerified:     a.EmailVerified,
		InfluenceScore:    a.InfluenceScore,
		Interests:         a.Interests,
		IsAdmin:           a.IsAdmin,
		IsBusiness:        a.IsBusiness,
		LocationSetEvent:  a.LocationSetEvent,
		Locations:         a.Locations,
		Lud06:             a.Lud06,
		Lud16:             a.Lud16,
		Name:              a.Name,
		Nip05:             a.Nip05,
		Notes:             a.Notes,
		Npub:              a.Npub,
		Picture:           a.Picture,
		PubKey:            a.PubKey,
		Website:           a.Website,
		Following:         following,
		FollowedBy:        followedBy,
	}, nil
}

func (a *Account) GetFollowingAccounts(db *gorm.DB, followerID uint) ([]AccountDTO, error) {
	var accounts []Account
	err := db.Table("follows").Select("accounts.*").
		Joins("join accounts on accounts.id = follows.following_id").
		Where("follows.follower_id = ?", followerID).
		Scan(&accounts).Error
	if err != nil {
		return nil, err
	}

	accountDTOs := make([]AccountDTO, len(accounts))

	for i, account := range accounts {
		accountDTOs[i] = account.ToDTO()
	}

	return accountDTOs, nil
}

func (a *Account) GetFollowersCount(db *gorm.DB, followingID uint) (int64, error) {
	var count int64
	err := db.Table("follows").Select("accounts.*").
		Joins("join accounts on accounts.id = follows.follower_id").
		Where("follows.following_id = ?", followingID).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (a *Account) GetFollowedByAccounts(db *gorm.DB, followingID uint) ([]AccountDTO, error) {
	var accounts []Account
	err := db.Table("follows").Select("accounts.*").
		Joins("join accounts on accounts.id = follows.follower_id").
		Where("follows.following_id = ?", followingID).
		Scan(&accounts).Error
	if err != nil {
		return nil, err
	}

	accountDTOs := make([]AccountDTO, len(accounts))

	for i, account := range accounts {
		accountDTOs[i] = account.ToDTO()
	}

	return accountDTOs, nil
}

func (a *Account) ToSearchAccountDTO(db *gorm.DB) (*SearchAccountDTO, error) {
	// Get the followers count
	followersCount, err := a.GetFollowersCount(db, a.ID)
	if err != nil {
		return nil, err
	}

	return &SearchAccountDTO{
		ID:             a.ID,
		Username:       a.Username,
		DisplayName:    a.DisplayName,
		About:          a.About,
		Picture:        a.Picture,
		Npub:           a.Npub,
		FollowersCount: followersCount,
	}, nil
}
