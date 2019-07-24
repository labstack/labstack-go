package domain

type (
	Record struct {
		Domain   string   `json:"domain"`
		Type     string   `json:"type"`
		Server   string   `json:"server"`
		A        string   `json:"a"`
		AAAA     string   `json:"aaaa"`
		CNAME    string   `json:"cname"`
		MX       string   `json:"mx"`
		NS       string   `json:"ns"`
		PTR      string   `json:"ptr"`
		Serial   uint32   `json:"serial"`
		Refresh  uint32   `json:"refresh"`
		Retry    uint32   `json:"retry"`
		Expire   uint32   `json:"expire"`
		Priority uint32   `json:"priority"`
		Weight   uint32   `json:"weight"`
		Port     uint32   `json:"port"`
		Target   string   `json:"target"`
		TXT      []string `json:"txt"`
		TTL      uint32   `json:"ttl"`
		Class    string   `json:"class"`
		SPF      []string `json:"spf"`
	}

	Result struct {
		Domain string `json:"domain"`
		Zone   string `json:"zone"`
	}

	Registrar struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Url         string `json:"url"`
		WhoisServer string `json:"whois_server"`
	}

	Registrant struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Organization string `json:"organization"`
		Street       string `json:"street"`
		City         string `json:"city"`
		State        string `json:"state"`
		Zip          string `json:"zip"`
		Country      string `json:"country"`
		Phone        string `json:"phone"`
		Fax          string `json:"fax"`
		Email        string `json:"email"`
	}

	DNSRequest struct {
		Type   string
		Domain string
	}

	DNSResponse struct {
		Records []*Record
	}

	SearchRequest struct {
		Domain string
	}

	SearchResponse struct {
		Results []*Result
	}

	StatusRequest struct {
		Domain string
	}

	StatusResponse struct {
		Domain string   `json:"domain"`
		Zone   string   `json:"zone"`
		Result string   `json:"result"`
		Flags  []string `json:"flags"`
	}

	WhoisRequest struct {
		Domain string
	}

	WhoisResponse struct {
		Domain      string      `json:"domain"`
		Id          string      `json:"id"`
		Status      string      `json:"status"`
		CreatedDate string      `json:"created_date"`
		UpdatedDate string      `json:"updated_date"`
		ExpiryDate  string      `json:"expiry_date"`
		NameServers []string    `json:"name_servers"`
		Dnssec      string      `json:"dnssec"`
		Registrar   *Registrar  `json:"registrar"`
		Registrant  *Registrant `json:"registrant"`
		Admin       *Registrant `json:"admin"`
		Technical   *Registrant `json:"technical"`
		Billing     *Registrant `json:"billing"`
		Raw         string      `json:"raw"`
	}
)
