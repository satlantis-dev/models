package models

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type AccountType int

const (
	BasicType AccountType = iota + 1
	AmbassadorType
	FounderType
	AdminType
)

type Account struct {
	ID                          uint                    `gorm:"primaryKey" json:"id"`
	CreatedAt                   time.Time               `json:"-"`
	About                       string                  `gorm:"type:text" json:"about"`
	AccountPlaceRoles           []AccountPlaceRole      `gorm:"foreignKey:AccountID" json:"accountPlaceRoles"`
	AccountType                 AccountType             `gorm:"type:smallint" json:"accountType"`
	AuthDetails                 []AuthenticationDetail  `gorm:"foreignKey:AccountID" json:"authDetails"`
	Banner                      string                  `gorm:"type:text" json:"banner"`
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
	Interests                   *string                 `gorm:"type:jsonb" json:"interests"`
	IsBusiness                  bool                    `json:"isBusiness"`
	LastSeen                    *time.Time              `json:"-"`
	LocationRatings             []AccountLocationRating `gorm:"foreignKey:AccountID" json:"locationRatings"`
	Lud06                       string                  `gorm:"default:NULL" json:"lud06"`
	Lud16                       string                  `gorm:"default:NULL" json:"lud16"`
	Name                        string                  `gorm:"type:text" json:"name"`
	Nip05                       string                  `gorm:"default:NULL" json:"nip05"`
	Notes                       []Note                  `gorm:"foreignKey:AccountID" json:"notes"`
	Npub                        string                  `gorm:"uniqueIndex" json:"npub"`
	Picture                     string                  `gorm:"type:text" json:"picture"`
	Phone                       string                  `json:"phone"`
	PlaceRatings                []AccountPlaceRating    `gorm:"foreignKey:AccountID" json:"placeRatings"`
	PubKey                      string                  `gorm:"uniqueIndex;default:NULL" json:"pubKey"`
	SocialMediaList             []SocialMedia           `gorm:"foreignKey:AccountID" json:"socialMediaList"`
	Website                     string                  `gorm:"type:text" json:"website"`
	ResetPasswordToken          *string                 `gorm:"type:text" json:"resetPasswordToken"`
	ResetPasswordTokenExpiresAt *time.Time              `json:"-"`
	Locations                   []LocationAccount       `gorm:"foreignKey:AccountID" json:"locations"`
}

type AccountPortable struct {
	ID                uint                    `json:"id"`
	About             string                  `json:"about"`
	AccountPlaceRoles []AccountPlaceRole      `gorm:"foreignKey:AccountID" json:"accountPlaceRoles"`
	AccountType       AccountType             `json:"accountType"`
	Banner            string                  `json:"banner"`
	ChatMemberships   []ChatMembership        `gorm:"foreignKey:AccountID" json:"chatMemberships"`
	CurrencyID        *uint                   `json:"currencyId"`
	DisplayName       string                  `json:"displayName"`
	Email             string                  `json:"email"`
	InfluenceScore    uint                    `json:"influenceScore"`
	Interests         *string                 `gorm:"type:jsonb" json:"interests"`
	IsBusiness        bool                    `json:"isBusiness"`
	LastSeen          *time.Time              `json:"-"`
	LocationRatings   []AccountLocationRating `gorm:"foreignKey:AccountID" json:"locationRatings"`
	Lud06             string                  `gorm:"default:NULL" json:"lud06"`
	Lud16             string                  `gorm:"default:NULL" json:"lud16"`
	Name              string                  `json:"name"`
	Nip05             string                  `gorm:"default:NULL" json:"nip05"`
	Notes             []Note                  `gorm:"foreignKey:AccountID" json:"notes"`
	Npub              string                  `json:"npub"`
	Picture           string                  `json:"picture"`
	PlaceRatings      []AccountPlaceRating    `gorm:"foreignKey:AccountID" json:"placeRatings"`
	PubKey            string                  `json:"pubKey"`
	Website           string                  `json:"website"`
	Following         []AccountDTO            `json:"following"`
	FollowedBy        []AccountDTO            `json:"followedBy"`
}

type AccountDTO struct {
	ID          uint        `json:"id"`
	About       string      `json:"about"`
	AccountType AccountType `json:"accountType"`
	IsBusiness  bool        `json:"isBusiness"`
	Name        string      `json:"name"`
	Nip05       string      `gorm:"default:NULL" json:"nip05"`
	Npub        string      `json:"npub"`
	Picture     string      `json:"picture"`
	PubKey      string      `json:"pubKey"`
}

func (AccountDTO) TableName() string {
	return "accounts"
}

func (at *AccountType) Scan(value interface{}) error {
	*at = AccountType(value.(int64))
	return nil
}

func (at AccountType) Value() (driver.Value, error) {
	return int64(at), nil
}

func (a *Account) ToDTO() AccountDTO {
	return AccountDTO{
		ID:          a.ID,
		About:       a.About,
		AccountType: a.AccountType,
		IsBusiness:  a.IsBusiness,
		Name:        a.Name,
		Nip05:       a.Nip05,
		Npub:        a.Npub,
		Picture:     a.Picture,
		PubKey:      a.PubKey,
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
		AccountType:       a.AccountType,
		Banner:            a.Banner,
		AccountPlaceRoles: a.AccountPlaceRoles,
		ChatMemberships:   a.ChatMemberships,
		CurrencyID:        a.CurrencyID,
		DisplayName:       a.DisplayName,
		Email:             a.Email,
		InfluenceScore:    a.InfluenceScore,
		Interests:         a.Interests,
		IsBusiness:        a.IsBusiness,
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
