package entry

import (
	"testing"
)

func TestEntry48(t *testing.T) {
	t.Run("encode", func(t *testing.T) {
		var (
			x  Entry48
			lo uint64
			hi uint16
		)
		lo, hi = 256, 128
		x.Encode(lo, hi)
		if x != 16777344 {
			t.Errorf("Encode fail: need %d, got %d", 16777344, x)
		}
	})
	t.Run("decode", func(t *testing.T) {
		x := Entry48(16777344)
		if lo, hi := x.Decode(); lo != 256 || hi != 128 {
			t.Errorf("Decode fail: need %d/%d, got %d/%d", 256, 128, lo, hi)
		}
	})
}

func BenchmarkEntry48(b *testing.B) {
	b.Run("encode", func(b *testing.B) {
		var x Entry48
		for i := 0; i < b.N; i++ {
			x.Encode(12, 16)
		}
	})
	b.Run("decode", func(b *testing.B) {
		var x Entry48
		for i := 0; i < b.N; i++ {
			x.Decode()
		}
	})
}
