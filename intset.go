package goset

// IntSet is a shortcut for Set[int]
type IntSet = Set[int]

func NewIntSet(vals ...int) *IntSet {
	s := &IntSet{}
	s.Add(vals...)
	return s
}
