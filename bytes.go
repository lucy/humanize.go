package humanize

import "strconv"

// IEC Sizes
const (
	KiB = 1024
	MiB = KiB * 1024
	GiB = MiB * 1024
	TiB = GiB * 1024
	PiB = TiB * 1024
	EiB = PiB * 1024
)

func itoa3(b []byte, x int) []byte {
	if x/100 != 0 {
		b = append(b, '0'+byte(x/100))
	}
	if x/100 != 0 || x/10%10 != 0 {
		b = append(b, '0'+byte(x/10%10))
	}
	b = append(b, '0'+byte(x%10))
	return b
}

// BytesM produces a human readable representation of an IEC size with custom
// suffixes.
func BytesM(s uint64, suf []string) string {
	if s < 999 {
		return string(append(strconv.AppendUint(nil, s, 10), suf[0]...))
	}

	m := 0
	div := float64(1)
	u := float64(s)
	for u > 999 {
		div *= 1024
		u /= 1024
		m++
	}
	v := float64(s) / div
	b := make([]byte, 0, 3+len(suf[m]))
	if v < 10 {
		d := byte(v*10 + 0.5)
		b = append(b, '0'+d/10, '.', '0'+d%10)
	} else {
		b = itoa3(b, int(v+0.5))
	}
	return string(append(b, suf[m]...))
}

// Bytes produces a human readable representation of an IEC size.
func Bytes(s uint64) string {
	var sizes = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
	return BytesM(s, sizes)
}
