# timeconv

Package timeconv provides some simple function to convert go time.Time values.

[![GoDoc](https://godoc.org/github.com/Andrew-M-C/go.timeconv?status.svg)](https://godoc.org/github.com/Andrew-M-C/go.timeconv)
[![Build status](https://travis-ci.org/Andrew-M-C/go.timeconv.svg?branch=master)](https://travis-ci.org/Andrew-M-C/go.timeconv)
[![Coverage Status](https://coveralls.io/repos/github/Andrew-M-C/go.timeconv/badge.svg?branch=master)](https://coveralls.io/github/Andrew-M-C/go.timeconv?branch=master)
[![Latest](https://img.shields.io/badge/latest-v0.4.0-blue.svg)](https://github.com/Andrew-M-C/go.timeconv/tree/v0.4.0)

The function `AddDate` in  [package time](https://golang.org/pkg/time/) adds dates passed in format of years, months and days. In the biginning, I thought it did handle different date of month. However, it is NOT.

For example, in many case in real life, if a month is added to 2019-01-31, the result should be 2019-01-28, identifying last day of the month. But the result with `AddDate` in package `time` will be 2019-03-03 ([Playground](https://play.golang.org/p/3eWrvAVzHUm))

This simple package `timeconv` provides a simple alternate of `AddDate` function, it focus on month operation and corrects result day for return. Please refer to the demo below for usage:

```go
package main

import (
	"fmt"
	"time"
	"github.com/Andrew-M-C/go.timeconv"
)

func main() {
	t := time.Date(2019, 1, 31, 0, 0, 0, 0, time.UTC)
	nt := t.AddDate(0, 1, 0)	// Add one month
	fmt.Printf("%v\n", nt)		// 2019-03-03 00:00:00 +0000 UTC, not expected

	nt = timeconv.AddDate(t, 0, 1, 0)
	fmt.Printf("%v\n", nt)		// 2019-02-28 00:00:00 +0000 UTC
}
```

[Playground](https://play.golang.org/p/-2tnI8Ejxwh)

Additionally, there is a pointer version: `func AddDateP(t *time.Time, years, months, days int)`
