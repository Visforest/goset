package goset

import (
	"reflect"
	"sync"
)

func NewFifoSet[T comparable]() *FifoSet[T] {
	return &FifoSet[T]{data: make(map[T]*setNode[T])}
}

type setNode[T comparable] struct {
	val  T
	pre  *setNode[T]
	next *setNode[T]
}

type FifoSet[T comparable] struct {
	m    sync.RWMutex
	head *setNode[T]
	tail *setNode[T]
	data map[T]*setNode[T]
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

func (l *FifoSet[T]) Delete(v ...T) {
	defer l.m.Unlock()
	l.m.Lock()

	for i := range v {
		if n, ok := l.data[v[i]]; ok {
			if n.pre != nil {
				n.pre.next = n.next
			}
			if n.next != nil {
				n.next.pre = n.pre
			}
			delete(l.data, v[i])
		}
	}
	if len(l.data) == 0 {
		l.head = nil
		l.tail = nil
	}
}

func (l *FifoSet[T]) Clear() {
	defer l.m.Unlock()
	l.m.Lock()

	l.head = nil
	l.tail = nil
	l.data = make(map[T]*setNode[T])
}

// Copy returns a deep copy of itself
func (l *FifoSet[T]) Copy() *FifoSet[T] {
	defer l.m.RUnlock()
	l.m.RLock()

	r := NewFifoSet[T]()
	cur := l.head
	for cur != nil {
		r.Add(cur.val)
		cur = cur.next
	}
	return r
}

// Length returns FifoSet length
func (l *FifoSet[T]) Length() int {
	return len(l.data)
}

// Has returns whether v exists in FifoSet
func (l *FifoSet[T]) Has(v any) bool {
	_, ok := l.data[v]
	return ok
}

// ToList returns data slice
func (l *FifoSet[T]) ToList() []any {
	defer l.m.RUnlock()
	l.m.RLock()

	r := make([]any, 0, len(l.data))
	cur := l.head
	for cur != nil {
		r = append(r, cur.val)
	}
	return r
}

// Equals returns whether FifoSet l has the same members with FifoSet t
func (l *FifoSet[T]) Equals(t *FifoSet[T]) bool {
	if t == nil {
		return false
	}
	return reflect.DeepEqual(l.data, t.data)
}
