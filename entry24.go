package entry

import "fmt"

// Entry24 compacts uint24 lo and uint8 hi values.
type Entry24 uint32

func NewEntry24(lo uint32, hi uint8) Entry24 {
	var e Entry24
	e.Encode(lo, hi)
	return e
}

// Encode combines uint24 lo and uint8 hi values into one uint64 int.
func (e *Entry24) Encode(lo uint32, hi uint8) {
	*e = Entry24(lo)<<8 | Entry24(hi)
}

// Decode splits e to lo/hi uint16 values.
func (e *Entry24) Decode() (lo uint32, hi uint8) {
	lo = uint32(*e >> 8)
	hi = uint8(*e)
	return
}

// Lo return low part of entry.
func (e *Entry24) Lo() uint32 {
	return uint32(*e >> 8)
}

// Hi return high part of entry.
func (e *Entry24) Hi() uint8 {
	return uint8(*e)
}

func (e *Entry24) Reset() {
	*e = 0
}

func (e *Entry24) DebugString() string {
	return fmt.Sprintf("%d: lo %d hi %d", *e, e.Lo(), e.Hi())
}
