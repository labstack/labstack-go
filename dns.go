package labstack

type (
	DNSLookupRequest struct {
		Domain string `json:"domain"`
		Type   string `json:"type"`
	}

	DNSLookupResponse struct {
		Records []DNSRecord `json:"records"`
	}

	DNSRecord struct {
		Type string `json:"type"`
		Name string `json:"name"`
		// Values - start
		A        string   `json:"a"`
		AAAA     string   `json:"aaaa"`
		CNAME    string   `json:"cname"`
		MX       string   `json:"mx"`
		NS       string   `json:"ms"`
		PTR      string   `json:"ptr"`
		Serial   uint32   `json:"serial"`
		Refresh  uint32   `json:"refresh"`
		Retry    uint32   `json:"retry"`
		Expire   uint32   `json:"expire"`
		Priority uint16   `json:"priority"`
		Weight   uint16   `json:"weight"`
		Port     uint16   `json:"port"`
		Target   string   `json:"target"`
		TXT      []string `json:"txt"`
		// Values - end
		TTL   uint32 `json:"ttl"`
		Class string `json:"class"`
	}
)

func (c *Client) DNSLookup(req *DNSLookupRequest) (*DNSLookupResponse, *APIError) {
	res := new(DNSLookupResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/dns/lookup")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if success(r) {
		return res, nil
	}
	return nil, err
}
