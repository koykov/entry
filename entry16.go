package entry

// Entry16 compacts uint8 lo/hi values.
type Entry16 uint16

func NewEntry16(lo, hi uint8) Entry16 {
	var e Entry16
	e.Encode(lo, hi)
	return e
}

// Encode combines lo/hi values into one uint16 int.
func (e *Entry16) Encode(lo, hi uint8) {
	*e = Entry16(lo)<<8 | Entry16(hi)
}

// Decode splits e to lo/hi uint8 values.
func (e *Entry16) Decode() (lo, hi uint8) {
	lo = uint8(*e >> 8)
	hi = uint8(*e)
	return
}

// Lo return low part of entry.
func (e *Entry16) Lo() uint8 {
	return uint8(*e >> 8)
}

// Hi return high part of entry.
func (e *Entry16) Hi() uint8 {
	return uint8(*e)
}

func (e *Entry16) Reset() {
	*e = 0
}
