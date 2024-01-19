package goset

// FifoSet is a set whose elements are stored by fifo
type FifoSet[T comparable] struct {
	*linearSet[T]
}

func NewFifoSet[T comparable](vals ...T) *FifoSet[T] {
	return &FifoSet[T]{newLinearSet[T](addFifo[T], vals...)}
}

func (s *FifoSet[T]) Add(vals ...T) {
	addFifo(s.linearSet, vals...)
}

// Copy returns a deep copy of itself
func (s *FifoSet[T]) Copy() *FifoSet[T] {
	return &FifoSet[T]{s.linearSet.copy(addFifo[T])}
}

func (s *FifoSet[T]) Equals(t *FifoSet[T]) bool {
	return s.linearSet.Equals(t.linearSet)
}

func (s *FifoSet[T]) IsSub(t *FifoSet[T]) bool {
	return s.linearSet.IsSub(t.linearSet)
}

func (s *FifoSet[T]) Union(t *FifoSet[T]) *FifoSet[T] {
	return &FifoSet[T]{s.linearSet.union(t.linearSet, addFifo[T])}
}

func (s *FifoSet[T]) Subtract(t *FifoSet[T]) *FifoSet[T] {
	return &FifoSet[T]{s.linearSet.subtract(t.linearSet, addFifo[T])}
}

func (s *FifoSet[T]) Intersect(t *FifoSet[T]) *FifoSet[T] {
	return &FifoSet[T]{s.linearSet.intersect(t.linearSet, addFifo[T])}
}

func (s *FifoSet[T]) Complement(t *FifoSet[T]) *FifoSet[T] {
	return &FifoSet[T]{s.linearSet.complement(t.linearSet, addFifo[T])}
}
