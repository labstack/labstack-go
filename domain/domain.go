package domain

type (
	Record struct {
		Domain   string `json:"domain"`
		Type     string `json:"type"`
		Server   string `json:"server"`
		A        string
		AAAA     string
		CNAME    string
		MX       string
		NS       string
		PTR      string
		Serial   int    `json:"serial"`
		Refresh  int    `json:"refresh"`
		Retry    int    `json:"retry"`
		Expire   int    `json:"expire"`
		Priority int    `json:"priority"`
		Weight   int    `json:"weight"`
		Port     int    `json:"port"`
		Target   string `json:"target"`
		TXT      []string
		TTL      int    `json:"ttl"`
		Class    string `json:"class"`
		SPF      []string
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
		Q string
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
