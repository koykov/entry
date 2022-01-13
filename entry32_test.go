package entry

import "testing"

func TestEntry32(t *testing.T) {
	t.Run("encode", func(t *testing.T) {
		var (
			x      Entry32
			lo, hi uint16
		)
		lo, hi = 12, 16
		x.Encode(lo, hi)
		if x != 786448 {
			t.Errorf("Encode fail: need %d, got %d", 786448, x)
		}
	})
	t.Run("decode", func(t *testing.T) {
		x := Entry32(786448)
		if lo, hi := x.Decode(); lo != 12 || hi != 16 {
			t.Errorf("Decode fail: need %d/%d, got %d/%d", 12, 16, lo, hi)
		}
	})
}

func BenchmarkEntry32(b *testing.B) {
	b.Run("encode", func(b *testing.B) {
		var x Entry32
		for i := 0; i < b.N; i++ {
			x.Encode(12, 16)
		}
	})
	b.Run("decode", func(b *testing.B) {
		var x Entry32
		for i := 0; i < b.N; i++ {
			x.Decode()
		}
	})
}
