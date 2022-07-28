package entry

// Entry24 compacts uint24 lo and uint8 hi values.
type Entry24 uint32

// Encode combines uint24 lo and uint8 hi values into one uint64 int.
func (e *Entry24) Encode(lo uint32, hi uint8) {
	*e = Entry24(lo)<<8 | Entry24(hi)
}

// Decode splits e to lo/hi uint16 values.
func (e Entry24) Decode() (lo uint32, hi uint8) {
	lo = uint32(e >> 8)
	hi = uint8(e)
	return
}

func (e *Entry24) Reset() {
	*e = 0
}
