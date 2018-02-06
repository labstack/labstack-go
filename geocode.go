package labstack

type (
	GeocodeAddressRequest struct {
		Location  string  `json:"location"`
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
		OSMTag    string  `json:"osm_tag"`
		Language  string  `json:"language"`
		Limit     int     `json:"limit"`
	}

	GeocodeIPRequest struct {
		IP string `json:"ip"`
	}

	GeocodeReverseRequest struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
		Language  string  `json:"language"`
		Limit     int     `json:"limit"`
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
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/geocode/address")
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
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/geocode/ip")
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
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/geocode/reverse")
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
