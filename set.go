package goset

import (
	"reflect"
	"sync"
)

type Set struct {
	m    sync.RWMutex
	data map[interface{}]struct{}
}

// NewSet creates a new Set
func NewSet(v ...interface{}) *Set {
	s := &Set{data: map[interface{}]struct{}{}}
	if len(v) > 0 {
		s.Add(v...)
	}
	return s
}

// Add add elements
func (s *Set) Add(v ...interface{}) {
	defer s.m.Unlock()
	s.m.Lock()
	for _, ele := range v {
		s.data[ele] = struct{}{}
	}
}

// Delete delete elements
func (s *Set) Delete(v ...interface{}) {
	defer s.m.Unlock()
	s.m.Lock()
	for _, ele := range v {
		delete(s.data, ele)
	}
}

// Clear clears all elements
func (s *Set) Clear() {
	defer s.m.Unlock()
	s.m.Lock()
	s.data = map[interface{}]struct{}{}
}

// Copy returns a deep copy of itself
func (s *Set) Copy() *Set {
	defer s.m.RUnlock()
	s.m.RLock()

	data := make(map[interface{}]struct{})
	for v := range s.data {
		data[v] = struct{}{}
	}
	return &Set{data: data}
}

// Length returns Set length
func (s *Set) Length() int {
	return len(s.data)
}

// Has returns whether v exists in Set
func (s *Set) Has(v interface{}) bool {
	_, ok := s.data[v]
	return ok
}

// ToList returns data slice
func (s *Set) ToList() []interface{} {
	defer s.m.RUnlock()
	s.m.RLock()

	var data = make([]interface{}, s.Length())
	var i int
	for d := range s.data {
		data[i] = d
		i++
	}
	return data
}

// Equals returns whether Set s has the same members with Set t
func (s *Set) Equals(t *Set) bool {
	if t == nil {
		return false
	}
	return reflect.DeepEqual(s.data, t.data)
}

// IsSub returns whether it's a part of Set t
func (s *Set) IsSub(t *Set) bool {
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
func (s *Set) Union(t *Set) *Set {
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
func (s *Set) Intersect(t *Set) *Set {
	var r = NewSet()
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
func (s *Set) Subtract(t *Set) *Set {
	if t == nil || t.Length() == 0 {
		return s.Copy()
	}
	var r = NewSet()
	if s == t {
		// subtract itself
		return r
	}

	s.m.RLock()
	t.m.RLock()

	defer s.m.RUnlock()
	defer t.m.RUnlock()

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
func (s *Set) Complement(t *Set) *Set {
	if t == nil || s.Length() == 0 || t.Length() == 0 {
		return s.Copy()
	}

	if s == t {
		return NewSet()
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
