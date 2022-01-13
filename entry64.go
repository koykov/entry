package entry

type Entry64 uint64

func (e *Entry64) Encode(lo, hi uint32) {
	*e = Entry64(lo)<<32 | Entry64(hi)
}

func (e Entry64) Decode() (lo, hi uint32) {
	lo = uint32(e >> 32)
	hi = uint32(e)
	return
}

func (e *Entry64) Reset() {
	*e = 0
}
