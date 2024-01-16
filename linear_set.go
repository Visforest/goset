package goset

import (
	"reflect"
	"sync"
)

type setNode[T comparable] struct {
	val  T
	pre  *setNode[T]
	next *setNode[T]
}

type linearSet[T comparable] struct {
	m    sync.RWMutex
	head *setNode[T]
	tail *setNode[T]
	data map[T]*setNode[T]
}

func newLinearSet[T comparable](vals ...T) *linearSet[T] {
	s := &linearSet[T]{data: make(map[T]*setNode[T])}
	s.Add(vals...)
	return s
}

func (s *linearSet[T]) Add(v ...T) {
	if len(v) == 0 {
		return
	}
	defer s.m.Unlock()
	s.m.Lock()

	var i int
	if s.head == nil {
		// first node
		n := &setNode[T]{
			val: v[i],
		}
		s.head = n
		s.tail = n
		s.data[v[i]] = n
		i++
	}
	for ; i < len(v); i++ {
		if _, ok := s.data[v[i]]; !ok {
			n := &setNode[T]{
				val:  v[i],
				pre:  s.tail,
				next: nil,
			}
			s.tail.next = n
			s.tail = n
			s.data[v[i]] = n
		}
	}
}

func (s *linearSet[T]) Delete(v ...T) {
	defer s.m.Unlock()
	s.m.Lock()

	for i := range v {
		if n, ok := s.data[v[i]]; ok {
			if n.pre != nil {
				n.pre.next = n.next
			}
			if n.next != nil {
				n.next.pre = n.pre
			}
			delete(s.data, v[i])
		}
	}
	if len(s.data) == 0 {
		s.head = nil
		s.tail = nil
	}
}

func (s *linearSet[T]) Clear() {
	defer s.m.Unlock()
	s.m.Lock()

	s.head = nil
	s.tail = nil
	s.data = make(map[T]*setNode[T])
}

// Length returns linearSet length
func (s *linearSet[T]) Length() int {
	return len(s.data)
}

// Has returns whether v exists in linearSet
func (s *linearSet[T]) Has(v T) bool {
	_, ok := s.data[v]
	return ok
}

// Copy returns a deep copy of itself
func (s *linearSet[T]) Copy() *linearSet[T] {
	defer s.m.RUnlock()
	s.m.RLock()

	r := newLinearSet[T]()
	cur := s.head
	for cur != nil {
		r.Add(cur.val)
		cur = cur.next
	}
	return r
}

// ToList returns data slice
func (s *linearSet[T]) ToList() []T {
	defer s.m.RUnlock()
	s.m.RLock()

	if s == nil || s.Length() == 0 {
		return nil
	}

	r := make([]T, 0, len(s.data))
	cur := s.head
	for cur != nil {
		r = append(r, cur.val)
		cur = cur.next
	}
	return r
}

// Equals returns whether linearSet s has the same members with linearSet t
func (s *linearSet[T]) Equals(t *linearSet[T]) bool {
	if t == nil {
		return false
	}
	return reflect.DeepEqual(s.data, t.data)
}

// IsSub returns if it's a part of Set t.
// Note that it's defined that nil is sub of any linearSet
func (s *linearSet[T]) IsSub(t *linearSet[T]) bool {
	if t == nil {
		return false
	}
	if s == nil || s == t {
		return true
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	if s.Length() > t.Length() {
		return false
	}
	for k := range s.data {
		if !t.Has(k) {
			return false
		}
	}
	return true
}

func (s *linearSet[T]) Union(t *linearSet[T]) *linearSet[T] {
	r := s.Copy()
	if t == nil || s == t {
		return r
	}
	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()
	for d := range t.data {
		r.Add(d)
	}
	return r
}

func (s *linearSet[T]) Intersect(t *linearSet[T]) *linearSet[T] {
	r := newLinearSet[T]()
	if s == nil || t == nil || s.Length() == 0 || t.Length() == 0 {
		return r
	}
	if s == t {
		// intersect itself
		return s.Copy()
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()
	if s.Length() >= t.Length() {
		for v := range t.data {
			if s.Has(v) {
				r.Add(v)
			}
		}
	} else {
		for v := range s.data {
			if t.Has(v) {
				r.Add(v)
			}
		}
	}
	return r
}

func (s *linearSet[T]) Subtract(t *linearSet[T]) *linearSet[T] {
	if t == nil || t.Length() == 0 {
		return s.Copy()
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	r := s.Copy()
	for v := range s.data {
		if t.Has(v) {
			r.Delete(v)
		}
	}
	return r
}

func (s *linearSet[T]) Complement(t *linearSet[T]) *linearSet[T] {
	if s == nil || t == nil || s.Length() == 0 || t.Length() == 0 {
		return s.Copy()
	}

	if s == t {
		return newLinearSet[T]()
	}

	s.m.RLock()
	t.m.RLock()
	defer s.m.RUnlock()
	defer t.m.RUnlock()

	var r = s.Union(t)
	if s.Length() >= t.Length() {
		for v := range t.data {
			if s.Has(v) {
				r.Delete(v)
			}
		}
	} else {
		for v := range s.data {
			if t.Has(v) {
				r.Delete(v)
			}
		}
	}
	return r
}
