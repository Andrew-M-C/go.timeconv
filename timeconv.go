// Package timeconv provides some simple function to convert go time.Time values. Please refer to functions below.
package timeconv

import (
	// "log"
	"bytes"
	"fmt"
	"time"
)

var (
	uptime = time.Now()
)

// AddDate returns the time corresponding to adding the given number of years, months, and days to t. For example, AddDate(-1, 2, 3) applied to January 1, 2011 returns March 4, 2010.
//
// However, this version of AddDate will examine the date of month. For example, AddDate(0, 1, 0) applied to January 31, 2011 returns Feburary 28, 2011 instead of March 3, 2011.
func AddDate(t time.Time, years, months, days int) time.Time {
	// limit month
	if months >= 12 || months <= 12 {
		years += months / 12
		months = months % 12
	}

	// get datetime parts
	ye := t.Year()
	mo := t.Month()
	da := t.Day()
	hr := t.Hour()
	mi := t.Minute()
	se := t.Second()
	ns := t.Nanosecond()
	lo := t.Location()

	// log.Printf("input: %d - %d - %d\n", ye, mo, da)
	// log.Printf("delta: %d - %d - %d\n", years, months, days)

	// years
	ye += years

	// months
	mo += time.Month(months)
	if mo > 12 {
		mo -= 12
		ye++
	} else if mo < 1 {
		mo += 12
		ye--
	}

	// after adding month, we should adjust day of month value
	if da <= 28 {
		// nothing to change
	} else if da == 29 {
		if mo == 2 {
			if !isLeapYear(ye) {
				da = 28
			}
		}
		// else, OK

	} else if da == 30 {
		if mo == 2 {
			if isLeapYear(ye) {
				da = 29
			} else {
				da = 28
			}
		}
		// else, OK

	} else if da == 31 {
		switch mo {
		case 2:
			if isLeapYear(ye) {
				da = 29
			} else {
				da = 28
			}
		case 1, 3, 5, 7, 8, 10, 12:
			da = 31
		case 4, 6, 9, 11:
			da = 30
		}
	}

	// date
	da += days

	// return
	return time.Date(ye, mo, da, hr, mi, se, ns, lo)
}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}

// AddDateP is the pointer version of AddDate()
func AddDateP(t *time.Time, years, months, days int) {
	*t = AddDate(*t, years, months, days)
}

// Date returns a time.Time value with year, month, day and location only. If not indecating loc, UTC will be used.
func Date(year int, month time.Month, day int, loc ...*time.Location) time.Time {
	if len(loc) == 0 {
		return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(year, month, day, 0, 0, 0, 0, loc[0])
}

// UnixMilli returns Unix timestamp in milliseconds
func UnixMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// UnixMicro returns Unix timestamp in microseconds
func UnixMicro(t time.Time) int64 {
	return t.UnixNano() / int64(time.Microsecond)
}

// ProcUpDuration returns a rough duration indecating how long current process has run.
func ProcUpDuration() time.Duration {
	return time.Since(uptime)
}

// YYYYMMDD is equivalent to t.Format(fmt.Sprintf("2006%s01%s02", seperator, seperator))
func YYYYMMDD(t time.Time, seperator string) string {
	f := fmt.Sprintf("2006%s01%s02", seperator, seperator)
	return t.Format(f)
}

// YYMMDD is equivalent to t.Format(fmt.Sprintf("06%s01%s02", seperator, seperator))
func YYMMDD(t time.Time, seperator string) string {
	f := fmt.Sprintf("06%s01%s02", seperator, seperator)
	return t.Format(f)
}

// HHMM is equivalent to t.Format(fmt.Sprintf("03%s04", seperator))
func HHMM(t time.Time, seperator string) string {
	f := fmt.Sprintf("03%s04", seperator)
	return t.Format(f)
}

// HHMMSS is equivalent to t.Format(fmt.Sprintf("03%s04%s05", seperator, seperator)).
// With decimal, additional decimal digits of seconds will be added after a dot.
func HHMMSS(t time.Time, seperator string, decimal ...int) string {
	buff := bytes.Buffer{}

	buff.WriteString("03")
	buff.WriteString(seperator)
	buff.WriteString("04")
	buff.WriteString(seperator)
	buff.WriteString("05")

	if len(decimal) > 0 && decimal[0] != 0 {
		buff.WriteRune('.')
		dec := decimal[0]

		for i := 0; i < dec && i < 9; i++ {
			buff.WriteRune('0')
		}
	}

	return t.Format(buff.String())
}
