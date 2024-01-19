package goset

// FiloSet is a set that first in, last out
type FiloSet[T comparable] struct {
	*linearSet[T]
}

func NewFiloSet[T comparable](vals ...T) *FiloSet[T] {
	return &FiloSet[T]{newLinearSet[T](addFilo[T], vals...)}
}

func (s *FiloSet[T]) Add(vals ...T) {
	addFilo(s.linearSet, vals...)
}

func (s *FiloSet[T]) Copy() *FiloSet[T] {
	return &FiloSet[T]{s.linearSet.copy(addFifo[T])}
}

func (s *FiloSet[T]) Equals(t *FiloSet[T]) bool {
	return s.linearSet.Equals(t.linearSet)
}

func (s *FiloSet[T]) IsSub(t *FiloSet[T]) bool {
	return s.linearSet.IsSub(t.linearSet)
}

func (s *FiloSet[T]) Union(t *FiloSet[T]) *FiloSet[T] {
	return &FiloSet[T]{s.linearSet.union(t.linearSet, addFifo[T])}
}

func (s *FiloSet[T]) Subtract(t *FiloSet[T]) *FiloSet[T] {
	return &FiloSet[T]{s.linearSet.subtract(t.linearSet, addFifo[T])}
}

func (s *FiloSet[T]) Intersect(t *FiloSet[T]) *FiloSet[T] {
	return &FiloSet[T]{s.linearSet.intersect(t.linearSet, addFifo[T])}
}

func (s *FiloSet[T]) Complement(t *FiloSet[T]) *FiloSet[T] {
	return &FiloSet[T]{s.linearSet.complement(t.linearSet, addFifo[T])}
}
