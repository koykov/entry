package entry

// Entry12 compacts uint12 lo and uint4 hi values.
type Entry12 uint16

// Encode combines uint12 lo and uint4 hi values into one uint16 int.
func (e *Entry12) Encode(lo uint16, hi uint8) {
	*e = Entry12(lo)<<4 | Entry12(hi)
}

// Decode splits e to lo/hi uint8 values.
func (e Entry12) Decode() (lo uint16, hi uint8) {
	lo = uint16(e >> 4)
	hi = uint8(e)
	return
}

func (e *Entry12) Reset() {
	*e = 0
}
