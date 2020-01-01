package timeconv_test

import (
	"fmt"

	timeconv "github.com/Andrew-M-C/go.timeconv"
)

func Example() {
	p := fmt.Println
	t := timeconv.Date(2020, 1, 31)

	p(t)
	p(t.AddDate(0, 1, 0))
	p(timeconv.AddDate(t, 0, 1, 0))

	// Output:
	// 2020-01-31 00:00:00 +0000 UTC
	// 2020-03-02 00:00:00 +0000 UTC
	// 2020-02-29 00:00:00 +0000 UTC
}
