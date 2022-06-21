package entry

type Entry16 uint32

func (e *Entry16) Encode(lo, hi uint8) {
	*e = Entry16(lo)<<8 | Entry16(hi)
}

func (e Entry16) Decode() (lo, hi uint8) {
	lo = uint8(e >> 8)
	hi = uint8(e)
	return
}

func (e *Entry16) Reset() {
	*e = 0
}
