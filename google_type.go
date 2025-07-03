package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONBMap map[string]string

func (m JSONBMap) Value() (driver.Value, error) {
	if m == nil {
		return "{}", nil
	}
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *JSONBMap) Scan(value interface{}) error {
	if value == nil {
		*m = make(JSONBMap)
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return errors.New("unsupported type for MapStringString")
	}
	return json.Unmarshal(bytes, m)
}

type GoogleType struct {
	Name        string    `gorm:"primaryKey" json:"name"`
	OSMStdTag   *JSONBMap `gorm:"type:jsonb" json:"osmStdTag"`
	OSMExtraTag *JSONBMap `gorm:"type:jsonb" json:"osmExtraTag"`
}

func (GoogleType) TableName() string {
	return "google_types"
}
