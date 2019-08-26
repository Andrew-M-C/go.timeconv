package timeconv

import (
	// "log"
	"time"
)

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
		ye += 1
	} else if mo < 1 {
		mo += 12
		ye -= 1
	}

	// after adding month, we should adjust day of month value
	if da <= 28 {
		// nothing to change
	} else if da == 29 {
		if mo == 2 {
			if false == isLeapYear(ye) {
				da = 28
			}
		} else {
			// OK
		}
	} else if da == 30 {
		if mo == 2 {
			if isLeapYear(ye) {
				da = 29
			} else {
				da = 28
			}
		} else {
			// OK
		}
	} else if da == 31 {
		switch mo {
		case 2:
			if isLeapYear(ye) {
				da = 29
			} else {
				da = 28
			}
		case 1, 3, 5, 7, 8, 10, 12:
			// OK
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
	if 0 == year%4 {
		if 0 == year%100 {
			return 0 == year%400
		} else {
			return true
		}
	} else {
		return false
	}
}
