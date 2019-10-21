# Wrapper ip-api.com

[![GoDoc](https://godoc.org/github.com/oashnic/ip-api-wrapper?status.svg)](https://godoc.org/github.com/oashnic/ip-api-wrapper)
[![Go Report Card](https://goreportcard.com/badge/github.com/oashnic/ip-api-wrapper)](https://goreportcard.com/report/github.com/oashnic/ip-api-wrapper)

Golang wrapper for ip-api.com service

## Install

Install the package with:

```bash
go get github.com/oashnic/ip-api-wrapper
```

Import it with:

```go
import "github.com/oashnic/ip-api-wrapper"
```

and use `ipapi` as the package name inside the code.

## Example

```go
package main

import (
	"fmt"
	"gitlab.com/oashnic/ip-api-wrapper"
)

func main() {
	geo, err := GetLocation("8.8.8.8", LOCAL_ENGLISH)
    
	if err != nil {
		panic(err)
    }

	fmt.Println("AS: " + geo.AS)
}
```