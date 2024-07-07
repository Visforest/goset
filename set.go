package goset

import (
	"reflect"
	"sync"
)

type Set[T comparable] struct {
	m    sync.RWMutex
	data map[T]struct{}
}

// NewSet creates a new Set
func NewSet[T comparable](v ...T) *Set[T] {
	s := &Set[T]{data: map[T]struct{}{}}
	if len(v) > 0 {
		s.Add(v...)
	}
	return s
}

// Add adds elements
func (s *Set[T]) Add(v ...T) {
	s.m.Lock()
	defer s.m.Unlock()

	if s.data == nil {
		s.data = make(map[T]struct{})
	}
	for _, ele := range v {
		s.data[ele] = struct{}{}
	}
}

// Delete delete elements
func (s *Set[T]) Delete(v ...T) {
	s.m.Lock()
	defer s.m.Unlock()

	if s.data == nil {
		s.data = make(map[T]struct{})
	}
	for _, ele := range v {
		delete(s.data, ele)
	}
}

// Clear clears all elements
func (s *Set[T]) Clear() {
	s.m.Lock()
	defer s.m.Unlock()

	s.data = make(map[T]struct{})
}

// Copy returns a deep copy of itself
func (s *Set[T]) Copy() *Set[T] {
	s.m.RLock()
	defer s.m.RUnlock()

	if s.data == nil {
		s.data = make(map[T]struct{})
	}

	data := make(map[T]struct{})
	for v := range s.data {
		data[v] = struct{}{}
	}
	return &Set[T]{data: data}
}

// Length returns Set length
func (s *Set[T]) Length() int {
	return len(s.data)
}

// Has returns whether v exists in Set
func (s *Set[T]) Has(v T) bool {
	s.m.RLock()
	defer s.m.RUnlock()

	if s.data == nil {
		s.data = make(map[T]struct{})
	}
	_, ok := s.data[v]
	return ok
}

// ToList returns data slice
func (s *Set[T]) ToList() []T {
	s.m.RLock()
	defer s.m.RUnlock()

	if s.data == nil {
		s.data = make(map[T]struct{})
	}

	var data = make([]T, s.Length())
	var i int
	for d := range s.data {
		data[i] = d
		i++
	}
	return data
}

// Equals returns whether Set s has the same members with Set t
func (s *Set[T]) Equals(t *Set[T]) bool {
	s.m.RLock()
	defer s.m.RUnlock()

	if t == nil {
		return false
	}
	return reflect.DeepEqual(s.data, t.data)
}

// IsSub returns whether it's a part of Set t
func (s *Set[T]) IsSub(t *Set[T]) bool {
	if t == nil {
		return false
	}
	if s == t {
		// compare with itself
		return true
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()
	if s.data == nil {
		s.data = make(map[T]struct{})
	}
	if t.data == nil {
		t.data = make(map[T]struct{})
	}

	if s.Length() > t.Length() {
		return false
	}
	for v := range s.data {
		if !t.Has(v) {
			return false
		}
	}
	return true
}

// Union unions with Set t and returns a new Set
//
// for example:
// var a=NewSet(1,2,3)
// var b=NewSet(2,3,4)
// a.Union(b) returns {1,2,3,4}
func (s *Set[T]) Union(t *Set[T]) *Set[T] {
	var r = s.Copy()
	if t == nil || s == t {
		return r
	}
	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	r.Add(t.ToList()...)
	return r
}

// Intersect returns a new Set Whose elements exist in both Set
//
// for example:
// var a=NewSet(1,2,3)
// var b=NewSet(2,3,4)
// a.Intersect(b) returns {2,3}
func (s *Set[T]) Intersect(t *Set[T]) *Set[T] {
	var r = NewSet[T]()
	if t == nil || s.Length() == 0 || t.Length() == 0 {
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

	if s.data == nil {
		s.data = make(map[T]struct{})
	}
	if t.data == nil {
		t.data = make(map[T]struct{})
	}

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

// Subtract returns a new Set Whose elements exist in itself but don't exist in Set t
//
// for example:
// var a=NewSet(1,2,3)
// var b=NewSet(2,3,4)
// a.Subtract(b) returns {1}
func (s *Set[T]) Subtract(t *Set[T]) *Set[T] {
	if t == nil || t.Length() == 0 {
		return s.Copy()
	}
	var r = NewSet[T]()
	if s == t {
		// subtract itself
		return r
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	if s.data == nil {
		s.data = make(map[T]struct{})
	}
	if t.data == nil {
		t.data = make(map[T]struct{})
	}

	for v := range s.data {
		if !t.Has(v) {
			r.Add(v)
		}
	}
	return r
}

// Complement returns a new Set Whose elements only exists in one Set
//
// for example:
// var a=NewSet(1,2,3)
// var b=NewSet(2,3,4)
// a.Complement(b) returns {1,4}
func (s *Set[T]) Complement(t *Set[T]) *Set[T] {
	if t == nil || s.Length() == 0 || t.Length() == 0 {
		return s.Copy()
	}

	if s == t {
		return NewSet[T]()
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

	if s.data == nil {
		s.data = make(map[T]struct{})
	}
	if t.data == nil {
		t.data = make(map[T]struct{})
	}

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
