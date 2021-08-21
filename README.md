# MobileNig Go

[![Build](https://github.com/NdoleStudio/mobilenig-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/mobilenig-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/mobilenig-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/mobilenig-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/mobilenig-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/mobilenig-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/mobilenig-go)](https://goreportcard.com/report/github.com/NdoleStudio/mobilenig-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/mobilenig-go)](https://github.com/NdoleStudio/mobilenig-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/mobilenig-go?color=brightgreen)](https://github.com/NdoleStudio/mobilenig-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/mobilenig-go)](https://pkg.go.dev/github.com/NdoleStudio/mobilenig-go)

This package provides a `go` client for interacting with the [MobileNig API](https://mobilenig.com/API/docs/index)

## Installation

`mobilenig-go` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/mobilenig-go
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/mobilenig-go"
```

## Implemented

- [Bills](#bills)
  - DStv
    - `GET /bills/user_check` - Validate a DStv user
    - `GET /bills/dstv` - Pay a dstv subscription
    - `GET /bills/query` - Fetch a dstv transaction

## Usage

### Initializing the Client

An instance of the `mobilenig` client can be created using `New()`.  The `http.Client` supplied will be used to make requests to the API.

```go
package main

import (
	"github.com/NdoleStudio/mobilenig-go"
)

func main()  {
	client := mobilenig.New(
		mobilenig.WithUsername("" /* MobileNig Username */),
		mobilenig.WithAPIKey("" /* MobileNig API Key */),
		mobilenig.WithEnvironment(mobilenig.TestEnvironment),
	)
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
payload, response, err := mobilenigClient.Token(context.Background())
if err != nil {
  //handle error
}
```

### Bills

This handles all API requests whose URL begins with `/bills/`

#### DStv

##### Validate DStv User

`GET /bills/user_check`: Validate a DStv user

```go
user, _, err := mobilenigClient.Bills.CheckDStvUser(context.Background(), "4131953321")

if err != nil {
    log.Fatal(err)
}

log.Println(user.Details.LastName) // e.g INI OBONG BASSEY
```

##### Pay a DStv subscription

`GET /bills/dstv` - Pay a DStv subscription

```go
user, _, err := mobilenigClient.Bills.PayDStv(context.Background(), "4131953321")

if err != nil {
    log.Fatal(err)
}

log.Println(user.Details.LastName) // e.g INI OBONG BASSEY
```

##### Fetch a DStv transaction

`GET /bills/query` - Fetch a DStv transaction

```go
user, _, err := mobilenigClient.Bills.PayDStv(context.Background(), "4131953321")

if err != nil {
    log.Fatal(err)
}

log.Println(user.Details.LastName) // e.g INI OBONG BASSEY
```

## Testing

You can run the unit tests for this SDK from the root directory using the command below:
```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
