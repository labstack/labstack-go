package labstack

import "os"

var (
	client = NewClient(os.Getenv("KEY"))
	cs = client.Currency()
	ds = client.Domain()
	es = client.Email()
	is = client.IP()
	ws = client.Webpage()
)
