package entry

import "testing"

func TestEntry64(t *testing.T) {
	t.Run("encode", func(t *testing.T) {
		lo, hi := uint32(23523), uint32(645623)
		x := NewEntry64(lo, hi)
		if x != 101030516349431 {
			t.Errorf("Encode fail: need %d, got %d", 101030516349431, x)
		}
	})
	t.Run("decode", func(t *testing.T) {
		x := Entry64(101030516349431)
		if lo, hi := x.Decode(); lo != 23523 || hi != 645623 {
			t.Errorf("Decode fail: need %d/%d, got %d/%d", 23523, 645623, lo, hi)
		}
	})
	t.Run("lo_hi", func(t *testing.T) {
		x := Entry64(101030516349431)
		if lo, hi := x.Lo(), x.Hi(); lo != 23523 || hi != 645623 {
			t.Errorf("Decode fail: need %d/%d, got %d/%d", 23523, 645623, lo, hi)
		}
	})
}

func BenchmarkEntry64(b *testing.B) {
	b.Run("encode", func(b *testing.B) {
		var x Entry64
		for i := 0; i < b.N; i++ {
			x.Encode(12, 16)
		}
	})
	b.Run("decode", func(b *testing.B) {
		var x Entry64
		for i := 0; i < b.N; i++ {
			x.Decode()
		}
	})
}
