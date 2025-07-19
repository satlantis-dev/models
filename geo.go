package models

type GeoLocation struct {
	Features []struct {
		Properties struct {
			TimeZone TimeZone `json:"timezone"`
		} `json:"properties"`
	} `json:"features"`
}

type TimeZone struct {
	Name   string  `json:"name"`
	Offset float64 `json:"offset"`
}
