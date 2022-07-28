package entry

// Entry16 compacts uint8 lo/hi values.
type Entry16 uint16

// Encode combines lo/hi values into one uint16 int.
func (e *Entry16) Encode(lo, hi uint8) {
	*e = Entry16(lo)<<8 | Entry16(hi)
}

// Decode splits e to lo/hi uint8 values.
func (e Entry16) Decode() (lo, hi uint8) {
	lo = uint8(e >> 8)
	hi = uint8(e)
	return
}

func (e *Entry16) Reset() {
	*e = 0
}
