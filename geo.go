package models

type GeoLocation struct {
	Features []struct {
		Properties struct {
			TimeZone TimeZone `json:"timezone"`
		} `json:"properties"`
	} `json:"features"`
}

type TimeZone struct {
	Name            string `json:"name"`
	OffsetSTD       int    `json:"offset_STD_seconds"`
	OffsetDST       int    `json:"offset_DST_seconds"`
	AbbreviationSTD string `json:"abbreviation_STD"`
	AbbreviationDST string `json:"abbreviation_DST"`
}
