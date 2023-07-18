package entry

// Entry64 compacts uint32 lo/hi values.
type Entry64 uint64

func NewEntry64(lo, hi uint32) Entry64 {
	var e Entry64
	e.Encode(lo, hi)
	return e
}

// Encode combines lo/hi values into one uint64 int.
func (e *Entry64) Encode(lo, hi uint32) Entry64 {
	*e = Entry64(lo)<<32 | Entry64(hi)
	return *e
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
