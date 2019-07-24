package labstack

import "os"

var (
	client = New(os.Getenv("KEY"))
)
