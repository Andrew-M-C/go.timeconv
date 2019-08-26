package timeconv

import (
	"testing"
	"time"
)

func TestAddDate(test *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	var t, expected time.Time

	t = time.Date(2019, 1, 31, 0, 0, 0, 0, loc)
	expected = time.Date(2019, 2, 28, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 1, 0)
	check(test, t, expected)

	t = time.Date(2019, 1, 29, 0, 0, 0, 0, loc)
	expected = time.Date(2019, 2, 28, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 1, 0)
	check(test, t, expected)

	t = time.Date(2000, 1, 30, 0, 0, 0, 0, loc)
	expected = time.Date(2000, 2, 29, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 1, 0)
	check(test, t, expected)

	t = time.Date(1900, 1, 29, 0, 0, 0, 0, loc)
	expected = time.Date(1900, 2, 28, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 1, 0)
	check(test, t, expected)

	t = time.Date(2020, 2, 1, 0, 0, 0, 0, loc)
	expected = time.Date(2020, 3, 3, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 0, 31)
	check(test, t, expected)

	t = time.Date(2020, 3, 1, 0, 0, 0, 0, loc)
	expected = time.Date(2020, 6, 1, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 0, 92)
	check(test, t, expected)

	t = time.Date(2020, 2, 29, 0, 0, 0, 0, loc)
	expected = time.Date(2021, 2, 28, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 12, 0)
	check(test, t, expected)

	t = time.Date(2019, 12, 31, 0, 0, 0, 0, loc)
	expected = time.Date(2020, 2, 29, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 2, 0)
	check(test, t, expected)

	t = time.Date(2019, 12, 31, 0, 0, 0, 0, loc)
	expected = time.Date(2016, 11, 30, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, -37, 0)
	check(test, t, expected)

	t = time.Date(2021, 1, 29, 0, 0, 0, 0, loc)
	expected = time.Date(2020, 2, 29, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, -11, 0)
	check(test, t, expected)

	t = time.Date(2019, 12, 30, 0, 0, 0, 0, loc)
	expected = time.Date(2019, 4, 30, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, -8, 0)
	check(test, t, expected)

	t = time.Date(2019, 12, 30, 0, 0, 0, 0, loc)
	expected = time.Date(2019, 2, 28, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, -10, 0)
	check(test, t, expected)

	t = time.Date(2019, 12, 29, 0, 0, 0, 0, loc)
	expected = time.Date(2019, 4, 29, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, -8, 0)
	check(test, t, expected)
}

func check(test *testing.T, t, expected time.Time) {
	if false == t.Equal(expected) {
		test.Errorf("Expected time %v, but got %v, error", expected, t)
	}
}
