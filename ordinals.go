package humanize

// Ordinal gives you the input number's ordinal string.
// Ordinal(3) -> "rd"
func Ordinal(x int) string {
	switch {
	case x%10 == 1 && x%100 != 11:
		return "st"
	case x%10 == 2 && x%100 != 12:
		return "nd"
	case x%10 == 3 && x%100 != 13:
		return "rd"
	}
	return "th"
}
