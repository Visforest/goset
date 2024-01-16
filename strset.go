package goset

// StrSet is a shortcut for Set[string]
type StrSet = Set[string]

func NewStrSet(vals ...string) *StrSet {
	s := &StrSet{Data: make(map[string]struct{})}
	s.Add(vals...)
	return s
}
