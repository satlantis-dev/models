package models

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `gorm:"index" json:"-,omitempty"`
	Code        string     `gorm:"type:char(2);uniqueIndex" json:"code"`
	Code_3      string     `gorm:"type:char(3);uniqueIndex" json:"code3"`
	Name        string     `gorm:"type:text" json:"name"`
	ContinentID uint       `json:"-"`
	Continent   Continent  `json:"-"`
}

// CreateOrUpdateCountry - Create or update a country.
func CreateOrUpdateCountry(db *gorm.DB, country Country) (*Country, error) {
	if err := db.Save(&country).Error; err != nil {
		return nil, err
	}
	return &country, nil
}

// GetCountry - Get a country by ID.
func GetCountry(db *gorm.DB, id uint) (*Country, error) {
	var country Country
	if err := db.Preload("Continent").First(&country, id).Error; err != nil {
		return nil, err
	}
	return &country, nil
}

// GetCountryByCode - Get a country by code.
func GetCountryByCode(db *gorm.DB, code string) (*Country, error) {
	var country Country
	if err := db.Preload("Continent").Where("code = ?", code).First(&country).Error; err != nil {
		return nil, err
	}
	return &country, nil
}

// GetCountryByName - Get a country by name.
func GetCountryByName(db *gorm.DB, name string) (*Country, error) {
	var country Country
	if err := db.Preload("Continent").Where("name = ?", name).First(&country).Error; err != nil {
		return nil, err
	}
	return &country, nil
}

// Get all countries.
func GetCountries(db *gorm.DB) (*[]Country, error) {
	var countries []Country
	if err := db.Preload("Continent").Find(&countries).Error; err != nil {
		return nil, err
	}
	return &countries, nil
}

// DeleteCountry - Delete a country.
func DeleteCountry(db *gorm.DB, country Country) error {
	if err := db.Delete(&country).Error; err != nil {
		return err
	}
	return nil
}
