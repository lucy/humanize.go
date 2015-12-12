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

// TimeRelativeTo formats a time into a relative string.
// TimeRelativeTo(now, now.Add(-3*Week)) -> "3 weeks ago"
func TimeRelativeTo(now, then time.Time) string {
	var te string
	var diff time.Duration
	if then.After(now) {
		te = "from now"
		diff = then.Sub(now)
	} else {
		te = "ago"
		diff = now.Sub(then)
	}

	switch {
	case diff/time.Second <= 0:
		return "now"

	case diff <= 2*time.Second:
		return fmt.Sprintf("1 second %s", te)
	case diff < 1*time.Minute:
		return fmt.Sprintf("%d seconds %s", diff/time.Second, te)

	case diff < 2*time.Minute:
		return fmt.Sprintf("1 minute %s", te)
	case diff < 1*time.Hour:
		return fmt.Sprintf("%d minutes %s", diff/time.Minute, te)

	case diff < 2*time.Hour:
		return fmt.Sprintf("1 hour %s", te)
	case diff < 1*Day:
		return fmt.Sprintf("%d hours %s", diff/time.Hour, te)

	case diff < 2*Day:
		return fmt.Sprintf("1 day %s", te)
	case diff < 1*Week:
		return fmt.Sprintf("%d days %s", diff/Day, te)

	case diff < 2*Week:
		return fmt.Sprintf("1 week %s", te)
	case diff < 1*Month:
		return fmt.Sprintf("%d weeks %s", diff/Week, te)

	case diff < 2*Month:
		return fmt.Sprintf("1 month %s", te)
	case diff < 1*Year:
		return fmt.Sprintf("%d months %s", diff/Month, te)

	case diff < 18*Month:
		return fmt.Sprintf("1 year %s", te)
	}

	return then.Format("2006-01-02")
}

// Time is TimeRelativeTo(time.Now(), then)
func Time(then time.Time) string {
	return TimeRelativeTo(time.Now(), then)
}
