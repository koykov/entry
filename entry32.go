package entry

type Entry32 uint32

func (e *Entry32) Encode(lo, hi uint16) {
	*e = Entry32(lo)<<16 | Entry32(hi)
}

func (e Entry32) Decode() (lo, hi uint16) {
	lo = uint16(e >> 16)
	hi = uint16(e)
	return
}

func (e *Entry32) Reset() {
	*e = 0
}
