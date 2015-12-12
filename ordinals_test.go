package humanize

import (
	"testing"
)

func TestOrdinals(t *testing.T) {
	testList{
		{"0", Ordinal(0), "th"},
		{"1", Ordinal(1), "st"},
		{"2", Ordinal(2), "nd"},
		{"3", Ordinal(3), "rd"},
		{"4", Ordinal(4), "th"},
		{"10", Ordinal(10), "th"},
		{"11", Ordinal(11), "th"},
		{"12", Ordinal(12), "th"},
		{"13", Ordinal(13), "th"},
		{"101", Ordinal(101), "st"},
		{"102", Ordinal(102), "nd"},
		{"103", Ordinal(103), "rd"},
	}.validate(t)
}
