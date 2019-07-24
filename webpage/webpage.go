package webpage

type (
	ImageRequest struct {
		URL      string
		Language string
		TTL      int
		FullPage bool
		Retina   bool
		Width    int
		Height   int
		Delay    int
	}

	ImageResponse struct {
		Image       string `json:"image"`
		Cached      bool   `json:"cached"`
		Tool        int    `json:"tool"`
		GeneratedAt string `json:"generated_at"`
	}

	PDFRequest struct {
		URL         string
		Language    string
		TTL         int
		Size        string
		Width       int
		Height      int
		Orientation string
		Delay       int
	}

	PDFResponse struct {
		PDF         string `json:"pdf"`
		Cached      bool   `json:"cached"`
		Took        int    `json:"took"`
		GeneratedAt string `json:"generated_at"`
	}
)
