package goset

// Int64Set is a shortcut for Set[int64]
type Int64Set = Set[int64]

func NewInt64Set(vals ...int64) *Int64Set {
	s := &Int64Set{}
	s.Add(vals...)
	return s
}
