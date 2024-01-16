package goset

// FiloSet is a set that first in, last out
type FiloSet[T comparable] struct {
	*FifoSet[T]
}

func NewFiloSet[T comparable](vals ...T) *FiloSet[T] {
	return &FiloSet[T]{NewFifoSet[T](vals...)}
}

func (s *FiloSet[T]) ToList() []T {
	defer s.m.RUnlock()
	s.m.RLock()

	if s == nil || s.Length() == 0 {
		return nil
	}

	r := make([]T, 0, len(s.data))
	cur := s.tail
	for cur != nil {
		r = append(r, cur.val)
		cur = cur.pre
	}
	return r
}
