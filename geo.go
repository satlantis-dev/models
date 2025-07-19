package models

type GeoLocation struct {
	Features []struct {
		Properties struct {
			TimeZone TimeZone `json:"timezone"`
		} `json:"properties"`
	} `json:"features"`
}

type TimeZone struct {
	Name            string  `json:"name"`
	OffsetSTD       float64 `json:"offset_STD"`
	OffsetDST       float64 `json:"offset_DST"`
	AbbreviationSTD string  `json:"abbreviation_STD"`
	AbbreviationDST string  `json:"abbreviation_DST"`
}
