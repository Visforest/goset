package goset

import (
	"cmp"
	"reflect"
	"sync"
)

func addFifo[T comparable](s *linearSet[T], vals ...T) {
	if len(vals) == 0 {
		return
	}
	defer s.m.Unlock()
	s.m.Lock()

	var i int
	if s.head == nil {
		// first node
		n := &setNode[T]{
			val: vals[i],
		}
		s.head = n
		s.tail = n
		s.data[vals[i]] = n
		i++
	}
	for ; i < len(vals); i++ {
		if _, ok := s.data[vals[i]]; !ok {
			n := &setNode[T]{
				val:  vals[i],
				pre:  s.tail,
				next: nil,
			}
			s.tail.next = n
			s.tail = n
			s.data[vals[i]] = n
		}
	}
}

func addFilo[T comparable](s *linearSet[T], vals ...T) {
	if len(vals) == 0 {
		return
	}
	defer s.m.Unlock()
	s.m.Lock()

	var i int
	if s.tail == nil {
		// put first value into tail node
		n := &setNode[T]{
			val: vals[i],
		}
		s.head = n
		s.tail = n
		s.data[vals[i]] = n
		i++
	}
	for ; i < len(vals); i++ {
		if _, ok := s.data[vals[i]]; !ok {
			n := &setNode[T]{
				val: vals[i],
			}
			n.next = s.head
			s.head.pre = n
			s.head = n
			s.data[vals[i]] = n
		}
	}
}

func addSorted[T cmp.Ordered](s *linearSet[T], vals ...T) {
	if len(vals) == 0 {
		return
	}
	defer s.m.Unlock()
	s.m.Lock()

	var i int
	if s.head == nil {
		// first node
		n := &setNode[T]{
			val: vals[i],
		}
		s.head = n
		s.tail = n
		s.data[vals[i]] = n
		i++
	}
	for ; i < len(vals); i++ {
		if _, ok := s.data[vals[i]]; !ok {
			n := &setNode[T]{
				val: vals[i],
			}
			if cmp.Less(vals[i], s.head.val) {
				// add to head
				n.next = s.head
				s.head.pre = n
				s.head = n
				s.data[vals[i]] = n
			} else if cmp.Less(s.tail.val, vals[i]) {
				//	add to tail
				n.pre = s.tail
				s.tail.next = n
				s.tail = n
				s.data[vals[i]] = n
			} else {
				// search and insert
				left := s.head
				right := left.next
				for right != nil {
					if cmp.Less(vals[i], right.val) {
						// insert and break
						left.next = n
						right.pre = n
						n.pre = left
						n.next = right
						s.data[vals[i]] = n
						break
					}
					// go on
					right = right.next
					left = left.next
				}
			}
		}
	}
}

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

func newLinearSet[T comparable](add func(s *linearSet[T], vals ...T), vals ...T) *linearSet[T] {
	s := &linearSet[T]{data: make(map[T]*setNode[T])}
	add(s, vals...)
	return s
}

func (s *linearSet[T]) Delete(vals ...T) {
	defer s.m.Unlock()
	s.m.Lock()

	for _, v := range vals {
		if n, ok := s.data[v]; ok {
			if n.pre == nil {
				s.head = n.next
			} else {
				n.pre.next = n.next
			}
			if n.next == nil {
				s.tail = n.pre
			} else {
				n.next.pre = n.pre
			}
			delete(s.data, v)
		}
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
func (s *linearSet[T]) copy(add func(s *linearSet[T], vals ...T)) *linearSet[T] {
	defer s.m.RUnlock()
	s.m.RLock()

	r := newLinearSet[T](add)
	cur := s.head
	for cur != nil {
		add(r, cur.val)
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

func (s *linearSet[T]) union(t *linearSet[T], add func(s *linearSet[T], vals ...T)) *linearSet[T] {
	r := s.copy(add)
	if t == nil || s == t {
		return r
	}
	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()
	for _, d := range t.ToList() {
		add(r, d)
	}
	return r
}

func (s *linearSet[T]) intersect(t *linearSet[T], add func(s *linearSet[T], vals ...T)) *linearSet[T] {
	r := newLinearSet[T](add)
	if s == nil || t == nil || s.Length() == 0 || t.Length() == 0 {
		return r
	}
	if s == t {
		// intersect itself
		return s.copy(add)
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()
	if s.Length() >= t.Length() {
		for _, v := range t.ToList() {
			if s.Has(v) {
				add(r, v)
			}
		}
	} else {
		for _, v := range s.ToList() {
			if t.Has(v) {
				add(r, v)
			}
		}
	}
	return r
}

func (s *linearSet[T]) subtract(t *linearSet[T], add func(s *linearSet[T], vals ...T)) *linearSet[T] {
	if t == nil || t.Length() == 0 {
		return s.copy(add)
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	r := s.copy(add)
	for _, v := range s.ToList() {
		if t.Has(v) {
			r.Delete(v)
		}
	}
	return r
}

func (s *linearSet[T]) complement(t *linearSet[T], add func(s *linearSet[T], vals ...T)) *linearSet[T] {
	if s == nil || t == nil || s.Length() == 0 || t.Length() == 0 {
		return s.copy(add)
	}

	if s == t {
		return newLinearSet[T](add)
	}

	s.m.RLock()
	t.m.RLock()
	defer s.m.RUnlock()
	defer t.m.RUnlock()

	var r = s.union(t, add)
	if s.Length() >= t.Length() {
		for _, v := range t.ToList() {
			if s.Has(v) {
				r.Delete(v)
			}
		}
	} else {
		for _, v := range s.ToList() {
			if t.Has(v) {
				r.Delete(v)
			}
		}
	}
	return r
}
