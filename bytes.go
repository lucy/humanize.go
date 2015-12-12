package humanize

import "fmt"

// IEC Sizes
const (
	KiB = 1024
	MiB = KiB * 1024
	GiB = MiB * 1024
	TiB = GiB * 1024
	PiB = TiB * 1024
	EiB = PiB * 1024
)

// BytesM produces a human readable representation of an IEC size with custom
// suffixes.
func BytesM(s uint64, suf []string) string {
	if s < 999 {
		return fmt.Sprintf("%d%s", s, suf[0])
	}

	m := 0
	div := float64(1)
	u := float64(s)
	for u > 999 {
		div *= 1024
		u /= 1024
		m++
	}
	if u < 10 {
		return fmt.Sprintf("%0.1f%s", float64(s)/float64(div), suf[m])
	}
	return fmt.Sprintf("%.0f%s", float64(s)/float64(div), suf[m])
}

// Bytes produces a human readable representation of an IEC size.
func Bytes(s uint64) string {
	var sizes = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
	return BytesM(s, sizes)
}
