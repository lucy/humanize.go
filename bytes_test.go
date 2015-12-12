package humanize

import (
	"testing"
)

func TestBytes(t *testing.T) {
	testList{
		{"bytes(0)", Bytes(0), "0B"},
		{"bytes(1)", Bytes(1), "1B"},
		{"bytes(803)", Bytes(803), "803B"},
		{"bytes(1023)", Bytes(1023), "1.0KiB"},
		{"bytes(1024)", Bytes(1024), "1.0KiB"},
		{"bytes(1MiB - 1)", Bytes(MiB - 1), "1.0MiB"},
		{"bytes(1MiB)", Bytes(1024 * 1024), "1.0MiB"},
		{"bytes(1GiB - 1K)", Bytes(GiB - KiB), "1.0GiB"},
		{"bytes(1GiB)", Bytes(GiB), "1.0GiB"},
		{"bytes(1TiB - 1M)", Bytes(TiB - MiB), "1.0TiB"},
		{"bytes(1TiB)", Bytes(TiB), "1.0TiB"},
		{"bytes(1PiB - 1T)", Bytes(PiB - TiB), "1.0PiB"},
		{"bytes(1PiB)", Bytes(PiB), "1.0PiB"},
		{"bytes(1PiB - 1T)", Bytes(EiB - PiB), "1.0EiB"},
		{"bytes(1EiB)", Bytes(EiB), "1.0EiB"},
		//{"bytes(1EB - 1P)", Bytes(1023 * EiB), "1023EB"},
		{"bytes(5.5GiB)", Bytes(5.5 * GiB), "5.5GiB"},
	}.validate(t)
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(16.5 * GiB)
	}
}
