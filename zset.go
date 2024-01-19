package goset

import (
	"cmp"
)

func NewSortedSet[T cmp.Ordered](vals ...T) *SortedSet[T] {
	s := &SortedSet[T]{newLinearSet[T](addSorted[T])}
	s.Add(vals...)
	return s
}

// SortedSet is a set whose elements are stored in asc order
type SortedSet[T cmp.Ordered] struct {
	*linearSet[T]
}

func (s *SortedSet[T]) Add(vals ...T) {
	addSorted(s.linearSet, vals...)
}

// Copy returns a deep copy of itself
func (s *SortedSet[T]) Copy() *SortedSet[T] {
	return &SortedSet[T]{s.linearSet.copy(addSorted[T])}
}

func (s *SortedSet[T]) Equals(t *SortedSet[T]) bool {
	return s.linearSet.Equals(t.linearSet)
}

func (s *SortedSet[T]) IsSub(t *SortedSet[T]) bool {
	return s.linearSet.IsSub(t.linearSet)
}

func (s *SortedSet[T]) Union(t *SortedSet[T]) *SortedSet[T] {
	return &SortedSet[T]{s.linearSet.union(t.linearSet, addSorted[T])}
}

func (s *SortedSet[T]) Subtract(t *SortedSet[T]) *SortedSet[T] {
	return &SortedSet[T]{s.linearSet.subtract(t.linearSet, addSorted[T])}
}

func (s *SortedSet[T]) Intersect(t *SortedSet[T]) *SortedSet[T] {
	return &SortedSet[T]{s.linearSet.intersect(t.linearSet, addSorted[T])}
}

func (s *SortedSet[T]) Complement(t *SortedSet[T]) *SortedSet[T] {
	return &SortedSet[T]{s.linearSet.complement(t.linearSet, addSorted[T])}
}
