package entry

import "testing"

func TestEntry16(t *testing.T) {
	t.Run("encode", func(t *testing.T) {
		var (
			x      Entry16
			lo, hi uint8
		)
		lo, hi = 12, 16
		x.Encode(lo, hi)
		if x != 3088 {
			t.Errorf("Encode fail: need %d, got %d", 3088, x)
		}
	})
	t.Run("decode", func(t *testing.T) {
		x := Entry16(3088)
		if lo, hi := x.Decode(); lo != 12 || hi != 16 {
			t.Errorf("Decode fail: need %d/%d, got %d/%d", 12, 16, lo, hi)
		}
	})
}

func BenchmarkEntry16(b *testing.B) {
	b.Run("encode", func(b *testing.B) {
		var x Entry16
		for i := 0; i < b.N; i++ {
			x.Encode(12, 16)
		}
	})
	b.Run("decode", func(b *testing.B) {
		var x Entry16
		for i := 0; i < b.N; i++ {
			x.Decode()
		}
	})
}
