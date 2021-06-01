package timeconv

import (
	"testing"
	"time"
)

func TestUpTime(tt *testing.T) {
	time.Sleep(time.Second)

	d := ProcUpDuration()
	if d > time.Second && d <= (time.Second+10*time.Millisecond) {
		tt.Logf("proc uptime: %v", d) // OK
	} else {
		tt.Errorf("expected up time about 1 second, but got %v", d)
		return
	}
}

func TestAddDate(tt *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	var t, expected time.Time

	t = time.Date(2019, 1, 31, 0, 0, 0, 0, loc)
	expected = time.Date(2019, 2, 28, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 1, 0)
	check(tt, t, expected)

	t = time.Date(2019, 1, 29, 0, 0, 0, 0, loc)
	expected = time.Date(2019, 2, 28, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 1, 0)
	check(tt, t, expected)

	t = time.Date(2000, 1, 30, 0, 0, 0, 0, loc)
	expected = time.Date(2000, 2, 29, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 1, 0)
	check(tt, t, expected)

	t = time.Date(1900, 1, 29, 0, 0, 0, 0, loc)
	expected = time.Date(1900, 2, 28, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 1, 0)
	check(tt, t, expected)

	t = time.Date(2020, 2, 1, 0, 0, 0, 0, loc)
	expected = time.Date(2020, 3, 3, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 0, 31)
	check(tt, t, expected)

	t = time.Date(2020, 3, 1, 0, 0, 0, 0, loc)
	expected = time.Date(2020, 6, 1, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 0, 92)
	check(tt, t, expected)

	t = time.Date(2020, 2, 29, 0, 0, 0, 0, loc)
	expected = time.Date(2021, 2, 28, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 12, 0)
	check(tt, t, expected)

	t = time.Date(2019, 12, 31, 0, 0, 0, 0, loc)
	expected = time.Date(2020, 2, 29, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, 2, 0)
	check(tt, t, expected)

	t = time.Date(2019, 12, 31, 0, 0, 0, 0, loc)
	expected = time.Date(2016, 11, 30, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, -37, 0)
	check(tt, t, expected)

	t = time.Date(2021, 1, 29, 0, 0, 0, 0, loc)
	expected = time.Date(2020, 2, 29, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, -11, 0)
	check(tt, t, expected)

	t = time.Date(2019, 12, 30, 0, 0, 0, 0, loc)
	expected = time.Date(2019, 4, 30, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, -8, 0)
	check(tt, t, expected)

	t = time.Date(2019, 12, 30, 0, 0, 0, 0, loc)
	expected = time.Date(2019, 2, 28, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, -10, 0)
	check(tt, t, expected)

	t = time.Date(2019, 12, 29, 0, 0, 0, 0, loc)
	expected = time.Date(2019, 4, 29, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, -8, 0)
	check(tt, t, expected)

	t = time.Date(2019, 12, 31, 0, 0, 0, 0, loc)
	expected = time.Date(2019, 10, 31, 0, 0, 0, 0, loc)
	t = AddDate(t, 0, -2, 0)
	check(tt, t, expected)

	t = time.Date(2019, 12, 31, 0, 0, 0, 0, loc)
	expected = time.Date(2019, 10, 31, 0, 0, 0, 0, loc)
	AddDateP(&t, 0, -2, 0)
	check(tt, t, expected)
}

func TestDate(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")

	tm := Date(2019, 12, 31)
	expected := time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC)
	check(t, tm, expected)

	tm = Date(2019, 12, 31, loc)
	expected = time.Date(2019, 12, 31, 0, 0, 0, 0, loc)
	check(t, tm, expected)
}

func check(tt *testing.T, t, expected time.Time) {
	if false == t.Equal(expected) {
		tt.Errorf("Expected time %v, but got %v, error", expected, t)
	}
}

func TestUnix(t *testing.T) {
	now := time.Now()
	nano := now.UnixNano()

	if nano/1000 != UnixMicro(now) {
		t.Errorf("%d != %d !!!", nano/1000, UnixMicro(now))
	}
	if nano/1000000 != UnixMilli(now) {
		t.Errorf("%d != %d !!!", nano/1000000, UnixMilli(now))
	}
}

func TestFormat(t *testing.T) {
	now := time.Now()

	test := func(s, format string) {
		expect := now.Format(format)
		if s != expect {
			t.Errorf("format '%s', expect '%s', but got '%s'", format, expect, s)
		}
	}

	test(YYYYMMDD(now, "-"), "2006-01-02")
	test(YYMMDD(now, ","), "06,01,02")
	test(HHMM(now, "."), "03.04")
	test(HHMMSS(now, ":"), "03:04:05")
	test(HHMMSS(now, ":", 3), "03:04:05.000")
	test(HHMMSS(now, ":", 9), "03:04:05.000000000")
	test(HHMMSS(now, ":", 11), "03:04:05.000000000")
}
