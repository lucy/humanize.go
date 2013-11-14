package humanize

import (
	"fmt"
	"time"
)

// Time units
const (
	Day   = 24 * time.Hour
	Week  = 7 * Day
	Month = 30 * Day
	Year  = 12 * Month
)

// Time formats a time into a relative string.
// Time(someT) -> "3 weeks ago"
func Time(then time.Time) string {
	now := time.Now()

	lbl := "ago"
	diff := now.Sub(then)
	if then.After(now) {
		lbl = "from now"
		diff = then.Sub(now)
	}

	switch {

	case diff <= 0:
		return "now"
	case diff <= 2*time.Second:
		return fmt.Sprintf("1 second %s", lbl)
	case diff < 1*time.Minute:
		return fmt.Sprintf("%d seconds %s", diff, lbl)

	case diff < 2*time.Minute:
		return fmt.Sprintf("1 minute %s", lbl)
	case diff < 1*time.Hour:
		return fmt.Sprintf("%d minutes %s", diff/time.Minute, lbl)

	case diff < 2*time.Hour:
		return fmt.Sprintf("1 hour %s", lbl)
	case diff < 1*Day:
		return fmt.Sprintf("%d hours %s", diff/time.Hour, lbl)

	case diff < 2*Day:
		return fmt.Sprintf("1 day %s", lbl)
	case diff < 1*Week:
		return fmt.Sprintf("%d days %s", diff/Day, lbl)

	case diff < 2*Week:
		return fmt.Sprintf("1 week %s", lbl)
	case diff < 1*Month:
		return fmt.Sprintf("%d weeks %s", diff/Week, lbl)

	case diff < 2*Month:
		return fmt.Sprintf("1 month %s", lbl)
	case diff < 1*Year:
		return fmt.Sprintf("%d months %s", diff/Month, lbl)

	case diff < 18*Month:
		return fmt.Sprintf("1 year %s", lbl)
	}

	return then.Format("2006-01-02")
}
