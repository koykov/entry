package entry

// Entry48 compacts uint64 lo and uint16 hi values.
type Entry48 uint64

func NewEntry48(lo uint64, hi uint16) Entry48 {
	var e Entry48
	e.Encode(lo, hi)
	return e
}

// Encode combines uint48 lo and uint16 hi values into one uint64 int.
func (e *Entry48) Encode(lo uint64, hi uint16) {
	*e = Entry48(lo)<<16 | Entry48(hi)
}

// Decode splits e to uint48 lo and uint16 hi values.
func (e Entry48) Decode() (lo uint64, hi uint16) {
	lo = uint64(e >> 16)
	hi = uint16(e)
	return
}

func (e *Entry48) Reset() {
	*e = 0
}
