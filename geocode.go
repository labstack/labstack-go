package labstack

import "strconv"

type (
	Geocode struct {
		*Client
	}

	GeocodeAddressOptions struct {
		Longitude float64
		Latitude  float64
		OSMTag    string
		Language  string
		Limit     int
	}

	GeocodeReverseOptions struct {
		Limit int
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

func (g *Geocode) Address(location string, options GeocodeAddressOptions) (*GeocodeResponse, *APIError) {
	res := new(GeocodeResponse)
	err := new(APIError)
	r, e := g.resty.R().
		SetQueryParams(map[string]string{
			"location":  location,
			"longitude": strconv.FormatFloat(options.Longitude, 'f', -1, 64),
			"latitude":  strconv.FormatFloat(options.Latitude, 'f', -1, 64),
			"osm_tag":   options.OSMTag,
			"language":  options.Language,
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
	res := new(GeocodeResponse)
	err := new(APIError)
	r, e := g.resty.R().
		SetQueryParams(map[string]string{
			"ip": ip,
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

func (g *Geocode) Reverse(longitude, latitude float64, options GeocodeReverseOptions) (*GeocodeResponse, *APIError) {
	res := new(GeocodeResponse)
	err := new(APIError)
	r, e := g.resty.R().
		SetQueryParams(map[string]string{
			"longitude": strconv.FormatFloat(longitude, 'f', -1, 64),
			"latitude":  strconv.FormatFloat(latitude, 'f', -1, 64),
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
