package entry

// Entry32 compacts uint16 lo/hi values.
type Entry32 uint32

func NewEntry32(lo, hi uint16) Entry32 {
	var e Entry32
	e.Encode(lo, hi)
	return e
}

// Encode combines lo/hi values into one uint32 int.
func (e *Entry32) Encode(lo, hi uint16) {
	*e = Entry32(lo)<<16 | Entry32(hi)
}

// Decode splits e to lo/hi uint16 values.
func (e Entry32) Decode() (lo, hi uint16) {
	lo = uint16(e >> 16)
	hi = uint16(e)
	return
}

func (e *Entry32) Reset() {
	*e = 0
}
