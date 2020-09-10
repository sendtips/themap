# TheMAP

[![GitHub Actions](https://github.com/sendtips/themap/workflows/Go/badge.svg)](https://github.com/sendtips/themap/actions?workflow=Go)
[![GoDoc](https://godoc.org/github.com/sendtips/themap?status.svg)](https://godoc.org/github.com/sendtips/themap)
[![codecov](https://codecov.io/gh/sendtips/themap/branch/master/graph/badge.svg)](https://codecov.io/gh/sendtips/themap)
[![Go Report Card](https://goreportcard.com/badge/github.com/sendtips/themap)](https://goreportcard.com/report/github.com/sendtips/themap)
[![Sourcegraph](https://sourcegraph.com/github.com/sendtips/themap/-/badge.svg)](https://sourcegraph.com/github.com/sendtips/themap?badge)

A Go library to work with [TheMAP](https://doc.mapcard.pro/).

## Install
Install by import `github.com/sendtips/themap` or via `go get github.com/sendtips/themap`

## Tests
Run tests using `THEMAPTERMID=TestTerminal THEMAPTERMPW=123 THEMAPMERCHID=TestMerchant THEMAPMERCHPW=123 THEMAPAPIHOST=https://api-stage.mapcard.pro go test -v .`.

*Note:* Use your credentials. The provided above will not work.
There also `THEMAPSIGNKEY` variable exists, but currently, no one test uses it.

## Example
To obtain a payment session you need to call `Init()` method.

```go
package main

import (
	"fmt"
	"context"
	"github.com/sendtips/themap"
)

func main() {
	pay := themap.New("TestTerminal", "TestOrder123")
	pay.SetTerm("123") // Sets Terminal password

	err := pay.Init(context.TODO(), 300) // Create a session for 3.00RUB
	if err != nil {
		fmt.Printf("Error occurred: %v", err)
	}

	fmt.Printf("%s", pay.SessionGUID) // TheMAP payment session identifier
}
```
