package ip

import (
	"time"
)

type (
	Organization struct {
		Name string `json:"name"`
	}

	Flag struct {
		SVG   string `json:"svg"`
		PNG   string `json:"png"`
		Emoji string `json:"emoji"`
	}

	TimeZone struct {
		ID           string    `json:"id"`
		Name         string    `json:"name"`
		Abbreviation string    `json:"abbreviation"`
		Offset       int       `json:"offset"`
		Time         time.Time `json:"time"`
	}

	AS struct {
		Number       int    `json:"number"`
		Name         string `json:"name"`
		Organization string `json:"organization"`
	}

	LookupRequest struct {
		IP string
	}

	LookupResponse struct {
		IP           string        `json:"ip"`
		Hostname     string        `json:"hostname"`
		Version      string        `json:"version"`
		City         string        `json:"city"`
		Region       string        `json:"region"`
		RegionCode   string        `json:"region_code"`
		Postal       string        `json:"postal"`
		Country      string        `json:"country"`
		CountryCode  string        `json:"country_code"`
		Latitude     float64       `json:"latitude"`
		Longitude    float64       `json:"longitude"`
		Organization *Organization `json:"organization"`
		Flag         *Flag         `json:"flag"`
		Currencies   []string      `json:"currencies"`
		TimeZone     *TimeZone     `json:"time_zone"`
		Language     []string      `json:"languages"`
		AS           *AS           `json:"as"`
		Flags        []string      `json:"flags"`
	}
)

