package humanize

// Comma produces a string form of the given number in base 10 with
// commas after every three orders of magnitude.
//
// e.g. Comma(834142) -> 834,142
func Comma(x int64) string {
	const len = 26 // sign + 6 commas + 19 numbers
	var b [len]byte
	i := len
	n := x < 0
	if n {
		x = -x
	}
	j := 0
	for ; x != 0; x /= 10 {
		if j == 3 {
			i--
			j = 0
			b[i] = ','
		}
		i--
		j++
		b[i] = '0' + byte(x%10)
	}

	if len-i < 1 {
		i--
		b[i] = '0'
	}

	if n {
		i--
		b[i] = '-'
	}

	return string(b[i:len])
}
