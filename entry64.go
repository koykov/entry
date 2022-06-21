package entry

// Entry64 compacts uint32 lo/hi values.
type Entry64 uint64

// Entry64 compacts uint32 lo/hi values.
func (e *Entry64) Encode(lo, hi uint32) {
	*e = Entry64(lo)<<32 | Entry64(hi)
}

// Decode splits e to lo/hi uint32 values.
func (e Entry64) Decode() (lo, hi uint32) {
	lo = uint32(e >> 32)
	hi = uint32(e)
	return
}

func (e *Entry64) Reset() {
	*e = 0
}
