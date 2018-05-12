package labstack

import "strconv"

type (
	Geocode struct {
		*Client
	}

	GeocodeAddressOptions struct {
		Latitude  float64
		Longitude float64
		OSMTag    string
		Language  string
		Foramt    string
		Limit     int
	}

	GeocodeIPOptions struct {
		Foramt string
	}

	GeocodeReverseOptions struct {
		Foramt string
		Limit  int
	}

	GeocodeGeometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}

	GeocodeFeature struct {
		Type       string           `json:"type"`
		Geometry   *GeocodeGeometry `json:"geometry"`
		Properties Properties       `json:"properties"`
	}

	GeocodeResponse struct {
		Type     string            `json:"type"`
		Features []*GeocodeFeature `json:"features"`
	}
)

func (g *Geocode) Address(location string) (*GeocodeResponse, *APIError) {
	return g.AddressWithOptions(location, GeocodeAddressOptions{})
}

func (g *Geocode) AddressWithOptions(location string, options GeocodeAddressOptions) (*GeocodeResponse, *APIError) {
	res := new(GeocodeResponse)
	err := new(APIError)
	r, e := g.resty.R().
		SetQueryParams(map[string]string{
			"location":  location,
			"latitude":  strconv.FormatFloat(options.Latitude, 'f', -1, 64),
			"longitude": strconv.FormatFloat(options.Longitude, 'f', -1, 64),
			"osm_tag":   options.OSMTag,
			"language":  options.Language,
			"foramt":    options.Foramt,
			"limit":     strconv.Itoa(options.Limit),
		}).
		SetResult(res).
		SetError(err).
		Get("/geocode/address")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if g.error(r) {
		return nil, err
	}
	return res, nil
}

func (g *Geocode) IP(ip string) (*GeocodeResponse, *APIError) {
	return g.IPWithOptions(ip, GeocodeIPOptions{})
}

func (g *Geocode) IPWithOptions(ip string, options GeocodeIPOptions) (*GeocodeResponse, *APIError) {
	res := new(GeocodeResponse)
	err := new(APIError)
	r, e := g.resty.R().
		SetQueryParams(map[string]string{
			"ip":     ip,
			"foramt": options.Foramt,
		}).
		SetResult(res).
		SetError(err).
		Get("/geocode/ip")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if g.error(r) {
		return nil, err
	}
	return res, nil
}

func (g *Geocode) Reverse(latitude, longitude float64) (*GeocodeResponse, *APIError) {
	return g.ReverseWithOptions(latitude, longitude, GeocodeReverseOptions{})
}

func (g *Geocode) ReverseWithOptions(latitude, longitude float64, options GeocodeReverseOptions) (*GeocodeResponse, *APIError) {
	res := new(GeocodeResponse)
	err := new(APIError)
	r, e := g.resty.R().
		SetQueryParams(map[string]string{
			"latitude":  strconv.FormatFloat(latitude, 'f', -1, 64),
			"longitude": strconv.FormatFloat(longitude, 'f', -1, 64),
			"foramt":    options.Foramt,
			"limit":     strconv.Itoa(options.Limit),
		}).
		SetResult(res).
		SetError(err).
		Get("/geocode/reverse")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if g.error(r) {
		return nil, err
	}
	return res, nil
}
