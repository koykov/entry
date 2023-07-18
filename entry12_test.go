package entry

import "testing"

func TestEntry12(t *testing.T) {
	t.Run("encode", func(t *testing.T) {
		lo, hi := uint16(128), uint8(8)
		x := NewEntry12(lo, hi)
		if x != 2056 {
			t.Errorf("Encode fail: need %d, got %d", 2056, x)
		}
	})
	t.Run("decode", func(t *testing.T) {
		x := Entry12(2056)
		if lo, hi := x.Decode(); lo != 128 || hi != 8 {
			t.Errorf("Decode fail: need %d/%d, got %d/%d", 128, 8, lo, hi)
		}
	})
}

func BenchmarkEntry12(b *testing.B) {
	b.Run("encode", func(b *testing.B) {
		var x Entry12
		for i := 0; i < b.N; i++ {
			x.Encode(128, 8)
		}
	})
	b.Run("decode", func(b *testing.B) {
		var x Entry12
		for i := 0; i < b.N; i++ {
			x.Decode()
		}
	})
}
