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
	client := labstack.NewClient("<ACCOUNT_ID>", "<API_KEY>")
	res, err := client.GeocodeAddress(&labstack.GeocodeAddressRequest{
		Location:  "eiffel tower",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", res)
	}
}
```

From terminal run your app:

```sh
go run app.go
```

## [Docs](https://labstack.com/docs) | [Forum](https://forum.labstack.com)
