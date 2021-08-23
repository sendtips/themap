# TheMAP

[![GitHub Actions](https://github.com/sendtips/themap/workflows/Go/badge.svg)](https://github.com/sendtips/themap/actions?workflow=Go)
[![GoDoc](https://godoc.org/github.com/sendtips/themap?status.svg)](https://godoc.org/github.com/sendtips/themap)
[![codecov](https://codecov.io/gh/sendtips/themap/branch/master/graph/badge.svg)](https://codecov.io/gh/sendtips/themap)
[![Go Report Card](https://goreportcard.com/badge/github.com/sendtips/themap)](https://goreportcard.com/report/github.com/sendtips/themap)
[![Sourcegraph](https://sourcegraph.com/github.com/sendtips/themap/-/badge.svg)](https://sourcegraph.com/github.com/sendtips/themap?badge)
[![sendtips](https://img.shields.io/badge/🍩_Sendtips-@awsom82-black?labelColor=3298dc)](https://sendtips.ru/pay/E2ZfzjVE)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsendtips%2Fthemap.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsendtips%2Fthemap?ref=badge_shield)


A Go library to work with [TheMAP](https://doc.mapcard.pro/).

## Install
Install by import `github.com/sendtips/themap` or via `go get github.com/sendtips/themap`

The library itself can be compiled on go1.13,
but to run tests you need go1.14 due to `strconv.NumError` received `Unwrap()` method only in v1.14

## Tests
Run tests using `THEMAPTERMID=TestTerminal THEMAPTERMPW=123 THEMAPMERCHID=TestMerchant THEMAPMERCHPW=123 THEMAPAPIHOST=https://api-stage.mapcard.pro go test -v .`

*Note:* Use your credentials. The provided above will not work.
There also `THEMAPSIGNKEY` variable exists, but currently, no one test uses it properly.

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
	pay.SetTerm("123") // Set Terminal password

	err := pay.Init(context.TODO(), 300) // Create a session for 3.00RUB
	if err != nil {
		fmt.Printf("Error occurred: %v", err)
	}

	fmt.Printf("%s", pay.SessionGUID) // TheMAP payment session identifier
}
```


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsendtips%2Fthemap.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsendtips%2Fthemap?ref=badge_large)