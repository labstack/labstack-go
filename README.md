<a href="https://labstack.com"><img height="80" src="https://cdn.labstack.com/images/labstack-logo.svg"></a>

## Go Client

## Installation

`go get github.com/labstack/labstack-go`

## Quick Start

[Sign up](https://labstack.com/signup) to get an API key

Create a file `app.go` with the following content:

```go
package main

import (
	"fmt"

	"github.com/labstack/labstack-go"
)

func main() {
	client := labstack.NewClient("<API_KEY>")
	res, err := client.BarcodeGenerate(&labstack.BarcodeGenerateRequest{
		Format:  "qr_code",
		Content: "https://labstack.com",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		client.Download(res.ID, "/tmp/"+res.Name)
	}
}
```

From terminal run your app:

```sh
go run app.go
```

## [API](https://labstack.com/api) | [Forum](https://forum.labstack.com)
