package entry

// Entry12 compacts uint12 lo and uint4 hi values.
type Entry12 uint16

func NewEntry12(lo uint16, hi uint8) Entry12 {
	var e Entry12
	e.Encode(lo, hi)
	return e
}

// Encode combines uint12 lo and uint4 hi values into one uint16 int.
func (e *Entry12) Encode(lo uint16, hi uint8) {
	*e = Entry12(lo)<<4 | Entry12(hi)
}

// Decode splits e to lo/hi uint8 values.
func (e *Entry12) Decode() (lo uint16, hi uint8) {
	lo = uint16(*e >> 4)
	hi = uint8(*e)
	return
}

// Lo return low part of entry.
func (e *Entry12) Lo() uint16 {
	return uint16(*e >> 4)
}

// Hi return high part of entry.
func (e *Entry12) Hi() uint8 {
	return uint8(*e)
}

func (e *Entry12) Reset() {
	*e = 0
}
