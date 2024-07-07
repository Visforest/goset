package goset

// StrSet is a shortcut for Set[string]
type StrSet = Set[string]

func NewStrSet(vals ...string) *StrSet {
	s := &StrSet{}
	s.Add(vals...)
	return s
}
