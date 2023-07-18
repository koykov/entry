package entry

import "testing"

func TestEntry24(t *testing.T) {
	t.Run("encode", func(t *testing.T) {
		lo, hi := uint32(12), uint8(16)
		x := NewEntry24(lo, hi)
		if x != 3088 {
			t.Errorf("Encode fail: need %d, got %d", 3088, x)
		}
	})
	t.Run("decode", func(t *testing.T) {
		x := Entry24(3088)
		if lo, hi := x.Decode(); lo != 12 || hi != 16 {
			t.Errorf("Decode fail: need %d/%d, got %d/%d", 12, 16, lo, hi)
		}
	})
}

func BenchmarkEntry24(b *testing.B) {
	b.Run("encode", func(b *testing.B) {
		var x Entry24
		for i := 0; i < b.N; i++ {
			x.Encode(12, 16)
		}
	})
	b.Run("decode", func(b *testing.B) {
		var x Entry24
		for i := 0; i < b.N; i++ {
			x.Decode()
		}
	})
}
