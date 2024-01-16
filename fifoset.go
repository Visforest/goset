package goset

// FifoSet is a set that first in,first out
type FifoSet[T comparable] struct {
	*linearSet[T]
}

func NewFifoSet[T comparable](vals ...T) *FifoSet[T] {
	return &FifoSet[T]{newLinearSet[T](vals...)}
}

func (l *FifoSet[T]) Add(v ...T) {
	if len(v) == 0 {
		return
	}
	defer l.m.Unlock()
	l.m.Lock()

	var i int
	if l.head == nil {
		// first node
		n := &setNode[T]{
			val: v[i],
		}
		l.head = n
		l.tail = n
		l.data[v[i]] = n
		i++
	}
	for ; i < len(v); i++ {
		if _, ok := l.data[v[i]]; !ok {
			n := &setNode[T]{
				val:  v[i],
				pre:  l.tail,
				next: nil,
			}
			l.tail.next = n
			l.tail = n
			l.data[v[i]] = n
		}
	}
}
