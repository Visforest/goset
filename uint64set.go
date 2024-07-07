package goset

// Uint64Set is a shortcut for Set[uint]
type Uint64Set = Set[uint64]

func NewUint64Set(vals ...uint64) *Uint64Set {
	s := &Uint64Set{}
	s.Add(vals...)
	return s
}
