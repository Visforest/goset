package goset

// UintSet is a shortcut for Set[uint]
type UintSet = Set[uint]

func NewUintSet(vals ...uint) *UintSet {
	s := &UintSet{}
	s.Add(vals...)
	return s
}
