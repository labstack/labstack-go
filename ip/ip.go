package ip

type (
	Currency struct {
		Name   string `json:"name"`
		Code   string `json:"code"`
		Symbol string `json:"symbol"`
	}

	Organization struct {
		Name string `json:"name"`
	}

	Flag struct {
		Image        string `json:"image"`
		Emoji        string `json:"emoji"`
		EmojiUnicode string `json:"emoji_unicode"`
	}

	TimeZone struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Abbreviation string `json:"abbreviation"`
		Offset       int32  `json:"offset"`
		Time         string `json:"time"`
	}

	Language struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}

	AS struct {
		Number       int64  `json:"number"`
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
		Currency     *Currency     `json:"currency"`
		TimeZone     *TimeZone     `json:"time_zone"`
		Language     []*Language   `json:"languages"`
		AS           *AS           `json:"as"`
		Flags        []string      `json:"flags"`
	}
)
