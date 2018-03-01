package labstack

import "strconv"

type (
	GeocodeAddressRequest struct {
		Location  string
		Longitude float64
		Latitude  float64
		OSMTag    string
		Language  string
		Limit     int
	}

	GeocodeIPRequest struct {
		IP string
	}

	GeocodeReverseRequest struct {
		Longitude float64
		Latitude  float64
		Language  string
		Limit     int
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

func (c *Client) GeocodeAddress(req *GeocodeAddressRequest) (*GeocodeResponse, *APIError) {
	res := new(GeocodeResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetQueryParams(map[string]string{
			"location":  req.Location,
			"longitude": strconv.FormatFloat(req.Longitude, 'f', -1, 64),
			"latitude":  strconv.FormatFloat(req.Latitude, 'f', -1, 64),
			"osm_tag":   req.OSMTag,
			"language":  req.Language,
			"limit":     strconv.Itoa(req.Limit),
		}).
		SetResult(res).
		SetError(err).
		Get("/geocode/address")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if c.error(r) {
		return nil, err
	}
	return res, nil
}

func (c *Client) GeocodeIP(req *GeocodeIPRequest) (*GeocodeResponse, *APIError) {
	res := new(GeocodeResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetQueryParams(map[string]string{
			"ip": req.IP,
		}).
		SetResult(res).
		SetError(err).
		Get("/geocode/ip")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if c.error(r) {
		return nil, err
	}
	return res, nil
}

func (c *Client) GeocodeReverse(req *GeocodeReverseRequest) (*GeocodeResponse, *APIError) {
	res := new(GeocodeResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetQueryParams(map[string]string{
			"longitude": strconv.FormatFloat(req.Longitude, 'f', -1, 64),
			"latitude":  strconv.FormatFloat(req.Latitude, 'f', -1, 64),
			"language":  req.Language,
			"limit":     strconv.Itoa(req.Limit),
		}).
		SetResult(res).
		SetError(err).
		Get("/geocode/reverse")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if c.error(r) {
		return nil, err
	}
	return res, nil
}
